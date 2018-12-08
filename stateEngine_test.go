package hsm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	mockState := NewMockState(nil)
	stateEngine := NewStateEngine(mockState, nil)
	assert.Equal(t, stateEngine.Name(), mockStateName)
}

var mockParentState = NewMockState(nil)

type stateEngineParentTest struct {
	parentStateEngine *StateEngine
}

var stateEngineParentTests = []stateEngineParentTest{
	{
		parentStateEngine: mockParentState.StateEngine(),
	},
	{
		parentStateEngine: nil,
	},
}

func TestParentStateEngine(t *testing.T) {
	for _, tt := range stateEngineParentTests {
		mockState := NewMockState(tt.parentStateEngine)
		stateEngine := NewStateEngine(mockState, tt.parentStateEngine)

		// Verify ParentStateEngine
		if tt.parentStateEngine != nil {
			assert.Equal(t, tt.parentStateEngine, stateEngine.ParentStateEngine())
		} else {
			assert.Nil(t, stateEngine.ParentStateEngine())
		}
	}
}

func TestOnEnter(t *testing.T) {
	mockState := NewMockState(nil)
	stateEngine := NewStateEngine(mockState, nil)
	event := NewMockEvent(mockStartEventId)
	assert.Equal(t, stateEngine.OnEnter(event).Name(), stateEngine.Name())
}

func TestOnExit(t *testing.T) {
	for _, tt := range stateEngineParentTests {
		mockState := NewMockState(tt.parentStateEngine)
		stateEngine := NewStateEngine(mockState, nil)
		event := NewMockEvent(mockStartEventId)

		// Verify ParentStateEngine
		if tt.parentStateEngine != nil {
			assert.Equal(t, tt.parentStateEngine, stateEngine.OnExit(event))
		} else {
			assert.Nil(t, stateEngine.ParentStateEngine())
		}
	}
}

type eventHandlerTest struct {
	eventId         string
	nilEventHandler bool
}

var eventHandlerTests = []eventHandlerTest{
	{
		eventId:         mockStartEventId,
		nilEventHandler: false,
	},
	{
		eventId:         mockSkipEventId,
		nilEventHandler: true,
	},
}

func TestStateEngine(t *testing.T) {
	for _, tt := range eventHandlerTests {
		mockState := NewMockState(nil)
		mockEvent := NewMockEvent(tt.eventId)
		stateEngine := NewStateEngine(mockState, nil)

		// Verify EventHandler()
		if tt.nilEventHandler {
			assert.Nil(t, stateEngine.EventHandler(mockEvent))
		} else {
			assert.NotNil(t, stateEngine.EventHandler(mockEvent))
		}
	}
}
