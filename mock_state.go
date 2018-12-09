package hsm

import (
	"go.uber.org/zap/zaptest"
	"testing"
)

type MockState struct {
	parentState State
	entered     bool
	exited      bool
}

func NewMockState(parentState State) *MockState {
	state := &MockState{}

	state.parentState = parentState

	return state
}

var mockStateName = "MockState"

func (s *MockState) Name() string {
	return mockStateName
}

func (s *MockState) OnEnter(event Event) State {
	s.entered = true
	return s
}

func (s *MockState) OnExit(event Event) State {
	s.exited = true
	return s.ParentState()
}

func (s *MockState) EventHandler(event Event) Transition {
	switch event.ID() {
	case "start":
		return NewInternalTransition(event, NopAction)
	case "end":
		return NewEndTransition(event, NopAction)
	}

	return NilTransition
}

func (s *MockState) ParentState() State {
	return s.parentState
}

func (s *MockState) Entered() bool {
	return s.entered
}

func (s *MockState) Exited() bool {
	return s.exited
}

func getStateMachine(t *testing.T, startState State) *StateMachine {
	logger := zaptest.NewLogger(t)
	stateMachine := NewStateMachine(logger, startState)

	return stateMachine
}
