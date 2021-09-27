//+build cgo

package ffiwrapper

import (
	ffi "github.com/filecoin-project/filecoin-ffi"
	proof5 "github.com/bitchina-io/specs-actors/v5/actors/runtime/proof"
)

var ProofProver = proofProver{}

var _ Prover = ProofProver

type proofProver struct{}

func (v proofProver) AggregateSealProofs(aggregateInfo proof5.AggregateSealVerifyProofAndInfos, proofs [][]byte) ([]byte, error) {
	return ffi.AggregateSealProofs(aggregateInfo, proofs)
}
