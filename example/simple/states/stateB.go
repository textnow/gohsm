package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
)

type StateB struct {
	stateEngine *hsm.StateEngine
	parentState *StateA
}

func NewStateB(parentState *StateA) *StateB {
	state := &StateB{
		stateEngine: nil,
		parentState: parentState,
	}

	state.stateEngine = hsm.NewStateEngine(state, parentState.GetStateEngine())

	return state
}

func (s *StateB) Name() string {
	return "B"
}

func (s *StateB) OnEnter(event hsm.Event) *hsm.StateEngine {
	fmt.Printf("->B;")
	return s.stateEngine
}

func (s *StateB) OnExit(event hsm.Event) *hsm.StateEngine {
	fmt.Printf("<-B;")
	return s.stateEngine.GetParentStateEngine()
}

func (s *StateB) GetEventHandler(event hsm.Event) *hsm.EventHandler {
	if event.ID() != ea.ID() {
		return nil
	}

	stateC := NewStateC(s.parentState)
	transition := hsm.NewExternalTransition(event, stateC.GetStateEngine(), daveAction1)
	return hsm.NewEventHandler(transition)
}

func (s *StateB) GetStateEngine() *hsm.StateEngine {
	return s.stateEngine
}

func daveAction1() {
	fmt.Printf("\nAction1\n")
}
