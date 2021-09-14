package storage

import (
	"context"
	sectorstorage "github.com/filecoin-project/lotus/extern/sector-storage"
)

type DistributeProver struct {
	prover sectorstorage.SectorManager
}

func NewDistributeProver(prover sectorstorage.SectorManager) *DistributeProver {
	return &DistributeProver{prover: prover}
}

func (dp *DistributeProver) Run(ctx context.Context) {
	// 启动服务端
	log.Warnw("[yunjie]: DistributeProver started!!")
}


