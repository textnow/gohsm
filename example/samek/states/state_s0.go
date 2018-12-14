package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
	"go.uber.org/zap"
)

// S0State represents State S0
type S0State struct {
	logger  *zap.Logger
	entered bool
	exited  bool
}

// NewS0State constructor
func NewS0State(logger *zap.Logger) *S0State {
	return &S0State{
		logger: logger,
	}
}

// Name returns the state's name
func (s *S0State) Name() string {
	return "S0"
}

// OnEnter enters this state
func (s *S0State) OnEnter(event hsm.Event) hsm.State {
	hsm.Precondition(s.logger, !s.entered, fmt.Sprintf("State %s has already been entered", s.Name()))
	s.logger.Debug("->S0;")
	s.entered = true

	stateS1 := NewS1State(s.logger, s)

	return stateS1.OnEnter(event)
}

// OnExit enters this state
func (s *S0State) OnExit(event hsm.Event) hsm.State {
	hsm.Precondition(s.logger, !s.exited, fmt.Sprintf("State %s has already been exited", s.Name()))
	s.logger.Debug("<-S0;")
	s.exited = true
	return s.ParentState()
}

// EventHandler returns a transition associated with the event or NilTransition if the event is not handled
func (s *S0State) EventHandler(event hsm.Event) hsm.Transition {
	switch event.ID() {
	case ee.ID():
		return hsm.NewExternalTransition(event, NewS2State(s.logger, s), hsm.NopAction)
	default:
		return hsm.NilTransition
	}
}

// ParentState returns the parentState or NilState
func (s *S0State) ParentState() hsm.State {
	return hsm.NilState
}

// Entered returns true if this state have been entered
func (s *S0State) Entered() bool {
	return s.entered
}

// Exited returns true if this state have been exited
func (s *S0State) Exited() bool {
	return s.exited
}
