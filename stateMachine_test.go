package hsm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStateMachine(t *testing.T) {
	startState := NewMockState(NilState)
	stateMachine := getStateMachine(t, startState)
	assert.Equal(t, startState, stateMachine.currentState)
}

type handleEventTest struct {
	eventId          string
	result           bool
	currentStateName string
}

var handleEventTests = []handleEventTest{
	{
		eventId:          mockStartEventId,
		result:           true,
		currentStateName: mockStateName,
	},
	{
		eventId:          mockSkipEventId,
		result:           false,
		currentStateName: mockStateName,
	},
	{
		eventId:          mockEndEventId,
		result:           true,
		currentStateName: "",
	},
}

func TestHandleEvent(t *testing.T) {
	for _, tt := range handleEventTests {
		startState := NewMockState(NilState)
		stateMachine := getStateMachine(t, startState)

		event := NewMockEvent(tt.eventId)
		result := stateMachine.HandleEvent(event)

		assert.Equal(t, result, tt.result)
		if tt.currentStateName != "" {
			assert.False(t, IsNilState(stateMachine.currentState))
			assert.Equal(t, stateMachine.currentState.Name(), tt.currentStateName)
		} else {
			assert.True(t, IsNilState(stateMachine.currentState))
		}
	}
}
