package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
)

type StateB struct {
	parentState *StateA
	entered     bool
	exited      bool
}

func NewStateB(parentState *StateA) *StateB {
	hsm.Precondition(parentState != nil, fmt.Sprintf("NewStateB: parentState cannot be nil"))

	return &StateB{
		parentState: parentState,
	}
}

func (s *StateB) Name() string {
	return "B"
}

func (s *StateB) OnEnter(ctx hsm.Context, event hsm.Event) hsm.State {
	hsm.Precondition(!s.entered, fmt.Sprintf("State %s has already been entered", s.Name()))
	ctx.Logger().Debug("->B;")
	s.entered = true
	return s
}

func (s *StateB) OnExit(ctx hsm.Context, event hsm.Event) hsm.State {
	hsm.Precondition(!s.exited, fmt.Sprintf("State %s has already been entered", s.Name()))
	ctx.Logger().Debug("<-B;")
	s.exited = true
	return s.ParentState()
}

func (s *StateB) EventHandler(c hsm.Context, event hsm.Event) hsm.Transition {
	if event.ID() != ea.ID() {
		return hsm.NilTransition
	}

	return hsm.NewExternalTransition(event, NewStateC(s.parentState), action1)
}

func (s *StateB) Entered() bool {
	return s.entered
}

func (s *StateB) Exited() bool {
	return s.exited
}

func (s *StateB) ParentState() hsm.State {
	return s.parentState
}

func action1(ctx hsm.Context) {
	ctx.Logger().Debug("Action1")
	LastActionIdExecuted = 1
}
