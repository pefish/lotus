package storage

import (
	"context"
	distribute_prover "github.com/filecoin-project/lotus/proto/distribute-prover"
	register_server "github.com/filecoin-project/lotus/proto/register-server"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
	"sync"
	"time"

	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/dline"
	"github.com/filecoin-project/specs-storage/storage"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/store"
	"github.com/filecoin-project/lotus/chain/types"
	sectorstorage "github.com/filecoin-project/lotus/extern/sector-storage"
	"github.com/filecoin-project/lotus/extern/sector-storage/ffiwrapper"
	"github.com/filecoin-project/lotus/journal"
	"github.com/filecoin-project/lotus/node/config"

	"go.opencensus.io/trace"
)

// WindowPoStScheduler is the coordinator for WindowPoSt submissions, fault
// declaration, and recovery declarations. It watches the chain for reverts and
// applies, and schedules/run those processes as partition deadlines arrive.
//
// WindowPoStScheduler watches the chain though the changeHandler, which in turn
// turn calls the scheduler when the time arrives to do work.
type WindowPoStScheduler struct {
	api              fullNodeFilteredAPI
	feeCfg           config.MinerFeeConfig
	addrSel          *AddressSelector
	prover           storage.Prover
	verifier         ffiwrapper.Verifier
	faultTracker     sectorstorage.FaultTracker
	proofType        abi.RegisteredPoStProof
	partitionSectors uint64
	ch               *changeHandler

	actor address.Address

	evtTypes [4]journal.EventType
	journal  journal.Journal

	// failed abi.ChainEpoch // eps
	// failLk sync.Mutex

	activeWdPosters sync.Map  // wdPoster url -> ActiveWdPosterData
	failedWdPosters sync.Map  // wdPoster url -> true/false
	partitionWdPoster sync.Map  // partition index -> wdPoster url
}

// 幂等
func (s *WindowPoStScheduler) Register(ctx context.Context, request *register_server.RegisterRequest) (*register_server.RegisterReply, error) {
	if _, ok := s.activeWdPosters.Load(request.Url); ok {
		return &register_server.RegisterReply{Msg: "ok"}, nil
	}
	err := s.connectOneWdPoster(request.Url)
	if err != nil {
		return nil, err
	}
	log.Infof("[yunjie]: new wdPoster %s", request.Url)
	return &register_server.RegisterReply{Msg: "ok"}, nil
}

type ActiveWdPosterData struct{
	Conn *grpc.ClientConn
	Client distribute_prover.DistributeProverClient
	Url string
}

// NewWindowedPoStScheduler creates a new WindowPoStScheduler scheduler.
func NewWindowedPoStScheduler(api fullNodeFilteredAPI,
	cfg config.MinerFeeConfig,
	as *AddressSelector,
	sp storage.Prover,
	verif ffiwrapper.Verifier,
	ft sectorstorage.FaultTracker,
	j journal.Journal,
	actor address.Address) (*WindowPoStScheduler, error) {
	mi, err := api.StateMinerInfo(context.TODO(), actor, types.EmptyTSK)
	if err != nil {
		return nil, xerrors.Errorf("getting sector size: %w", err)
	}

	return &WindowPoStScheduler{
		api:              api,
		feeCfg:           cfg,
		addrSel:          as,
		prover:           sp,
		verifier:         verif,
		faultTracker:     ft,
		proofType:        mi.WindowPoStProofType,
		partitionSectors: mi.WindowPoStPartitionSectors,

		actor: actor,
		evtTypes: [...]journal.EventType{
			evtTypeWdPoStScheduler:  j.RegisterEventType("wdpost", "scheduler"),
			evtTypeWdPoStProofs:     j.RegisterEventType("wdpost", "proofs_processed"),
			evtTypeWdPoStRecoveries: j.RegisterEventType("wdpost", "recoveries_processed"),
			evtTypeWdPoStFaults:     j.RegisterEventType("wdpost", "faults_processed"),
		},
		journal: j,
	}, nil
}

func (s *WindowPoStScheduler) connectWdPosters(wdPostServers []string)  {
	for _, wdPostServerUrl := range wdPostServers {
		s.connectOneWdPoster(wdPostServerUrl)
	}
}

func (s *WindowPoStScheduler) connectOneWdPoster(wdPostServerUrl string) error {
	log.Infof("[yunjie]: WindowPoStScheduler connecting wdPoster %s", wdPostServerUrl)
	var conn *grpc.ClientConn
	count := 3  // 重试次数
	i := 0
	for {
		i++
		dialCtx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
		conn_, err := grpc.DialContext(dialCtx, wdPostServerUrl, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Warnf("[yunjie]: WindowPoStScheduler connect wdPoster %s failed. err: %v", wdPostServerUrl, err)
			if i == count {
				s.activeWdPosters.Delete(wdPostServerUrl)
				s.failedWdPosters.Store(wdPostServerUrl, true)
				return errors.New("connect failed")
			}
			time.Sleep(2 * time.Second)
			continue
		}
		conn = conn_
		break
	}
	client := distribute_prover.NewDistributeProverClient(conn)
	s.failedWdPosters.Delete(wdPostServerUrl)
	s.activeWdPosters.Store(wdPostServerUrl, ActiveWdPosterData{
		Conn:   conn,
		Client: client,
		Url: wdPostServerUrl,
	})
	log.Infof("[yunjie]: WindowPoStScheduler connected wdPoster %s", wdPostServerUrl)
	return nil
}

func (s *WindowPoStScheduler) pingOneWdPoster(activeWdPoster ActiveWdPosterData)  {
	log.Infof("[yunjie]: WindowPoStScheduler pinging wdPoster %s", activeWdPoster.Url)
	count := 3  // 重试次数
	i := 0
	for {
		i++
		ctx, _ := context.WithTimeout(context.Background(), 3 * time.Second)
		reply, err := activeWdPoster.Client.Ping(ctx, &distribute_prover.PingRequest{})
		if err != nil || reply == nil || reply.Msg != "ok" {
			log.Warnf("[yunjie]: WindowPoStScheduler ping wdPoster %s failed. reply: %s, err: %v", activeWdPoster.Url, reply, err)
			if i == count {
				activeWdPoster.Conn.Close()
				s.activeWdPosters.Delete(activeWdPoster.Url)
				s.failedWdPosters.Store(activeWdPoster.Url, true)
				return
			}
			time.Sleep(2 * time.Second)
			continue
		}
		break
	}
	log.Infof("[yunjie]: WindowPoStScheduler ping wdPoster %s succeed", activeWdPoster.Url)
}

func (s *WindowPoStScheduler) heartbeatWdPosters()  {
	for {
		// 连接失败的反复尝试
		s.failedWdPosters.Range(func(key, value interface{}) bool {
			s.connectOneWdPoster(key.(string))
			return true
		})
		// 连接成功的反复 ping
		s.activeWdPosters.Range(func(key, value interface{}) bool {
			s.pingOneWdPoster(value.(ActiveWdPosterData))
			return true
		})

		time.Sleep(10 * time.Second)
	}
}

func (s *WindowPoStScheduler) startRegisterServer(ctx context.Context, registerServerUrl string) {
	lis, err := net.Listen("tcp", registerServerUrl)
	if err != nil {
		log.Errorf("[yunjie]: [RegisterServer] failed to listen: %v", err)
		return
	}

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
			defer func() {
				if err_ := recover(); err_ != nil {
					log.Error(err_)
					err = status.Errorf(codes.Internal, "%#v", err_)
				}
			}()
			log.Infof("[yunjie]: [RegisterServer] method: %s, param: %#v", info.FullMethod, srv)
			return handler(srv, ss)
		}),
		grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
			defer func() {
				if err_ := recover(); err_ != nil {
					log.Error(err_)
					err = status.Errorf(codes.Internal, "%#v", err_)
				}
			}()
			log.Infof("[yunjie]: [RegisterServer] method: %s, param: %#v", info.FullMethod, req)
			return handler(ctx, req)
		}),
	)
	register_server.RegisterRegisterServerServer(grpcServer, s)

	log.Infof(`[yunjie]: [RegisterServer] grpc server started. address: %s`, registerServerUrl)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Errorf("[yunjie]: [RegisterServer] failed to serve: %v", err)
		}
	}()

	<- ctx.Done()
	grpcServer.Stop()
	log.Infof(`[yunjie]: [RegisterServer] grpc server stopped. address: %s`, registerServerUrl)

}

func (s *WindowPoStScheduler) Close()  {
	s.activeWdPosters.Range(func(key, value interface{}) bool {
		value.(ActiveWdPosterData).Conn.Close()
		return true
	})
}

func (s *WindowPoStScheduler) Run(ctx context.Context, wdPostConfig config.WdPostConfig) {
	// Initialize change handler.

	// callbacks is a union of the fullNodeFilteredAPI and ourselves.
	callbacks := struct {
		fullNodeFilteredAPI
		*WindowPoStScheduler
	}{s.api, s}

	s.ch = newChangeHandler(callbacks, s.actor)
	defer s.ch.shutdown()
	s.ch.start()

	var (
		notifs <-chan []*api.HeadChange
		err    error
		gotCur bool
	)

	// 如果是分布式 wdposter ，则连接 wdposters
	if len(wdPostConfig.WdPostServers) > 0 {
		s.connectWdPosters(wdPostConfig.WdPostServers)
		// 协程开启心跳
		go s.heartbeatWdPosters()
	}
	// 启动注册中心
	go s.startRegisterServer(ctx, wdPostConfig.RegisterServerUrl)

	// not fine to panic after this point
	for {
		if notifs == nil {
			notifs, err = s.api.ChainNotify(ctx)  // 监听高度变化
			if err != nil {
				log.Errorf("ChainNotify error: %+v", err)

				build.Clock.Sleep(10 * time.Second)
				continue
			}

			gotCur = false
		}

		select {
		case changes, ok := <-notifs:  // 高度有变化
			if !ok {
				log.Warn("window post scheduler notifs channel closed")
				notifs = nil
				continue
			}

			if !gotCur {
				if len(changes) != 1 {
					log.Errorf("expected first notif to have len = 1")
					continue
				}
				chg := changes[0]
				if chg.Type != store.HCCurrent {
					log.Errorf("expected first notif to tell current ts")
					continue
				}

				ctx, span := trace.StartSpan(ctx, "WindowPoStScheduler.headChange")

				s.update(ctx, nil, chg.Val)

				span.End()
				gotCur = true
				continue
			}

			ctx, span := trace.StartSpan(ctx, "WindowPoStScheduler.headChange")

			var lowest, highest *types.TipSet = nil, nil

			for _, change := range changes {
				if change.Val == nil {
					log.Errorf("change.Val was nil")
				}
				switch change.Type {
				case store.HCRevert:
					lowest = change.Val
				case store.HCApply:
					highest = change.Val
				}
			}

			s.update(ctx, lowest, highest)

			span.End()
		case <-ctx.Done():
			s.Close()
			return
		}
	}
}

func (s *WindowPoStScheduler) update(ctx context.Context, revert, apply *types.TipSet) {
	if apply == nil {
		log.Error("no new tipset in window post WindowPoStScheduler.update")
		return
	}
	err := s.ch.update(ctx, revert, apply)
	if err != nil {
		log.Errorf("handling head updates in window post sched: %+v", err)
	}
}

// onAbort is called when generating proofs or submitting proofs is aborted
func (s *WindowPoStScheduler) onAbort(ts *types.TipSet, deadline *dline.Info) {
	s.journal.RecordEvent(s.evtTypes[evtTypeWdPoStScheduler], func() interface{} {
		c := evtCommon{}
		if ts != nil {
			c.Deadline = deadline
			c.Height = ts.Height()
			c.TipSet = ts.Cids()
		}
		return WdPoStSchedulerEvt{
			evtCommon: c,
			State:     SchedulerStateAborted,
		}
	})
}

func (s *WindowPoStScheduler) getEvtCommon(err error) evtCommon {
	c := evtCommon{Error: err}
	currentTS, currentDeadline := s.ch.currentTSDI()
	if currentTS != nil {
		c.Deadline = currentDeadline
		c.Height = currentTS.Height()
		c.TipSet = currentTS.Cids()
	}
	return c
}
