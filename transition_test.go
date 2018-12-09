package hsm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewExternalTransition(t *testing.T) {
	mockState := NewMockState(NilState)
	mockEvent := NewMockEvent(mockStartEventId)

	transition := NewExternalTransition(mockEvent, mockState, NopAction)
	assert.NotNil(t, transition)
}

func TestExternalTransition_Execute(t *testing.T) {
	mockState := NewMockState(NilState)
	mockEvent := NewMockEvent(mockStartEventId)

	transition := NewExternalTransition(mockEvent, mockState, NopAction)
	assert.Equal(t, mockState, transition.Execute(mockState))
}

func TestNewInternalTransition(t *testing.T) {
	mockEvent := NewMockEvent(mockStartEventId)

	transition := NewInternalTransition(mockEvent, NopAction)
	assert.NotNil(t, transition)
}

func TestInternalTransition_Execute(t *testing.T) {
	mockState := NewMockState(NilState)
	mockEvent := NewMockEvent(mockStartEventId)

	transition := NewInternalTransition(mockEvent, NopAction)
	assert.Equal(t, mockState, transition.Execute(mockState))
}

func TestNewEndTransition(t *testing.T) {
	mockEvent := NewMockEvent(mockStartEventId)

	transition := NewEndTransition(mockEvent, NopAction)
	assert.NotNil(t, transition)
}

func TestEndTransition_Execute(t *testing.T) {
	mockState := NewMockState(NilState)
	mockEvent := NewMockEvent(mockStartEventId)

	transition := NewEndTransition(mockEvent, NopAction)
	assert.Equal(t, NilState, transition.Execute(mockState))
}
