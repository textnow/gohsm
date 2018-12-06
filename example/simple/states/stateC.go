package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
)

type StateC struct {
	stateEngine *hsm.StateEngine
}

func NewStateC(parentState *StateA) *StateC {
	state := &StateC{
		stateEngine: nil,
	}

	state.stateEngine = hsm.NewStateEngine(state, parentState.StateEngine())

	return state
}

func (s *StateC) Initialize(state *hsm.StateEngine) {
	s.stateEngine = state
}

func (s *StateC) Name() string {
	return "C"
}

func (s *StateC) OnEnter(event hsm.Event) *hsm.StateEngine {
	fmt.Printf("->C;")
	return s.stateEngine
}

func (s *StateC) OnExit(event hsm.Event) *hsm.StateEngine {
	fmt.Printf("<-C;")
	return s.stateEngine.ParentStateEngine()
}

func (s *StateC) EventHandler(event hsm.Event) *hsm.EventHandler {
	switch event.ID() {
	case ex.ID():
		transition := hsm.NewExternalTransition(event, s.stateEngine, action6)
		return hsm.NewEventHandler(transition)
	case ey.ID():
		transition := hsm.NewInternalTransition(event, action7)
		return hsm.NewEventHandler(transition)
	default:
		return nil
	}
}

func (s *StateC) StateEngine() *hsm.StateEngine {
	return s.stateEngine
}

func action6() {
	fmt.Printf("\nAction6\n")
}

func action7() {
	fmt.Printf("\nAction7\n")
}
