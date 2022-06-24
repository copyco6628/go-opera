package concurrent

import (
	"sync"

	"github.com/copyco6628/lachesis-base/inter/idx"
)

type ValidatorBlocksSet struct {
	sync.RWMutex
	Val map[idx.ValidatorID]idx.Block
}

func WrapValidatorBlocksSet(v map[idx.ValidatorID]idx.Block) *ValidatorBlocksSet {
	return &ValidatorBlocksSet{
		RWMutex: sync.RWMutex{},
		Val:     v,
	}
}
