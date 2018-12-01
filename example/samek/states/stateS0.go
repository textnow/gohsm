package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
)

type S0State struct {
	stateEngine *hsm.StateEngine
}

func NewS0State() *S0State {
	state := &S0State{
		stateEngine: nil,
	}

	state.stateEngine = hsm.NewStateEngine(state, nil)

	return state
}

func (s *S0State) Name() string {
	return "S0"
}

func (s *S0State) OnEnter(event hsm.Event) *hsm.StateEngine {
	fmt.Printf("->S0;")

	stateS1 := NewS1State(s)

	return stateS1.OnEnter(event)
}

func (s *S0State) OnExit(event hsm.Event) *hsm.StateEngine {
	fmt.Printf("<-S0;")
	return s.stateEngine.GetParentStateEngine()
}

func (s *S0State) GetEventHandler(event hsm.Event) *hsm.EventHandler {
	switch event.ID() {
	case ee.ID():
		toState := NewS211State(nil)
		transition := hsm.NewExternalTransition(event, toState.GetStateEngine(), hsm.NopAction)
		return hsm.NewEventHandler(transition)
	default:
		return nil
	}
}

func (s *S0State) GetStateEngine() *hsm.StateEngine {
	return s.stateEngine
}
