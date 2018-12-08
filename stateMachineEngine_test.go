package hsm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStateMachineEngine(t *testing.T) {
	startState := NewMockState(nil)
	stateMachineEngine := getStateMachine(t, startState)
	assert.Equal(t, startState.StateEngine(), stateMachineEngine.currentStateEngine)
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
		startState := NewMockState(nil)
		stateMachineEngine := getStateMachine(t, startState)

		event := NewMockEvent(tt.eventId)
		result := stateMachineEngine.HandleEvent(event)

		assert.Equal(t, result, tt.result)
		if tt.currentStateName != "" {
			assert.NotNil(t, stateMachineEngine.currentStateEngine)
			assert.Equal(t, stateMachineEngine.currentStateEngine.Name(), tt.currentStateName)
		} else {
			assert.Nil(t, stateMachineEngine.currentStateEngine)
		}
	}
}
