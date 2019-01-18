package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
	"go.uber.org/zap"
)

// S21State represents State S21
type S21State struct {
	hsm.BaseState
}

// NewS21State constructor
func NewS21State(logger *zap.Logger, parentState hsm.State) *S21State {
	hsm.Precondition(logger, parentState != nil, fmt.Sprintf("NewS21State: parentState cannot be nil"))

	state := &S21State{
		BaseState: *hsm.NewBaseState("S21", parentState, logger),
	}

	return state
}

// OnEnter enters this state
func (s *S21State) OnEnter(event hsm.Event) hsm.State {
	s.VerifyNotEntered()
	s.Logger().Debug("->S21;")

	return NewS211State(s.Logger(), s).OnEnter(event)
}

// OnExit enters this state
func (s *S21State) OnExit(event hsm.Event) hsm.State {
	s.VerifyNotExited()
	s.Logger().Debug("<-S21;")
	return s.ParentState()
}

// EventHandler returns a transition associated with the event or NilTransition if the event is not handled
func (s *S21State) EventHandler(event hsm.Event) hsm.Transition {
	switch event.ID() {
	case eb.ID():
		return hsm.NewExternalTransition(event, NewS211State(s.Logger(), s), hsm.NopAction)
	case eh.ID():
		return hsm.NewExternalTransition(event, NewS21State(s.Logger(), s.ParentState()), hsm.NopAction)
	default:
		return nil
	}
}
