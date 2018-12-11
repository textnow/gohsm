package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
)

type S21State struct {
	parentState *S2State
	entered     bool
	exited      bool
}

func NewS21State(srv hsm.Service, parentState *S2State) *S21State {
	hsm.Precondition(srv, parentState != nil, fmt.Sprintf("NewS21State: parentState cannot be nil"))

	state := &S21State{
		parentState: parentState,
	}

	return state
}

func (s *S21State) Name() string {
	return "S21"
}

func (s *S21State) OnEnter(srv hsm.Service, event hsm.Event) hsm.State {
	hsm.Precondition(srv, !s.entered, fmt.Sprintf("State %s has already been entered", s.Name()))
	srv.Logger().Debug("->S21;")
	s.entered = true

	stateS211 := NewS211State(srv, s)

	return stateS211.OnEnter(srv, event)
}

func (s *S21State) OnExit(srv hsm.Service, event hsm.Event) hsm.State {
	hsm.Precondition(srv, !s.exited, fmt.Sprintf("State %s has already been exited", s.Name()))
	srv.Logger().Debug("<-S21;")
	s.exited = true
	return s.ParentState()
}

func (s *S21State) EventHandler(srv hsm.Service, event hsm.Event) hsm.Transition {
	switch event.ID() {
	case eb.ID():
		return hsm.NewExternalTransition(event, NewS211State(srv, s), hsm.NopAction)
	case eh.ID():
		return hsm.NewExternalTransition(event, NewS21State(srv, s.parentState), hsm.NopAction)
	default:
		return hsm.NilTransition
	}
}

func (s *S21State) ParentState() hsm.State {
	return s.parentState
}

func (s *S21State) Entered() bool {
	return s.entered
}

func (s *S21State) Exited() bool {
	return s.exited
}
