package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
)

type S211State struct {
	parentState *S21State
	entered     bool
	exited      bool
}

func NewS211State(srv hsm.Service, parentState *S21State) *S211State {
	hsm.Precondition(srv, parentState != nil, fmt.Sprintf("NewS211State: parentState cannot be nil"))

	state := &S211State{
		parentState: parentState,
	}

	return state
}

func (s *S211State) Name() string {
	return "S211"
}

func (s *S211State) OnEnter(srv hsm.Service, event hsm.Event) hsm.State {
	hsm.Precondition(srv, !s.entered, fmt.Sprintf("State %s has already been entered", s.Name()))
	srv.Logger().Debug("->S211;")
	s.entered = true
	return s
}

func (s *S211State) OnExit(srv hsm.Service, event hsm.Event) hsm.State {
	hsm.Precondition(srv, !s.exited, fmt.Sprintf("State %s has already been exited", s.Name()))
	srv.Logger().Debug("<-S211;")
	s.exited = true
	return s.ParentState()
}

func (s *S211State) EventHandler(srv hsm.Service, event hsm.Event) hsm.Transition {
	switch event.ID() {
	case ed.ID():
		return hsm.NewExternalTransition(event, NewS21State(srv, s.parentState.parentState), hsm.NopAction)
	case eg.ID():
		return hsm.NewExternalTransition(event, NewS0State(), hsm.NopAction)
	default:
		return hsm.NilTransition
	}
}

func (s *S211State) ParentState() hsm.State {
	return s.parentState
}

func (s *S211State) Entered() bool {
	return s.entered
}

func (s *S211State) Exited() bool {
	return s.exited
}
