package storage

import (
	"context"
	sectorstorage "github.com/filecoin-project/lotus/extern/sector-storage"
	distribute_prover "github.com/filecoin-project/lotus/proto/distribute-prover"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
)

type DistributeProver struct {
	prover sectorstorage.SectorManager
}

func (dp *DistributeProver) GenerateWindowPoSt(ctx context.Context, request *distribute_prover.GenerateWindowPoStRequest) (*distribute_prover.GenerateWindowPoStReply, error) {
	panic("implement me")
}

func NewDistributeProver(prover sectorstorage.SectorManager) *DistributeProver {
	return &DistributeProver{prover: prover}
}

func (dp *DistributeProver) Run(ctx context.Context) {
	// 启动服务端
	log.Info("[yunjie]: DistributeProver started!!")

	address := `0.0.0.0:8000`
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Errorf("[yunjie]: failed to listen: %v", err)
		return
	}

	s := grpc.NewServer(
		grpc.StreamInterceptor(func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
			defer func() {
				if err_ := recover(); err_ != nil {
					log.Error(err_)
					err = status.Errorf(codes.Internal, "%#v", err_)
				}
			}()
			log.Debugf("[yunjie]: method: %s, param: %#v", info.FullMethod, srv)
			return handler(srv, ss)
		}),
		grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
			defer func() {
				if err_ := recover(); err_ != nil {
					log.Error(err_)
					err = status.Errorf(codes.Internal, "%#v", err_)
				}
			}()
			log.Debugf("[yunjie]: method: %s, param: %#v", info.FullMethod, req)
			return handler(ctx, req)
		}),
	)
	distribute_prover.RegisterDistributeProverServer(s, dp)

	log.Infof(`[yunjie]: grpc server started. address: %s`, address)
	if err := s.Serve(lis); err != nil {
		log.Errorf("[yunjie]: failed to serve: %v", err)
	}
}


