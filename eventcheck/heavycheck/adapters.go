package heavycheck

import (
	"github.com/copyco6628/lachesis-base/inter/dag"

	"github.com/copyco6628ra/inter"
)

type EventsOnly struct {
	*Checker
}

func (c *EventsOnly) Enqueue(e dag.Event, onValidated func(error)) error {
	return c.Checker.EnqueueEvent(e.(inter.EventPayloadI), onValidated)
}

type BVsOnly struct {
	*Checker
}

func (c *BVsOnly) Enqueue(bvs inter.LlrSignedBlockVotes, onValidated func(error)) error {
	return c.Checker.EnqueueBVs(bvs, onValidated)
}

type EVOnly struct {
	*Checker
}

func (c *EVOnly) Enqueue(ers inter.LlrSignedEpochVote, onValidated func(error)) error {
	return c.Checker.EnqueueEV(ers, onValidated)
}
