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

func (s *StateA) OnEnter(event hsm.Event) hsm.State {
	hsm.Precondition(!s.entered, fmt.Sprintf("State %s has already been entered", s.Name()))
	fmt.Printf("->A;")
	s.entered = true

	if s.a {
		return NewStateB(s).OnEnter(event)
	} else {
		return NewStateC(s).OnEnter(event)
	}
}

func (s *StateA) OnExit(event hsm.Event) hsm.State {
	hsm.Precondition(!s.exited, fmt.Sprintf("State %s has already been entered", s.Name()))
	fmt.Printf("<-A;")
	s.exited = true
	return s.ParentState()
}

func (s *StateA) EventHandler(event hsm.Event) hsm.Transition {
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

func action2() {
	fmt.Printf("\nAction2\n")
	LastActionIdExecuted = 2
}

func action3() {
	fmt.Printf("\nAction3\n")
	LastActionIdExecuted = 3
}

func action4() {
	fmt.Printf("\nAction4\n")
	LastActionIdExecuted = 4
}
