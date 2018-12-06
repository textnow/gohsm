package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
)

type S1State struct {
	stateEngine *hsm.StateEngine
	parentState *S0State
}

func NewS1State(parentState *S0State) *S1State {
	state := &S1State{
		stateEngine: nil,
		parentState: parentState,
	}

	state.stateEngine = hsm.NewStateEngine(state, parentState.StateEngine())

	return state
}

func (s *S1State) Name() string {
	return "S1"
}

func (s *S1State) OnEnter(event hsm.Event) *hsm.StateEngine {
	fmt.Printf("->S1;")

	stateS1 := NewS11State(s)

	return stateS1.OnEnter(event)
}

func (s *S1State) OnExit(event hsm.Event) *hsm.StateEngine {
	fmt.Printf("<-S1;")
	return s.stateEngine.ParentStateEngine()
}

func (s *S1State) EventHandler(event hsm.Event) *hsm.EventHandler {
	switch event.ID() {
	case ea.ID():
		transition := hsm.NewExternalTransition(event, s.stateEngine, hsm.NopAction)
		return hsm.NewEventHandler(transition)
	case eb.ID():
		toState := NewS11State(s)
		transition := hsm.NewExternalTransition(event, toState.StateEngine(), hsm.NopAction)
		return hsm.NewEventHandler(transition)
	case ec.ID():
		toState := NewS2State(s.parentState)
		transition := hsm.NewExternalTransition(event, toState.StateEngine(), hsm.NopAction)
		return hsm.NewEventHandler(transition)
	case ed.ID():
		toState := NewS0State()
		transition := hsm.NewExternalTransition(event, toState.StateEngine(), hsm.NopAction)
		return hsm.NewEventHandler(transition)
	case ef.ID():
		toState := NewS211State(nil)
		transition := hsm.NewExternalTransition(event, toState.StateEngine(), hsm.NopAction)
		return hsm.NewEventHandler(transition)
	default:
		return nil
	}
}

func (s *S1State) StateEngine() *hsm.StateEngine {
	return s.stateEngine
}
