package states

import (
	"github.com/Enflick/gohsm"
	"go.uber.org/zap"
)

// StateA represents State A
type StateA struct {
	hsm.BaseState
	a bool
}

// NewStateA constructor
func NewStateA(logger *zap.Logger, a bool) *StateA {
	return &StateA{
		BaseState: *hsm.NewBaseState("A", nil, logger),
		a:         a,
	}
}

// OnEnter enters this state and returns the new current state
func (s *StateA) OnEnter(event hsm.Event) hsm.State {
	s.VerifyNotEntered()
	s.Logger().Debug("->A;")

	if s.a {
		return NewStateB(s.Logger(), s).OnEnter(event)
	}
	return NewStateC(s.Logger(), s).OnEnter(event)
}

// OnExit exits this state and returns the parentState or nil if this state does not have a parent
func (s *StateA) OnExit(event hsm.Event) hsm.State {
	s.VerifyNotExited()
	s.Logger().Debug("<-A;")
	return s.ParentState()
}

// EventHandler returns the transition associated with the event or nil if this state does not handle the event
func (s *StateA) EventHandler(event hsm.Event) hsm.Transition {
	switch event.ID() {
	case ec.ID():
		return hsm.NewExternalTransition(event, NewStateA(s.Logger(), s.a), action3)
	case eb.ID():
		return hsm.NewInternalTransition(event, action2)
	case ed.ID():
		return hsm.NewExternalTransition(event, NewStateD(s.Logger()), action4)
	default:
		return nil
	}
}

func action2(logger *zap.Logger) {
	logger.Debug("Action2")
	LastActionIDExecuted = 2
}

func action3(logger *zap.Logger) {
	logger.Debug("Action3")
	LastActionIDExecuted = 3
}

func action4(logger *zap.Logger) {
	logger.Debug("Action4")
	LastActionIDExecuted = 4
}
