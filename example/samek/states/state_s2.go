package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
	"go.uber.org/zap"
)

// S2State represents State S2
type S2State struct {
	hsm.BaseState
}

// NewS2State constructor
func NewS2State(logger *zap.Logger, parentState hsm.State) *S2State {
	hsm.Precondition(logger, parentState != nil, fmt.Sprintf("NewS2State: parentState cannot be nil"))

	state := &S2State{
		BaseState: *hsm.NewBaseState("S2", parentState, logger),
	}

	return state
}

// OnEnter enters this state
func (s *S2State) OnEnter(event hsm.Event) hsm.State {
	s.VerifyNotEntered()
	s.Logger().Debug("->S2;")
	return NewS21State(s.Logger(), s).OnEnter(event)
}

// OnExit exits this state
func (s *S2State) OnExit(event hsm.Event) hsm.State {
	s.VerifyNotExited()
	s.Logger().Debug("<-S2;")
	return s.ParentState()
}

// EventHandler returns a transition associated with the event or NilTransition if the event is not handled
func (s *S2State) EventHandler(event hsm.Event) hsm.Transition {
	switch event.ID() {
	case ec.ID():
		return hsm.NewExternalTransition(event, NewS1State(s.Logger(), s.ParentState()), hsm.NopAction)
	case ef.ID():
		return hsm.NewExternalTransition(event, NewS1State(s.Logger(), s.ParentState()), hsm.NopAction)
	default:
		return nil
	}
}
