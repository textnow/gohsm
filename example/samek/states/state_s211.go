package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
	"go.uber.org/zap"
)

// S211State represents State S211
type S211State struct {
	logger      *zap.Logger
	parentState *S21State
	entered     bool
	exited      bool
}

// NewS211State constructor
func NewS211State(logger *zap.Logger, parentState *S21State) *S211State {
	hsm.Precondition(logger, parentState != nil, fmt.Sprintf("NewS211State: parentState cannot be nil"))

	state := &S211State{
		logger:      logger,
		parentState: parentState,
	}

	return state
}

// Name returns the state's name
func (s *S211State) Name() string {
	return "S211"
}

// OnEnter enters this state
func (s *S211State) OnEnter(event hsm.Event) hsm.State {
	hsm.Precondition(s.logger, !s.entered, fmt.Sprintf("State %s has already been entered", s.Name()))
	s.logger.Debug("->S211;")
	s.entered = true
	return s
}

// OnExit enters this state
func (s *S211State) OnExit(event hsm.Event) hsm.State {
	hsm.Precondition(s.logger, !s.exited, fmt.Sprintf("State %s has already been exited", s.Name()))
	s.logger.Debug("<-S211;")
	s.exited = true
	return s.ParentState()
}

// EventHandler returns a transition associated with the event or NilTransition if the event is not handled
func (s *S211State) EventHandler(event hsm.Event) hsm.Transition {
	switch event.ID() {
	case ed.ID():
		return hsm.NewExternalTransition(event, NewS21State(s.logger, s.parentState.parentState), hsm.NopAction)
	case eg.ID():
		return hsm.NewExternalTransition(event, NewS0State(s.logger), hsm.NopAction)
	default:
		return hsm.NilTransition
	}
}

// ParentState returns the parentState or NilState
func (s *S211State) ParentState() hsm.State {
	return s.parentState
}

// Entered returns true if this state have been entered
func (s *S211State) Entered() bool {
	return s.entered
}

// Exited returns true if this state have been exited
func (s *S211State) Exited() bool {
	return s.exited
}
