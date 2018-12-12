package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
	"go.uber.org/zap"
)

// S2State represents State S2
type S2State struct {
	logger      *zap.Logger
	parentState *S0State
	entered     bool
	exited      bool
}

// NewS2State constructor
func NewS2State(logger *zap.Logger, parentState *S0State) *S2State {
	hsm.Precondition(logger, parentState != nil, fmt.Sprintf("NewS2State: parentState cannot be nil"))

	state := &S2State{
		logger:      logger,
		parentState: parentState,
	}

	return state
}

// Name returns the state's name
func (s *S2State) Name() string {
	return "S2"
}

// OnEnter enters this state
func (s *S2State) OnEnter(event hsm.Event) hsm.State {
	hsm.Precondition(s.logger, !s.entered, fmt.Sprintf("State %s has already been entered", s.Name()))
	s.logger.Debug("->S2;")
	s.entered = true

	stateS21 := NewS21State(s.logger, s)

	return stateS21.OnEnter(event)
}

// OnExit enters this state
func (s *S2State) OnExit(event hsm.Event) hsm.State {
	hsm.Precondition(s.logger, !s.exited, fmt.Sprintf("State %s has already been exited", s.Name()))
	s.logger.Debug("<-S2;")
	s.exited = true
	return s.ParentState()
}

// EventHandler returns a transition associated with the event or NilTransition if the event is not handled
func (s *S2State) EventHandler(event hsm.Event) hsm.Transition {
	switch event.ID() {
	case ec.ID():
		return hsm.NewExternalTransition(event, NewS1State(s.logger, s.parentState), hsm.NopAction)
	case ef.ID():
		return hsm.NewExternalTransition(event, NewS1State(s.logger, s.parentState), hsm.NopAction)
	default:
		return hsm.NilTransition
	}
}

// ParentState returns the parentState or NilState
func (s *S2State) ParentState() hsm.State {
	return s.parentState
}

// Entered returns true if this state have been entered
func (s *S2State) Entered() bool {
	return s.entered
}

// Exited returns true if this state have been exited
func (s *S2State) Exited() bool {
	return s.exited
}
