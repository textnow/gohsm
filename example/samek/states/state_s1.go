package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
	"go.uber.org/zap"
)

// S1State represents State S1
type S1State struct {
	hsm.BaseState
}

// NewS1State constructor
func NewS1State(logger *zap.Logger, parentState hsm.State) *S1State {
	hsm.Precondition(logger, parentState != nil, fmt.Sprintf("NewS1State: parentState cannot be nil"))

	state := &S1State{
		BaseState: *hsm.NewBaseState("S1", parentState, logger),
	}

	return state
}

// OnEnter enters this state
func (s *S1State) OnEnter(event hsm.Event) hsm.State {
	s.VerifyNotEntered()
	s.Logger().Debug("->S1;")
	return NewS11State(s.Logger(), s).OnEnter(event)
}

// OnExit enters this state
func (s *S1State) OnExit(event hsm.Event) hsm.State {
	s.VerifyNotExited()
	s.Logger().Debug("<-S1;")
	return s.ParentState()
}

// EventHandler returns a transition associated with the event or NilTransition if the event is not handled
func (s *S1State) EventHandler(event hsm.Event) hsm.Transition {
	switch event.ID() {
	case ea.ID():
		return hsm.NewExternalTransition(event, NewS1State(s.Logger(), s.ParentState()), hsm.NopAction)
	case eb.ID():
		return hsm.NewExternalTransition(event, NewS11State(s.Logger(), s), hsm.NopAction)
	case ec.ID():
		return hsm.NewExternalTransition(event, NewS2State(s.Logger(), s.ParentState()), hsm.NopAction)
	case ed.ID():
		return hsm.NewExternalTransition(event, NewS0State(s.Logger()), hsm.NopAction)
	case ef.ID():
		return hsm.NewExternalTransition(event, NewS2State(s.Logger(), s.ParentState()), hsm.NopAction)
	default:
		return nil
	}
}
