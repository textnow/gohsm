package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
)

type S11State struct {
	stateEngine       *hsm.StateEngine
	parentStateEngine *hsm.StateEngine
}

func NewS11State(parentState *S1State) *S11State {
	state := &S11State{
		stateEngine:       nil,
		parentStateEngine: nil,
	}

	if parentState == nil {
		state.stateEngine = hsm.NewStateEngine(state, nil)
	} else {
		state.stateEngine = hsm.NewStateEngine(state, parentState.GetStateEngine())
		state.parentStateEngine = parentState.GetStateEngine()
	}

	return state
}

func (s *S11State) Name() string {
	return "S11"
}

func (s *S11State) OnEnter(event hsm.Event) *hsm.StateEngine {
	fmt.Printf("->S11;")
	return s.stateEngine
}

func (s *S11State) OnExit(event hsm.Event) *hsm.StateEngine {
	fmt.Printf("<-S11;")
	return s.stateEngine.GetParentStateEngine()
}

func (s *S11State) GetEventHandler(event hsm.Event) *hsm.EventHandler {
	switch event.ID() {
	case eg.ID():
		toState := NewS211State(nil)
		transition := hsm.NewExternalTransition(event, toState.GetStateEngine(), hsm.NopAction)
		return hsm.NewEventHandler(transition)
	default:
		return nil
	}
}

func (s *S11State) GetStateEngine() *hsm.StateEngine {
	return s.stateEngine
}
