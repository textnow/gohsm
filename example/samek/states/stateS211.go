package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
)

type S211State struct {
	stateEngine       *hsm.StateEngine
	parentStateEngine *hsm.StateEngine
}

func NewS211State(parentState *S21State) *S211State {
	state := &S211State{
		stateEngine:       nil,
		parentStateEngine: nil,
	}

	if parentState != nil {
		parentStateEngine := parentState.GetStateEngine()
		state.stateEngine = hsm.NewStateEngine(state, parentStateEngine)
		state.parentStateEngine = parentStateEngine
	} else {
		state.stateEngine = hsm.NewStateEngine(state, nil)
	}

	return state
}

func (s *S211State) Name() string {
	return "S211"
}

func (s *S211State) OnEnter(event hsm.Event) *hsm.StateEngine {
	fmt.Printf("->S211;")
	return s.stateEngine
}

func (s *S211State) OnExit(event hsm.Event) *hsm.StateEngine {
	fmt.Printf("<-S211;")
	return s.stateEngine.GetParentStateEngine()
}

func (s *S211State) GetEventHandler(event hsm.Event) *hsm.EventHandler {
	switch event.ID() {
	case ed.ID():
		if s.parentStateEngine != nil {
			transition := hsm.NewExternalTransition(event, s.parentStateEngine, hsm.NopAction)
			return hsm.NewEventHandler(transition)
		} else {
			toState := NewS21State(nil)
			transition := hsm.NewExternalTransition(event, toState.GetStateEngine(), hsm.NopAction)
			return hsm.NewEventHandler(transition)
		}
	case eg.ID():
		toState := NewS0State()
		transition := hsm.NewExternalTransition(event, toState.GetStateEngine(), hsm.NopAction)
		return hsm.NewEventHandler(transition)
	default:
		return nil
	}
}

func (s *S211State) GetStateEngine() *hsm.StateEngine {
	return s.stateEngine
}
