package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
	"go.uber.org/zap"
)

// S11State represents State S11
type S11State struct {
	hsm.BaseState
}

// NewS11State constructor
func NewS11State(logger *zap.Logger, parentState *S1State) *S11State {
	hsm.Precondition(logger, parentState != nil, fmt.Sprintf("NewS11State: parentState cannot be nil"))

	state := &S11State{
		BaseState: *hsm.NewBaseState("S11", parentState, logger),
	}

	return state
}

// OnEnter enters this state
func (s *S11State) OnEnter(event hsm.Event) hsm.State {
	s.VerifyNotEntered()
	s.Logger().Debug("->S11;")
	return s
}

// OnExit enters this state
func (s *S11State) OnExit(event hsm.Event) hsm.State {
	s.VerifyNotExited()
	s.Logger().Debug("<-S11;")
	return s.ParentState()
}

// EventHandler returns a transition associated with the event or NilTransition if the event is not handled
func (s *S11State) EventHandler(event hsm.Event) hsm.Transition {
	switch event.ID() {
	case eg.ID():
		return hsm.NewExternalTransition(event, NewS2State(s.Logger(), s.ParentState().ParentState()), hsm.NopAction)
	default:
		return nil
	}
}
