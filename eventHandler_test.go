package hsm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEventHandler(t *testing.T) {
	mockEvent := NewMockEvent(mockStartEventId)
	transition := NewInternalTransition(mockEvent, NopAction)

	eventHandler := NewEventHandler(transition)
	assert.NotNil(t, eventHandler)
}

func TestExecute(t *testing.T) {
	mockEvent := NewMockEvent(mockStartEventId)
	transition := NewInternalTransition(mockEvent, NopAction)
	mockFromState := NewMockState(nil)

	eventHandler := NewEventHandler(transition)
	stateEngine := eventHandler.Execute(mockFromState.StateEngine())
	assert.Equal(t, stateEngine, mockFromState.StateEngine())
}
