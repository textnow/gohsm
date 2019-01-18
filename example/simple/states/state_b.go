package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
	"go.uber.org/zap"
)

// StateB represents State B
type StateB struct {
	hsm.BaseState
}

// NewStateB constructor
func NewStateB(logger *zap.Logger, parentState *StateA) *StateB {
	hsm.Precondition(logger, parentState != nil, fmt.Sprintf("NewStateB: parentState cannot be nil"))

	return &StateB{
		BaseState: *hsm.NewBaseState("B", parentState, logger),
	}
}

// OnEnter enters this state and returns the new current state
func (s *StateB) OnEnter(event hsm.Event) hsm.State {
	s.VerifyNotEntered()
	s.Logger().Debug("->B;")
	return s
}

// OnExit exits this state and returns the parentState or nil if this state does not have a parent
func (s *StateB) OnExit(event hsm.Event) hsm.State {
	s.VerifyNotExited()
	s.Logger().Debug("<-B;")
	return s.ParentState()
}

// EventHandler returns the transition associated with the event or nil if this state does not handle the event
func (s *StateB) EventHandler(event hsm.Event) hsm.Transition {
	if event.ID() != ea.ID() {
		return nil
	}

	return hsm.NewExternalTransition(event, NewStateC(s.Logger(), s.ParentState()), action1)
}

func action1(logger *zap.Logger) {
	logger.Debug("Action1")
	LastActionIDExecuted = 1
}
