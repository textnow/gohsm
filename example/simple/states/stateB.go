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

func NewStateB(srv hsm.Service, parentState *StateA) *StateB {
	hsm.Precondition(srv, parentState != nil, fmt.Sprintf("NewStateB: parentState cannot be nil"))

	return &StateB{
		parentState: parentState,
	}
}

func (s *StateB) Name() string {
	return "B"
}

func (s *StateB) OnEnter(srv hsm.Service, event hsm.Event) hsm.State {
	sc := ToSimpleService(srv)
	hsm.Precondition(srv, !s.entered, fmt.Sprintf("State %s has already been entered", s.Name()))
	sc.Logger().Debug("->B;")
	sc.Logger().Debug(fmt.Sprintf("Got test value in state B: %s", sc.GetTest()))
	s.entered = true
	return s
}

func (s *StateB) OnExit(srv hsm.Service, event hsm.Event) hsm.State {
	hsm.Precondition(srv, !s.exited, fmt.Sprintf("State %s has already been entered", s.Name()))
	srv.Logger().Debug("<-B;")
	s.exited = true
	return s.ParentState()
}

func (s *StateB) EventHandler(srv hsm.Service, event hsm.Event) hsm.Transition {
	if event.ID() != ea.ID() {
		return hsm.NilTransition
	}

	return hsm.NewExternalTransition(event, NewStateC(srv, s.parentState), action1)
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

func action1(srv hsm.Service) {
	srv.Logger().Debug("Action1")
	LastActionIdExecuted = 1
}
