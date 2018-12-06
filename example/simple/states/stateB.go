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

	state.stateEngine = hsm.NewStateEngine(state, parentState.StateEngine())

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
	return s.stateEngine.ParentStateEngine()
}

func (s *StateB) EventHandler(event hsm.Event) *hsm.EventHandler {
	if event.ID() != ea.ID() {
		return nil
	}

	stateC := NewStateC(s.parentState)
	transition := hsm.NewExternalTransition(event, stateC.StateEngine(), action1)
	return hsm.NewEventHandler(transition)
}

func (s *StateB) StateEngine() *hsm.StateEngine {
	return s.stateEngine
}

func action1() {
	fmt.Printf("\nAction1\n")
}
