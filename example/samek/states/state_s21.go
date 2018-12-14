package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
	"go.uber.org/zap"
)

// S21State represents State S21
type S21State struct {
	logger      *zap.Logger
	parentState *S2State
	entered     bool
	exited      bool
}

// NewS21State constructor
func NewS21State(logger *zap.Logger, parentState *S2State) *S21State {
	hsm.Precondition(logger, parentState != nil, fmt.Sprintf("NewS21State: parentState cannot be nil"))

	state := &S21State{
		logger:      logger,
		parentState: parentState,
	}

	return state
}

// Name returns the state's name
func (s *S21State) Name() string {
	return "S21"
}

// OnEnter enters this state
func (s *S21State) OnEnter(event hsm.Event) hsm.State {
	hsm.Precondition(s.logger, !s.entered, fmt.Sprintf("State %s has already been entered", s.Name()))
	s.logger.Debug("->S21;")
	s.entered = true

	stateS211 := NewS211State(s.logger, s)

	return stateS211.OnEnter(event)
}

// OnExit enters this state
func (s *S21State) OnExit(event hsm.Event) hsm.State {
	hsm.Precondition(s.logger, !s.exited, fmt.Sprintf("State %s has already been exited", s.Name()))
	s.logger.Debug("<-S21;")
	s.exited = true
	return s.ParentState()
}

// EventHandler returns a transition associated with the event or NilTransition if the event is not handled
func (s *S21State) EventHandler(event hsm.Event) hsm.Transition {
	switch event.ID() {
	case eb.ID():
		return hsm.NewExternalTransition(event, NewS211State(s.logger, s), hsm.NopAction)
	case eh.ID():
		return hsm.NewExternalTransition(event, NewS21State(s.logger, s.parentState), hsm.NopAction)
	default:
		return hsm.NilTransition
	}
}

// ParentState returns the parentState or NilState
func (s *S21State) ParentState() hsm.State {
	return s.parentState
}

// Entered returns true if this state have been entered
func (s *S21State) Entered() bool {
	return s.entered
}

// Exited returns true if this state have been exited
func (s *S21State) Exited() bool {
	return s.exited
}
