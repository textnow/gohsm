package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
)

type StateD struct {
	stateEngine *hsm.StateEngine
}

func NewStateD() *StateD {
	state := &StateD{
		stateEngine: nil,
	}

	state.stateEngine = hsm.NewStateEngine(state, nil)

	return state
}

func (s *StateD) Initialize(state *hsm.StateEngine) {
	s.stateEngine = state
}

func (s *StateD) Name() string {
	return "D"
}

func (s *StateD) OnEnter(event hsm.Event) *hsm.StateEngine {
	fmt.Printf("->D;")
	return s.stateEngine
}

func (s *StateD) OnExit(event hsm.Event) *hsm.StateEngine {
	fmt.Printf("<-D;")
	return s.stateEngine.GetParentStateEngine()
}

func (s *StateD) GetEventHandler(event hsm.Event) *hsm.EventHandler {
	switch event.ID() {
	case ee.ID():
		transition := hsm.NewEndTransition(event, daveAction5)
		return hsm.NewEventHandler(transition)
	default:
		return nil
	}
}

func (s *StateD) GetStateEngine() *hsm.StateEngine {
	return s.stateEngine
}

func daveAction5() {
	fmt.Printf("\nAction5\n")
}
