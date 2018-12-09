package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
)

type S21State struct {
	parentState *S2State
	entered     bool
	exited      bool
}

func NewS21State(parentState *S2State) *S21State {
	hsm.Precondition(parentState != nil, fmt.Sprintf("NewS21State: parentState cannot be nil"))

	state := &S21State{
		parentState: parentState,
	}

	return state
}

func (s *S21State) Name() string {
	return "S21"
}

func (s *S21State) OnEnter(event hsm.Event) hsm.State {
	hsm.Precondition(!s.entered, fmt.Sprintf("State %s has already been entered", s.Name()))
	fmt.Printf("->S21;")
	s.entered = true

	stateS211 := NewS211State(s)

	return stateS211.OnEnter(event)
}

func (s *S21State) OnExit(event hsm.Event) hsm.State {
	hsm.Precondition(!s.exited, fmt.Sprintf("State %s has already been exited", s.Name()))
	fmt.Printf("<-S21;")
	s.exited = true
	return s.ParentState()
}

func (s *S21State) EventHandler(event hsm.Event) hsm.Transition {
	switch event.ID() {
	case eb.ID():
		return hsm.NewExternalTransition(event, NewS211State(s), hsm.NopAction)
	case eh.ID():
		return hsm.NewExternalTransition(event, NewS21State(s.parentState), hsm.NopAction)
	default:
		return hsm.NilTransition
	}
}

func (s *S21State) ParentState() hsm.State {
	return s.parentState
}

func (s *S21State) Entered() bool {
	return s.entered
}

func (s *S21State) Exited() bool {
	return s.exited
}
