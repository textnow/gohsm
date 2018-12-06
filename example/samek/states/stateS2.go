package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
)

type S2State struct {
	stateEngine *hsm.StateEngine
	parentState *S0State
}

func NewS2State(parentState *S0State) *S2State {
	state := &S2State{
		stateEngine: nil,
		parentState: parentState,
	}

	state.stateEngine = hsm.NewStateEngine(state, parentState.StateEngine())

	return state
}

func (s *S2State) Name() string {
	return "S2"
}

func (s *S2State) OnEnter(event hsm.Event) *hsm.StateEngine {
	fmt.Printf("->S2;")

	stateS21 := NewS21State(s)

	return stateS21.OnEnter(event)
}

func (s *S2State) OnExit(event hsm.Event) *hsm.StateEngine {
	fmt.Printf("<-S2;")
	return s.stateEngine.ParentStateEngine()
}

func (s *S2State) EventHandler(event hsm.Event) *hsm.EventHandler {
	switch event.ID() {
	case ec.ID():
		toState := NewS1State(s.parentState)
		transition := hsm.NewExternalTransition(event, toState.StateEngine(), hsm.NopAction)
		return hsm.NewEventHandler(transition)
	case ef.ID():
		toState := NewS11State(nil)
		transition := hsm.NewExternalTransition(event, toState.StateEngine(), hsm.NopAction)
		return hsm.NewEventHandler(transition)
	default:
		return nil
	}
}

func (s *S2State) StateEngine() *hsm.StateEngine {
	return s.stateEngine
}
