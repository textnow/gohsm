package states

import (
	"github.com/Enflick/gohsm"
	"go.uber.org/zap"
)

// StateD represents State D
type StateD struct {
	hsm.BaseState
}

// NewStateD constructor
func NewStateD(logger *zap.Logger) *StateD {
	return &StateD{
		BaseState: *hsm.NewBaseState("D", nil, logger),
	}
}

// OnEnter enters this state and returns the new current state
func (s *StateD) OnEnter(event hsm.Event) hsm.State {
	s.VerifyNotEntered()
	s.Logger().Debug("->D;")
	return s
}

// OnExit exits this state and returns the parentState or nil if this state does not have a parent
func (s *StateD) OnExit(event hsm.Event) hsm.State {
	s.VerifyNotExited()
	s.Logger().Debug("<-D;")
	return s.ParentState()
}

// EventHandler returns the transition associated with the event or nil if this state does not handle the event
func (s *StateD) EventHandler(event hsm.Event) hsm.Transition {
	switch event.ID() {
	case ee.ID():
		return hsm.NewEndTransition(event, action5)
	default:
		return nil
	}
}

func action5(logger *zap.Logger) {
	logger.Debug("Action5")
	LastActionIDExecuted = 5
}
