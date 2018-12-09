package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
)

type StateC struct {
	parentState *StateA
	entered     bool
	exited      bool
}

func NewStateC(parentState *StateA) *StateC {
	hsm.Precondition(parentState != nil, fmt.Sprintf("NewStateC: parentState cannot be nil"))

	return &StateC{
		parentState: parentState,
	}
}

func (s *StateC) Name() string {
	return "C"
}

func (s *StateC) OnEnter(event hsm.Event) hsm.State {
	hsm.Precondition(!s.entered, fmt.Sprintf("State %s has already been entered", s.Name()))
	fmt.Printf("->C;")
	s.entered = true
	return s
}

func (s *StateC) OnExit(event hsm.Event) hsm.State {
	hsm.Precondition(!s.exited, fmt.Sprintf("State %s has already been entered", s.Name()))
	fmt.Printf("<-C;")
	s.exited = true
	return s.ParentState()
}

func (s *StateC) EventHandler(event hsm.Event) hsm.Transition {
	switch event.ID() {
	case ex.ID():
		return hsm.NewExternalTransition(event, NewStateC(s.parentState), action6)
	case ey.ID():
		return hsm.NewInternalTransition(event, action7)
	default:
		return hsm.NilTransition
	}
}

func (s *StateC) Entered() bool {
	return s.entered
}

func (s *StateC) Exited() bool {
	return s.exited
}

func (s *StateC) ParentState() hsm.State {
	return s.parentState
}

func action6() {
	fmt.Printf("\nAction6\n")
	LastActionIdExecuted = 6
}

func action7() {
	fmt.Printf("\nAction7\n")
	LastActionIdExecuted = 7
}
