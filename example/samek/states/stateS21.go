package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
)

type S21State struct {
	stateEngine       *hsm.StateEngine
	parentStateEngine *hsm.StateEngine
}

func NewS21State(parentState *S2State) *S21State {
	state := &S21State{
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

func (s *S21State) Name() string {
	return "S21"
}

func (s *S21State) OnEnter(event hsm.Event) *hsm.StateEngine {
	fmt.Printf("->S21;")

	stateS211 := NewS211State(s)

	return stateS211.OnEnter(event)
}

func (s *S21State) OnExit(event hsm.Event) *hsm.StateEngine {
	fmt.Printf("<-S21;")
	return s.stateEngine.GetParentStateEngine()
}

func (s *S21State) GetEventHandler(event hsm.Event) *hsm.EventHandler {
	switch event.ID() {
	case eb.ID():
		toState := NewS211State(s)
		transition := hsm.NewExternalTransition(event, toState.GetStateEngine(), hsm.NopAction)
		return hsm.NewEventHandler(transition)
	case eh.ID():
		transition := hsm.NewExternalTransition(event, s.stateEngine, hsm.NopAction)
		return hsm.NewEventHandler(transition)
	default:
		return nil
	}
}

func (s *S21State) GetStateEngine() *hsm.StateEngine {
	return s.stateEngine
}
