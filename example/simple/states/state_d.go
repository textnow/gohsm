package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
	"go.uber.org/zap"
)

// StateD represents State D
type StateD struct {
	logger  *zap.Logger
	entered bool
	exited  bool
}

// NewStateD constructor
func NewStateD(logger *zap.Logger) *StateD {
	return &StateD{
		logger: logger,
	}
}

// Name returns this state's name
func (s *StateD) Name() string {
	return "D"
}

// OnEnter enters this state and returns the new current state
func (s *StateD) OnEnter(event hsm.Event) hsm.State {
	hsm.Precondition(s.logger, !s.entered, fmt.Sprintf("State %s has already been entered", s.Name()))
	s.logger.Debug("->D;")
	s.entered = true
	return s
}

// OnExit exits this state and returns the parentState or NilParentState
func (s *StateD) OnExit(event hsm.Event) hsm.State {
	hsm.Precondition(s.logger, !s.exited, fmt.Sprintf("State %s has already been exited", s.Name()))
	s.logger.Debug("<-D;")
	s.exited = true
	return s.ParentState()
}

// EventHandler returns the transition associated with the event or NilTransition if this state does not handle the event
func (s *StateD) EventHandler(event hsm.Event) hsm.Transition {
	switch event.ID() {
	case ee.ID():
		return hsm.NewEndTransition(event, action5)
	default:
		return hsm.NilTransition
	}
}

// Entered returns true if this state has been entered
func (s *StateD) Entered() bool {
	return s.entered
}

// Exited returns true if this state has been exited
func (s *StateD) Exited() bool {
	return s.exited
}

// ParentState returns this state's parentState or NilState if the state does not have a parent
func (s *StateD) ParentState() hsm.State {
	return hsm.NilState
}

func action5(logger *zap.Logger) {
	logger.Debug("Action5")
	LastActionIDExecuted = 5
}
