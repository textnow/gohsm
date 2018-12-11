package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
)

type S11State struct {
	parentState *S1State
	entered     bool
	exited      bool
}

func NewS11State(srv hsm.Service, parentState *S1State) *S11State {
	hsm.Precondition(srv, parentState != nil, fmt.Sprintf("NewS11State: parentState cannot be nil"))

	state := &S11State{
		parentState: parentState,
	}

	return state
}

func (s *S11State) Name() string {
	return "S11"
}

func (s *S11State) OnEnter(srv hsm.Service, event hsm.Event) hsm.State {
	hsm.Precondition(srv, !s.entered, fmt.Sprintf("State %s has already been entered", s.Name()))
	srv.Logger().Debug("->S11;")
	s.entered = true
	return s
}

func (s *S11State) OnExit(srv hsm.Service, event hsm.Event) hsm.State {
	hsm.Precondition(srv, !s.exited, fmt.Sprintf("State %s has already been exited", s.Name()))
	srv.Logger().Debug("<-S11;")
	s.exited = true
	return s.ParentState()
}

func (s *S11State) EventHandler(srv hsm.Service, event hsm.Event) hsm.Transition {
	switch event.ID() {
	case eg.ID():
		return hsm.NewExternalTransition(event, NewS2State(srv, s.parentState.parentState), hsm.NopAction)
	default:
		return hsm.NilTransition
	}
}

func (s *S11State) ParentState() hsm.State {
	return s.parentState
}

func (s *S11State) Entered() bool {
	return s.entered
}

func (s *S11State) Exited() bool {
	return s.exited
}
