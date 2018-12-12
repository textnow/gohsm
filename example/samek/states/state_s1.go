package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
	"go.uber.org/zap"
)

// S1State represents State S1
type S1State struct {
	logger      *zap.Logger
	parentState *S0State
	entered     bool
	exited      bool
}

// NewS1State constructor
func NewS1State(logger *zap.Logger, parentState *S0State) *S1State {
	hsm.Precondition(logger, parentState != nil, fmt.Sprintf("NewS1State: parentState cannot be nil"))

	state := &S1State{
		logger:      logger,
		parentState: parentState,
	}

	return state
}

// Name returns the state's name
func (s *S1State) Name() string {
	return "S1"
}

// OnEnter enters this state
func (s *S1State) OnEnter(event hsm.Event) hsm.State {
	hsm.Precondition(s.logger, !s.entered, fmt.Sprintf("State %s has already been entered", s.Name()))
	s.logger.Debug("->S1;")
	s.entered = true

	stateS1 := NewS11State(s.logger, s)

	return stateS1.OnEnter(event)
}

// OnExit enters this state
func (s *S1State) OnExit(event hsm.Event) hsm.State {
	hsm.Precondition(s.logger, !s.exited, fmt.Sprintf("State %s has already been exited", s.Name()))
	s.logger.Debug("<-S1;")
	s.exited = true
	return s.ParentState()
}

// EventHandler returns a transition associated with the event or NilTransition if the event is not handled
func (s *S1State) EventHandler(event hsm.Event) hsm.Transition {
	switch event.ID() {
	case ea.ID():
		return hsm.NewExternalTransition(event, NewS1State(s.logger, s.parentState), hsm.NopAction)
	case eb.ID():
		return hsm.NewExternalTransition(event, NewS11State(s.logger, s), hsm.NopAction)
	case ec.ID():
		return hsm.NewExternalTransition(event, NewS2State(s.logger, s.parentState), hsm.NopAction)
	case ed.ID():
		return hsm.NewExternalTransition(event, NewS0State(s.logger), hsm.NopAction)
	case ef.ID():
		return hsm.NewExternalTransition(event, NewS2State(s.logger, s.parentState), hsm.NopAction)
	default:
		return hsm.NilTransition
	}
}

// ParentState returns the parentState or NilState
func (s *S1State) ParentState() hsm.State {
	return s.parentState
}

// Entered returns true if this state have been entered
func (s *S1State) Entered() bool {
	return s.entered
}

// Exited returns true if this state have been exited
func (s *S1State) Exited() bool {
	return s.exited
}
