package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
)

type S1State struct {
	parentState *S0State
	entered     bool
	exited      bool
}

func NewS1State(srv hsm.Service, parentState *S0State) *S1State {
	hsm.Precondition(srv, parentState != nil, fmt.Sprintf("NewS1State: parentState cannot be nil"))

	state := &S1State{
		parentState: parentState,
	}

	return state
}

func (s *S1State) Name() string {
	return "S1"
}

func (s *S1State) OnEnter(srv hsm.Service, event hsm.Event) hsm.State {
	hsm.Precondition(srv, !s.entered, fmt.Sprintf("State %s has already been entered", s.Name()))
	srv.Logger().Debug("->S1;")
	s.entered = true

	stateS1 := NewS11State(srv, s)

	return stateS1.OnEnter(srv, event)
}

func (s *S1State) OnExit(srv hsm.Service, event hsm.Event) hsm.State {
	hsm.Precondition(srv, !s.exited, fmt.Sprintf("State %s has already been exited", s.Name()))
	srv.Logger().Debug("<-S1;")
	s.exited = true
	return s.ParentState()
}

func (s *S1State) EventHandler(srv hsm.Service, event hsm.Event) hsm.Transition {
	switch event.ID() {
	case ea.ID():
		return hsm.NewExternalTransition(event, NewS1State(srv, s.parentState), hsm.NopAction)
	case eb.ID():
		return hsm.NewExternalTransition(event, NewS11State(srv, s), hsm.NopAction)
	case ec.ID():
		return hsm.NewExternalTransition(event, NewS2State(srv, s.parentState), hsm.NopAction)
	case ed.ID():
		return hsm.NewExternalTransition(event, NewS0State(), hsm.NopAction)
	case ef.ID():
		return hsm.NewExternalTransition(event, NewS2State(srv, s.parentState), hsm.NopAction)
	default:
		return hsm.NilTransition
	}
}

func (s *S1State) ParentState() hsm.State {
	return s.parentState
}

func (s *S1State) Entered() bool {
	return s.entered
}

func (s *S1State) Exited() bool {
	return s.exited
}
