package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
)

type StateA struct {
	srv		*SimpleService
	a       bool
	entered bool
	exited  bool
}

func NewStateA(srv *SimpleService, a bool) *StateA {
	return &StateA{
		srv: srv,
		a: a,
	}
}

func (s *StateA) Name() string {
	return "A"
}

func (s *StateA) OnEnter(event hsm.Event) hsm.State {
	hsm.Precondition(s.srv, !s.entered, fmt.Sprintf("State %s has already been entered", s.Name()))
	s.srv.Logger().Debug("->A;")
	s.entered = true

	s.srv.Logger().Debug(fmt.Sprintf("Got original test value: %s", s.srv.GetTest()))
	s.srv.SetTest("This value was set in state A OnEnter()")

	if s.a {
		return NewStateB(s.srv, s).OnEnter(event)
	} else {
		return NewStateC(s.srv, s).OnEnter(event)
	}
}

func (s *StateA) OnExit(event hsm.Event) hsm.State {
	hsm.Precondition(s.srv, !s.exited, fmt.Sprintf("State %s has already been entered", s.Name()))
	s.srv.Logger().Debug("<-A;")
	s.exited = true
	return s.ParentState()
}

func (s *StateA) EventHandler(event hsm.Event) hsm.Transition {
	switch event.ID() {
	case ec.ID():
		return hsm.NewExternalTransition(event, NewStateA(s.srv, s.a), action3)
	case eb.ID():
		return hsm.NewInternalTransition(event, action2)
	case ed.ID():
		return hsm.NewExternalTransition(event, NewStateD(s.srv), action4)
	default:
		return hsm.NilTransition
	}
}

func (s *StateA) Entered() bool {
	return s.entered
}

func (s *StateA) Exited() bool {
	return s.exited
}

func (s *StateA) ParentState() hsm.State {
	return hsm.NilState
}

func action2(srv hsm.Service) {
	srv.Logger().Debug("Action2")
	LastActionIdExecuted = 2
}

func action3(srv hsm.Service) {
	srv.Logger().Debug("Action3")
	LastActionIdExecuted = 3
}

func action4(srv hsm.Service) {
	srv.Logger().Debug("Action4")
	LastActionIdExecuted = 4
}
