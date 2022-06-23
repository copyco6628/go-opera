package iblockproc

import (
	"github.com/copyco6628/lachesis-base/hash"
	"github.com/copyco6628/lachesis-base/inter/idx"
	"github.com/copyco6628/lachesis-base/inter/pos"

	"github.com/copyco6628/go-opera/inter"
	"github.com/copyco6628/go-opera/opera"
)

type ValidatorEpochStateV0 struct {
	GasRefund      uint64
	PrevEpochEvent hash.Event
}

type EpochStateV0 struct {
	Epoch          idx.Epoch
	EpochStart     inter.Timestamp
	PrevEpochStart inter.Timestamp

	EpochStateRoot hash.Hash

	Validators        *pos.Validators
	ValidatorStates   []ValidatorEpochStateV0
	ValidatorProfiles ValidatorProfiles

	Rules opera.Rules
}
