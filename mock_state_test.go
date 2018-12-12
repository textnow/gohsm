package hsm

import (
	"go.uber.org/zap/zaptest"
	"testing"
)

// MockState mocks a State for testing
type MockState struct {
	parentState State
	entered     bool
	exited      bool
}

// NewMockState constructor
func NewMockState(parentState State) *MockState {
	state := &MockState{}

	state.parentState = parentState

	return state
}

var mockStateName = "MockState"

// Name gets the state's name
func (s *MockState) Name() string {
	return mockStateName
}

// OnEnter enters the state
func (s *MockState) OnEnter(event Event) State {
	s.entered = true
	return s
}

// OnExit exits the state
func (s *MockState) OnExit(event Event) State {
	s.exited = true
	return s.ParentState()
}

// EventHandler returns a transtion if the event is handled; otherwise NilTransition is returned
func (s *MockState) EventHandler(event Event) Transition {
	switch event.ID() {
	case "start":
		return NewInternalTransition(event, NopAction)
	case "end":
		return NewEndTransition(event, NopAction)
	}

	return NilTransition
}

// ParentState gets this state's parentState
func (s *MockState) ParentState() State {
	return s.parentState
}

// Entered returns true if this state has been entered
func (s *MockState) Entered() bool {
	return s.entered
}

// Exited returns true if this state has been exited
func (s *MockState) Exited() bool {
	return s.exited
}

func getStateMachine(t *testing.T, startState State) *StateMachine {
	logger := zaptest.NewLogger(t)
	stateMachine := NewStateMachine(logger, startState, StartEvent)

	return stateMachine
}
