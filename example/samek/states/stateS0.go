package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
)

type S0State struct {
	entered bool
	exited  bool
}

func NewS0State() *S0State {
	return &S0State{}
}

func (s *S0State) Name() string {
	return "S0"
}

func (s *S0State) OnEnter(event hsm.Event) hsm.State {
	hsm.Precondition(!s.entered, fmt.Sprintf("State %s has already been entered", s.Name()))
	fmt.Printf("->S0;")
	s.entered = true

	stateS1 := NewS1State(s)

	return stateS1.OnEnter(event)
}

func (s *S0State) OnExit(event hsm.Event) hsm.State {
	hsm.Precondition(!s.exited, fmt.Sprintf("State %s has already been exited", s.Name()))
	fmt.Printf("<-S0;")
	s.exited = true
	return s.ParentState()
}

func (s *S0State) EventHandler(event hsm.Event) hsm.Transition {
	switch event.ID() {
	case ee.ID():
		return hsm.NewExternalTransition(event, NewS2State(s), hsm.NopAction)
	default:
		return hsm.NilTransition
	}
}

func (s *S0State) ParentState() hsm.State {
	return hsm.NilState
}

func (s *S0State) Entered() bool {
	return s.entered
}

func (s *S0State) Exited() bool {
	return s.exited
}
