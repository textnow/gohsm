package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
)

type S2State struct {
	parentState *S0State
	entered     bool
	exited      bool
}

func NewS2State(srv hsm.Service, parentState *S0State) *S2State {
	hsm.Precondition(srv, parentState != nil, fmt.Sprintf("NewS2State: parentState cannot be nil"))

	state := &S2State{
		parentState: parentState,
	}

	return state
}

func (s *S2State) Name() string {
	return "S2"
}

func (s *S2State) OnEnter(srv hsm.Service, event hsm.Event) hsm.State {
	hsm.Precondition(srv, !s.entered, fmt.Sprintf("State %s has already been entered", s.Name()))
	srv.Logger().Debug("->S2;")
	s.entered = true

	stateS21 := NewS21State(srv, s)

	return stateS21.OnEnter(srv, event)
}

func (s *S2State) OnExit(srv hsm.Service, event hsm.Event) hsm.State {
	hsm.Precondition(srv, !s.exited, fmt.Sprintf("State %s has already been exited", s.Name()))
	srv.Logger().Debug("<-S2;")
	s.exited = true
	return s.ParentState()
}

func (s *S2State) EventHandler(srv hsm.Service, event hsm.Event) hsm.Transition {
	switch event.ID() {
	case ec.ID():
		return hsm.NewExternalTransition(event, NewS1State(srv, s.parentState), hsm.NopAction)
	case ef.ID():
		return hsm.NewExternalTransition(event, NewS1State(srv, s.parentState), hsm.NopAction)
	default:
		return hsm.NilTransition
	}
}

func (s *S2State) ParentState() hsm.State {
	return s.parentState
}

func (s *S2State) Entered() bool {
	return s.entered
}

func (s *S2State) Exited() bool {
	return s.exited
}
