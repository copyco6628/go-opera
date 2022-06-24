package iep

import (
	"github.com/copyco6628/go-opera/inter"
	"github.com/copyco6628/go-opera/inter/ier"
)

type LlrEpochPack struct {
	Votes  []inter.LlrSignedEpochVote
	Record ier.LlrIdxFullEpochRecord
}
