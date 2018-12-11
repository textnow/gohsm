package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
)

type S211State struct {
	srv 		hsm.Service
	parentState *S21State
	entered     bool
	exited      bool
}

func NewS211State(srv hsm.Service, parentState *S21State) *S211State {
	hsm.Precondition(srv, parentState != nil, fmt.Sprintf("NewS211State: parentState cannot be nil"))

	state := &S211State{
		srv: srv,
		parentState: parentState,
	}

	return state
}

func (s *S211State) Name() string {
	return "S211"
}

func (s *S211State) OnEnter(event hsm.Event) hsm.State {
	hsm.Precondition(s.srv, !s.entered, fmt.Sprintf("State %s has already been entered", s.Name()))
	s.srv.Logger().Debug("->S211;")
	s.entered = true
	return s
}

func (s *S211State) OnExit(event hsm.Event) hsm.State {
	hsm.Precondition(s.srv, !s.exited, fmt.Sprintf("State %s has already been exited", s.Name()))
	s.srv.Logger().Debug("<-S211;")
	s.exited = true
	return s.ParentState()
}

func (s *S211State) EventHandler(event hsm.Event) hsm.Transition {
	switch event.ID() {
	case ed.ID():
		return hsm.NewExternalTransition(event, NewS21State(s.srv, s.parentState.parentState), hsm.NopAction)
	case eg.ID():
		return hsm.NewExternalTransition(event, NewS0State(s.srv), hsm.NopAction)
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
