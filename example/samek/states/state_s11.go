package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
	"go.uber.org/zap"
)

// S11State represents State S11
type S11State struct {
	logger      *zap.Logger
	parentState *S1State
	entered     bool
	exited      bool
}

// NewS11State constructor
func NewS11State(logger *zap.Logger, parentState *S1State) *S11State {
	hsm.Precondition(logger, parentState != nil, fmt.Sprintf("NewS11State: parentState cannot be nil"))

	state := &S11State{
		logger:      logger,
		parentState: parentState,
	}

	return state
}

// Name returns the state's name
func (s *S11State) Name() string {
	return "S11"
}

// OnEnter enters this state
func (s *S11State) OnEnter(event hsm.Event) hsm.State {
	hsm.Precondition(s.logger, !s.entered, fmt.Sprintf("State %s has already been entered", s.Name()))
	s.logger.Debug("->S11;")
	s.entered = true
	return s
}

// OnExit enters this state
func (s *S11State) OnExit(event hsm.Event) hsm.State {
	hsm.Precondition(s.logger, !s.exited, fmt.Sprintf("State %s has already been exited", s.Name()))
	s.logger.Debug("<-S11;")
	s.exited = true
	return s.ParentState()
}

// EventHandler returns a transition associated with the event or NilTransition if the event is not handled
func (s *S11State) EventHandler(event hsm.Event) hsm.Transition {
	switch event.ID() {
	case eg.ID():
		return hsm.NewExternalTransition(event, NewS2State(s.logger, s.parentState.parentState), hsm.NopAction)
	default:
		return hsm.NilTransition
	}
}

// ParentState returns the parentState or NilState
func (s *S11State) ParentState() hsm.State {
	return s.parentState
}

// Entered returns true if this state have been entered
func (s *S11State) Entered() bool {
	return s.entered
}

// Exited returns true if this state have been exited
func (s *S11State) Exited() bool {
	return s.exited
}
