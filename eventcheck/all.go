package eventcheck

import (
	"github.com/copyco6628/go-opera/eventcheck/basiccheck"
	"github.com/copyco6628ra/eventcheck/epochcheck"
	"github.com/copyco6628ra/eventcheck/gaspowercheck"
	"github.com/copyco6628ra/eventcheck/heavycheck"
	"github.com/copyco6628ra/eventcheck/parentscheck"
	"github.com/copyco6628ra/inter"
)

// Checkers is collection of all the checkers
type Checkers struct {
	Basiccheck    *basiccheck.Checker
	Epochcheck    *epochcheck.Checker
	Parentscheck  *parentscheck.Checker
	Gaspowercheck *gaspowercheck.Checker
	Heavycheck    *heavycheck.Checker
}

// Validate runs all the checks except Poset-related
func (v *Checkers) Validate(e inter.EventPayloadI, parents inter.EventIs) error {
	if err := v.Basiccheck.Validate(e); err != nil {
		return err
	}
	if err := v.Epochcheck.Validate(e); err != nil {
		return err
	}
	if err := v.Parentscheck.Validate(e, parents); err != nil {
		return err
	}
	var selfParent inter.EventI
	if e.SelfParent() != nil {
		selfParent = parents[0]
	}
	if err := v.Gaspowercheck.Validate(e, selfParent); err != nil {
		return err
	}
	if err := v.Heavycheck.ValidateEvent(e); err != nil {
		return err
	}
	return nil
}
