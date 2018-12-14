package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
	"go.uber.org/zap"
)

// StateC represents State C
type StateC struct {
	logger      *zap.Logger
	parentState *StateA
	entered     bool
	exited      bool
}

// NewStateC constructor
func NewStateC(logger *zap.Logger, parentState *StateA) *StateC {
	hsm.Precondition(logger, parentState != nil, fmt.Sprintf("NewStateC: parentState cannot be nil"))

	return &StateC{
		logger:      logger,
		parentState: parentState,
	}
}

// Name returns this state's name
func (s *StateC) Name() string {
	return "C"
}

// OnEnter enters this state and returns the new current state
func (s *StateC) OnEnter(event hsm.Event) hsm.State {
	hsm.Precondition(s.logger, !s.entered, fmt.Sprintf("State %s has already been entered", s.Name()))
	s.logger.Debug("->C;")
	s.entered = true
	return s
}

// OnExit exits this state and returns the parentState or NilParentState
func (s *StateC) OnExit(event hsm.Event) hsm.State {
	hsm.Precondition(s.logger, !s.exited, fmt.Sprintf("State %s has already been exited", s.Name()))
	s.logger.Debug("<-C;")
	s.exited = true
	return s.ParentState()
}

// EventHandler returns the transition associated with the event or NilTransition if this state does not handle the event
func (s *StateC) EventHandler(event hsm.Event) hsm.Transition {
	switch event.ID() {
	case ex.ID():
		return hsm.NewExternalTransition(event, NewStateC(s.logger, s.parentState), action6)
	case ey.ID():
		return hsm.NewInternalTransition(event, action7)
	default:
		return hsm.NilTransition
	}
}

// Entered returns true if this state has been entered
func (s *StateC) Entered() bool {
	return s.entered
}

// Exited returns true if this state has been exited
func (s *StateC) Exited() bool {
	return s.exited
}

// ParentState returns this state's parentState or NilState if the state does not have a parent
func (s *StateC) ParentState() hsm.State {
	return s.parentState
}

func action6(logger *zap.Logger) {
	logger.Debug("Action6")
	LastActionIDExecuted = 6
}

func action7(logger *zap.Logger) {
	logger.Debug("Action7")
	LastActionIDExecuted = 7
}
