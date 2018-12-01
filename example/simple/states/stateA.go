package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
)

type StateA struct {
	stateEngine *hsm.StateEngine
	a           bool
}

func NewStateA() *StateA {
	state := &StateA{
		stateEngine: nil,
		a:           true,
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
		return stateB.OnEnter(event)
	} else {
		stateC := NewStateC(s)
		return stateC.OnEnter(event)
	}
}

func (s *StateA) OnExit(event hsm.Event) *hsm.StateEngine {
	fmt.Printf("<-A;")
	return s.stateEngine.GetParentStateEngine()
}

func (s *StateA) GetEventHandler(event hsm.Event) *hsm.EventHandler {
	switch event.ID() {
	case ec.ID():
		transition := hsm.NewExternalTransition(event, s.stateEngine, daveAction3)
		return hsm.NewEventHandler(transition)
	case eb.ID():
		transition := hsm.NewInternalTransition(event, daveAction2)
		return hsm.NewEventHandler(transition)
	case ed.ID():
		toState := NewStateD()
		transition := hsm.NewExternalTransition(event, toState.GetStateEngine(), daveAction4)
		return hsm.NewEventHandler(transition)
	default:
		return nil
	}
}

func (s *StateA) GetStateEngine() *hsm.StateEngine {
	return s.stateEngine
}

func daveAction2() {
	fmt.Printf("\nAction2\n")
}

func daveAction3() {
	fmt.Printf("\nAction3\n")
}

func daveAction4() {
	fmt.Printf("\nAction4\n")
}
