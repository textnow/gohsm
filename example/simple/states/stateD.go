package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
)

type StateD struct {
	srv 	*SimpleService
	entered bool
	exited  bool
}

func NewStateD(srv *SimpleService) *StateD {
	return &StateD{
		srv: srv,
	}
}

func (s *StateD) Name() string {
	return "D"
}

func (s *StateD) OnEnter(event hsm.Event) hsm.State {
	hsm.Precondition(s.srv, !s.entered, fmt.Sprintf("State %s has already been entered", s.Name()))
	s.srv.Logger().Debug("->D;")
	s.entered = true
	return s
}

func (s *StateD) OnExit(event hsm.Event) hsm.State {
	hsm.Precondition(s.srv, !s.exited, fmt.Sprintf("State %s has already been entered", s.Name()))
	s.srv.Logger().Debug("<-D;")
	s.exited = true
	return s.ParentState()
}

func (s *StateD) EventHandler(event hsm.Event) hsm.Transition {
	switch event.ID() {
	case ee.ID():
		return hsm.NewEndTransition(event, action5)
	default:
		return hsm.NilTransition
	}
}

func (s *StateD) Entered() bool {
	return s.entered
}

func (s *StateD) Exited() bool {
	return s.exited
}

func (s *StateD) ParentState() hsm.State {
	return hsm.NilState
}

func action5(srv hsm.Service) {
	srv.Logger().Debug("Action5")
	LastActionIdExecuted = 5
}
