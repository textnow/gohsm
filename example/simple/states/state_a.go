package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
	"go.uber.org/zap"
)

// StateA represents State A
type StateA struct {
	logger  *zap.Logger
	a       bool
	entered bool
	exited  bool
}

// NewStateA constructor
func NewStateA(logger *zap.Logger, a bool) *StateA {
	return &StateA{
		logger: logger,
		a:      a,
	}
}

// Name returns this state's name
func (s *StateA) Name() string {
	return "A"
}

// OnEnter enters this state and returns the new current state
func (s *StateA) OnEnter(event hsm.Event) hsm.State {
	hsm.Precondition(s.logger, !s.entered, fmt.Sprintf("State %s has already been entered", s.Name()))
	s.logger.Debug("->A;")
	s.entered = true

	if s.a {
		return NewStateB(s.logger, s).OnEnter(event)
	} else {
		return NewStateC(s.logger, s).OnEnter(event)
	}
}

// OnExit exits this state and returns the parentState or NilParentState
func (s *StateA) OnExit(event hsm.Event) hsm.State {
	hsm.Precondition(s.logger, !s.exited, fmt.Sprintf("State %s has already been exited", s.Name()))
	s.logger.Debug("<-A;")
	s.exited = true
	return s.ParentState()
}

// EventHandler returns the transition associated with the event or NilTransition if this state does not handle the event
func (s *StateA) EventHandler(event hsm.Event) hsm.Transition {
	switch event.ID() {
	case ec.ID():
		return hsm.NewExternalTransition(event, NewStateA(s.logger, s.a), action3)
	case eb.ID():
		return hsm.NewInternalTransition(event, action2)
	case ed.ID():
		return hsm.NewExternalTransition(event, NewStateD(s.logger), action4)
	default:
		return hsm.NilTransition
	}
}

// Entered returns true if this state has been entered
func (s *StateA) Entered() bool {
	return s.entered
}

// Exited returns true if this state has been exited
func (s *StateA) Exited() bool {
	return s.exited
}

// ParentState returns this state's parentState or NilState if the state does not have a parent
func (s *StateA) ParentState() hsm.State {
	return hsm.NilState
}

func action2(logger *zap.Logger) {
	logger.Debug("Action2")
	LastActionIdExecuted = 2
}

func action3(logger *zap.Logger) {
	logger.Debug("Action3")
	LastActionIdExecuted = 3
}

func action4(logger *zap.Logger) {
	logger.Debug("Action4")
	LastActionIdExecuted = 4
}
