package hsm

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func NewMockService() (Service) {
	logger, _ := zap.NewDevelopment()
	return NewDefaultService(logger)
}

func TestNewExternalTransition(t *testing.T) {
	mockState := NewMockState(NilState)
	mockEvent := NewMockEvent(mockStartEventId)

	transition := NewExternalTransition(mockEvent, mockState, NopAction)
	assert.NotNil(t, transition)
}

func TestExternalTransition_Execute(t *testing.T) {
	mockState := NewMockState(NilState)
	mockEvent := NewMockEvent(mockStartEventId)
	mockService := NewMockService()

	transition := NewExternalTransition(mockEvent, mockState, NopAction)
	assert.Equal(t, mockState, transition.Execute(mockService, mockState))
}

func TestNewInternalTransition(t *testing.T) {
	mockEvent := NewMockEvent(mockStartEventId)

	transition := NewInternalTransition(mockEvent, NopAction)
	assert.NotNil(t, transition)
}

func TestInternalTransition_Execute(t *testing.T) {
	mockState := NewMockState(NilState)
	mockEvent := NewMockEvent(mockStartEventId)
	mockService := NewMockService()

	transition := NewInternalTransition(mockEvent, NopAction)
	assert.Equal(t, mockState, transition.Execute(mockService, mockState))
}

func TestNewEndTransition(t *testing.T) {
	mockEvent := NewMockEvent(mockStartEventId)

	transition := NewEndTransition(mockEvent, NopAction)
	assert.NotNil(t, transition)
}

func TestEndTransition_Execute(t *testing.T) {
	mockState := NewMockState(NilState)
	mockEvent := NewMockEvent(mockStartEventId)
	mockService := NewMockService()

	transition := NewEndTransition(mockEvent, NopAction)
	assert.Equal(t, NilState, transition.Execute(mockService, mockState))
}
