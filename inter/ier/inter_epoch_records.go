package ier

import (
	"github.com/copyco6628/lachesis-base/hash"
	"github.com/copyco6628/lachesis-base/inter/idx"

	"github.com/copyco6628/go-opera/inter/iblockproc"
)

type LlrFullEpochRecord struct {
	BlockState iblockproc.BlockState
	EpochState iblockproc.EpochState
}

type LlrIdxFullEpochRecord struct {
	LlrFullEpochRecord
	Idx idx.Epoch
}

func (er LlrFullEpochRecord) Hash() hash.Hash {
	return hash.Of(er.BlockState.Hash().Bytes(), er.EpochState.Hash().Bytes())
}
