package hsm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewExternalTransition(t *testing.T) {
	mockState := NewMockState(nil)
	mockEvent := NewMockEvent(mockStartEventId)

	transition := NewExternalTransition(mockEvent, mockState.StateEngine(), NopAction)
	assert.NotNil(t, transition)
}

func TestExternalTransition_Execute(t *testing.T) {
	mockState := NewMockState(nil)
	mockEvent := NewMockEvent(mockStartEventId)

	transition := NewExternalTransition(mockEvent, mockState.StateEngine(), NopAction)
	assert.Equal(t, mockState.StateEngine(), transition.Execute(mockState.StateEngine()))
}

func TestNewInternalTransition(t *testing.T) {
	mockEvent := NewMockEvent(mockStartEventId)

	transition := NewInternalTransition(mockEvent, NopAction)
	assert.NotNil(t, transition)
}

func TestInternalTransition_Execute(t *testing.T) {
	mockState := NewMockState(nil)
	mockEvent := NewMockEvent(mockStartEventId)

	transition := NewInternalTransition(mockEvent, NopAction)
	assert.Equal(t, mockState.StateEngine(), transition.Execute(mockState.StateEngine()))
}

func TestNewEndTransition(t *testing.T) {
	mockEvent := NewMockEvent(mockStartEventId)

	transition := NewEndTransition(mockEvent, NopAction)
	assert.NotNil(t, transition)
}

func TestEndTransition_Execute(t *testing.T) {
	mockState := NewMockState(nil)
	mockEvent := NewMockEvent(mockStartEventId)

	transition := NewEndTransition(mockEvent, NopAction)
	assert.Nil(t, transition.Execute(mockState.StateEngine()))
}
