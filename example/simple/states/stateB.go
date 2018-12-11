package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
)

type StateB struct {
	srv			*SimpleService
	parentState *StateA
	entered     bool
	exited      bool
}

func NewStateB(srv *SimpleService, parentState *StateA) *StateB {
	hsm.Precondition(srv, parentState != nil, fmt.Sprintf("NewStateB: parentState cannot be nil"))

	return &StateB{
		srv: srv,
		parentState: parentState,
	}
}

func (s *StateB) Name() string {
	return "B"
}

func (s *StateB) OnEnter(event hsm.Event) hsm.State {
	hsm.Precondition(s.srv, !s.entered, fmt.Sprintf("State %s has already been entered", s.Name()))
	s.srv.Logger().Debug("->B;")
	s.srv.Logger().Debug(fmt.Sprintf("Got test value in state B: %s", s.srv.GetTest()))
	s.entered = true
	return s
}

func (s *StateB) OnExit(event hsm.Event) hsm.State {
	hsm.Precondition(s.srv, !s.exited, fmt.Sprintf("State %s has already been entered", s.Name()))
	s.srv.Logger().Debug("<-B;")
	s.exited = true
	return s.ParentState()
}

func (s *StateB) EventHandler(event hsm.Event) hsm.Transition {
	if event.ID() != ea.ID() {
		return hsm.NilTransition
	}

	return hsm.NewExternalTransition(event, NewStateC(s.srv, s.parentState), action1)
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
