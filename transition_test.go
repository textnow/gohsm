package hsm

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestNewExternalTransition(t *testing.T) {
	mockState := NewMockState(NilState)
	mockEvent := NewMockEvent(mockStartEventID)

	transition := NewExternalTransition(mockEvent, mockState, NopAction)
	assert.NotNil(t, transition)
}

func TestExternalTransition_Execute(t *testing.T) {
	mockState := NewMockState(NilState)
	mockEvent := NewMockEvent(mockStartEventID)
	logger, _ := zap.NewDevelopment()

	transition := NewExternalTransition(mockEvent, mockState, NopAction)
	assert.Equal(t, mockState, transition.Execute(logger, mockState))
}

func TestNewInternalTransition(t *testing.T) {
	mockEvent := NewMockEvent(mockStartEventID)

	transition := NewInternalTransition(mockEvent, NopAction)
	assert.NotNil(t, transition)
}

func TestInternalTransition_Execute(t *testing.T) {
	mockState := NewMockState(NilState)
	mockEvent := NewMockEvent(mockStartEventID)
	logger, _ := zap.NewDevelopment()

	transition := NewInternalTransition(mockEvent, NopAction)
	assert.Equal(t, mockState, transition.Execute(logger, mockState))
}

func TestNewEndTransition(t *testing.T) {
	mockEvent := NewMockEvent(mockStartEventID)

	transition := NewEndTransition(mockEvent, NopAction)
	assert.NotNil(t, transition)
}

func TestEndTransition_Execute(t *testing.T) {
	mockState := NewMockState(NilState)
	mockEvent := NewMockEvent(mockStartEventID)
	logger, _ := zap.NewDevelopment()

	transition := NewEndTransition(mockEvent, NopAction)
	assert.Equal(t, NilState, transition.Execute(logger, mockState))
}
