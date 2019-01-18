package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
	"go.uber.org/zap"
)

// S211State represents State S211
type S211State struct {
	hsm.BaseState
}

// NewS211State constructor
func NewS211State(logger *zap.Logger, parentState hsm.State) *S211State {
	hsm.Precondition(logger, parentState != nil, fmt.Sprintf("NewS211State: parentState cannot be nil"))

	state := &S211State{
		BaseState: *hsm.NewBaseState("S211", parentState, logger),
	}

	return state
}

// OnEnter enters this state
func (s *S211State) OnEnter(event hsm.Event) hsm.State {
	s.VerifyNotEntered()
	s.Logger().Debug("->S211;")
	return s
}

// OnExit enters this state
func (s *S211State) OnExit(event hsm.Event) hsm.State {
	s.VerifyNotExited()
	s.Logger().Debug("<-S211;")
	return s.ParentState()
}

// EventHandler returns a transition associated with the event or NilTransition if the event is not handled
func (s *S211State) EventHandler(event hsm.Event) hsm.Transition {
	switch event.ID() {
	case ed.ID():
		return hsm.NewExternalTransition(event, NewS21State(s.Logger(), s.ParentState().ParentState()), hsm.NopAction)
	case eg.ID():
		return hsm.NewExternalTransition(event, NewS0State(s.Logger()), hsm.NopAction)
	default:
		return nil
	}
}
