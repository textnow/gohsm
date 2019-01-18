package states

import (
	"github.com/Enflick/gohsm"
	"go.uber.org/zap"
)

// S0State represents State S0
type S0State struct {
	hsm.BaseState
}

// NewS0State constructor
func NewS0State(logger *zap.Logger) *S0State {
	return &S0State{
		BaseState: *hsm.NewBaseState("S0", nil, logger),
	}
}

// OnEnter enters this state
func (s *S0State) OnEnter(event hsm.Event) hsm.State {
	s.VerifyNotEntered()
	s.Logger().Debug("->S0;")
	return NewS1State(s.Logger(), s).OnEnter(event)
}

// OnExit enters this state
func (s *S0State) OnExit(event hsm.Event) hsm.State {
	s.VerifyNotExited()
	s.Logger().Debug("<-S0;")
	return s.ParentState()
}

// EventHandler returns a transition associated with the event or NilTransition if the event is not handled
func (s *S0State) EventHandler(event hsm.Event) hsm.Transition {
	switch event.ID() {
	case ee.ID():
		return hsm.NewExternalTransition(event, NewS2State(s.Logger(), s), hsm.NopAction)
	default:
		return nil
	}
}
