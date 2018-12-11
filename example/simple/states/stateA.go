package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
)

type StateA struct {
	a       bool
	entered bool
	exited  bool
}

func NewStateA(a bool) *StateA {
	return &StateA{
		a: a,
	}
}

func (s *StateA) Name() string {
	return "A"
}

func (s *StateA) OnEnter(srv hsm.Service, event hsm.Event) hsm.State {
	sc := ToSimpleService(srv)

	hsm.Precondition(srv, !s.entered, fmt.Sprintf("State %s has already been entered", s.Name()))
	sc.Logger().Debug("->A;")
	s.entered = true

	sc.Logger().Debug(fmt.Sprintf("Got original test value: %s", sc.GetTest()))
	sc.SetTest("This value was set in state A OnEnter()")

	if s.a {
		return NewStateB(srv, s).OnEnter(srv, event)
	} else {
		return NewStateC(srv, s).OnEnter(srv, event)
	}
}

func (s *StateA) OnExit(srv hsm.Service, event hsm.Event) hsm.State {
	hsm.Precondition(srv, !s.exited, fmt.Sprintf("State %s has already been entered", s.Name()))
	srv.Logger().Debug("<-A;")
	s.exited = true
	return s.ParentState()
}

func (s *StateA) EventHandler(srv hsm.Service, event hsm.Event) hsm.Transition {
	switch event.ID() {
	case ec.ID():
		return hsm.NewExternalTransition(event, NewStateA(s.a), action3)
	case eb.ID():
		return hsm.NewInternalTransition(event, action2)
	case ed.ID():
		return hsm.NewExternalTransition(event, NewStateD(), action4)
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
