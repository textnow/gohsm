package hsm

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
	"testing"
)

func TestNewBaseState(t *testing.T) {
	logger := zaptest.NewLogger(t)
	baseState := NewBaseState("baseStateName", nil, logger)

	assert.Equal(t, "baseStateName", baseState.Name())
	assert.Nil(t, baseState.ParentState())
	assert.False(t, baseState.Entered())
	assert.False(t, baseState.Exited())
	assert.Equal(t, logger, baseState.Logger())
}

func TestBaseState_VerifyNotEntered(t *testing.T) {
	logger := zaptest.NewLogger(t)
	baseState := NewBaseState("baseStateName", nil, logger)

	assert.False(t, baseState.Entered())
	baseState.VerifyNotEntered()
	assert.True(t, baseState.Entered())
}

func TestBaseState_VerifyNotExited(t *testing.T) {
	logger := zaptest.NewLogger(t)
	baseState := NewBaseState("baseStateName", nil, logger)

	assert.False(t, baseState.Exited())
	baseState.VerifyNotExited()
	assert.True(t, baseState.Exited())
}
