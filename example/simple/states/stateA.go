package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
)

type StateA struct {
	stateEngine *hsm.StateEngine
	a           bool
}

func NewStateA(a bool) *StateA {
	state := &StateA{
		stateEngine: nil,
		a:           a,
	}

	state.stateEngine = hsm.NewStateEngine(state, nil)

	return state
}

func (s *StateA) Name() string {
	return "A"
}

func (s *StateA) OnEnter(event hsm.Event) *hsm.StateEngine {
	fmt.Printf("->A;")

	if s.a {
		stateB := NewStateB(s)
		return stateB.StateEngine().OnEnter(event)
	} else {
		stateC := NewStateC(s)
		return stateC.StateEngine().OnEnter(event)
	}
}

func (s *StateA) OnExit(event hsm.Event) *hsm.StateEngine {
	fmt.Printf("<-A;")
	return s.stateEngine.ParentStateEngine()
}

func (s *StateA) EventHandler(event hsm.Event) *hsm.EventHandler {
	switch event.ID() {
	case ec.ID():
		toState := NewStateA(s.a)
		transition := hsm.NewExternalTransition(event, toState.StateEngine(), action3)
		return hsm.NewEventHandler(transition)
	case eb.ID():
		transition := hsm.NewInternalTransition(event, action2)
		return hsm.NewEventHandler(transition)
	case ed.ID():
		toState := NewStateD()
		transition := hsm.NewExternalTransition(event, toState.StateEngine(), action4)
		return hsm.NewEventHandler(transition)
	default:
		return nil
	}
}

func (s *StateA) StateEngine() *hsm.StateEngine {
	return s.stateEngine
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
