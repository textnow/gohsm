package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
)

type StateD struct {
	entered bool
	exited  bool
}

func NewStateD() *StateD {
	return &StateD{}
}

func (s *StateD) Name() string {
	return "D"
}

func (s *StateD) OnEnter(event hsm.Event) hsm.State {
	hsm.Precondition(!s.entered, fmt.Sprintf("State %s has already been entered", s.Name()))
	fmt.Printf("->D;")
	s.entered = true
	return s
}

func (s *StateD) OnExit(event hsm.Event) hsm.State {
	hsm.Precondition(!s.exited, fmt.Sprintf("State %s has already been entered", s.Name()))
	fmt.Printf("<-D;")
	s.exited = true
	return s.ParentState()
}

func (s *StateD) EventHandler(event hsm.Event) hsm.Transition {
	switch event.ID() {
	case ee.ID():
		return hsm.NewEndTransition(event, action5)
	default:
		return hsm.NilTransition
	}
}

func (s *StateD) Entered() bool {
	return s.entered
}

func (s *StateD) Exited() bool {
	return s.exited
}

func (s *StateD) ParentState() hsm.State {
	return hsm.NilState
}

func action5() {
	fmt.Printf("\nAction5\n")
	LastActionIdExecuted = 5
}
