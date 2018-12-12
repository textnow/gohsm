package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
	"go.uber.org/zap"
)

// StateA represents State A
type StateB struct {
	logger      *zap.Logger
	parentState *StateA
	entered     bool
	exited      bool
}

// NewStateB constructor
func NewStateB(logger *zap.Logger, parentState *StateA) *StateB {
	hsm.Precondition(logger, parentState != nil, fmt.Sprintf("NewStateB: parentState cannot be nil"))

	return &StateB{
		logger:      logger,
		parentState: parentState,
	}
}

// Name returns this state's name
func (s *StateB) Name() string {
	return "B"
}

// OnEnter enters this state and returns the new current state
func (s *StateB) OnEnter(event hsm.Event) hsm.State {
	hsm.Precondition(s.logger, !s.entered, fmt.Sprintf("State %s has already been entered", s.Name()))
	s.logger.Debug("->B;")
	s.entered = true
	return s
}

// OnExit exits this state and returns the parentState or NilParentState
func (s *StateB) OnExit(event hsm.Event) hsm.State {
	hsm.Precondition(s.logger, !s.exited, fmt.Sprintf("State %s has already been exited", s.Name()))
	s.logger.Debug("<-B;")
	s.exited = true
	return s.ParentState()
}

// EventHandler returns the transition associated with the event or NilTransition if this state does not handle the event
func (s *StateB) EventHandler(event hsm.Event) hsm.Transition {
	if event.ID() != ea.ID() {
		return hsm.NilTransition
	}

	return hsm.NewExternalTransition(event, NewStateC(s.logger, s.parentState), action1)
}

// Entered returns true if this state has been entered
func (s *StateB) Entered() bool {
	return s.entered
}

// Exited returns true if this state has been exited
func (s *StateB) Exited() bool {
	return s.exited
}

// ParentState returns this state's parentState or NilState if the state does not have a parent
func (s *StateB) ParentState() hsm.State {
	return s.parentState
}

func action1(logger *zap.Logger) {
	logger.Debug("Action1")
	LastActionIdExecuted = 1
}
