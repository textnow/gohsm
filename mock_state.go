package hsm

import (
	"go.uber.org/zap/zaptest"
	"testing"
)

type MockState struct {
	stateEngine *StateEngine
}

func NewMockState(parentStateEngine *StateEngine) *MockState {
	state := &MockState{
		stateEngine: nil,
	}

	state.stateEngine = NewStateEngine(state, parentStateEngine)

	return state
}

func (s *MockState) Initialize(state *StateEngine) {
	s.stateEngine = state
}

var mockStateName = "MockState"

func (s *MockState) Name() string {
	return mockStateName
}

func (s *MockState) OnEnter(event Event) *StateEngine {
	return s.stateEngine
}

func (s *MockState) OnExit(event Event) *StateEngine {
	return s.stateEngine.ParentStateEngine()
}

func (s *MockState) EventHandler(event Event) *EventHandler {
	switch event.ID() {
	case "start":
		transition := NewInternalTransition(event, NopAction)
		return NewEventHandler(transition)
	case "end":
		transition := NewEndTransition(event, NopAction)
		return NewEventHandler(transition)
	}

	return nil
}

func (s *MockState) StateEngine() *StateEngine {
	return s.stateEngine
}

func getStateMachine(t *testing.T, startState State) *StateMachineEngine {
	logger := zaptest.NewLogger(t)
	stateMachineEngine := NewStateMachineEngine(logger, startState)

	return stateMachineEngine
}
