package hsm

import (
	"fmt"
	"go.uber.org/zap"
)

// BaseState implements the InternalState interface
type BaseState struct {
	name        string
	parentState State
	logger      *zap.Logger
	entered     bool
	exited      bool
}

// NewBaseState constructor
func NewBaseState(name string, parentState State, logger *zap.Logger) *BaseState {
	return &BaseState{
		name:        name,
		parentState: parentState,
		logger:      logger,
	}
}

// Name getter
func (b *BaseState) Name() string {
	return b.name
}

// ParentState getter
func (b *BaseState) ParentState() State {
	return b.parentState
}

// Entered getter
func (b *BaseState) Entered() bool {
	return b.entered
}

// Exited getter
func (b *BaseState) Exited() bool {
	return b.exited
}

// Logger getter
func (b *BaseState) Logger() *zap.Logger {
	return b.logger
}

// VerifyNotEntered makes sure that this state is only entered once
func (b *BaseState) VerifyNotEntered() {
	Precondition(b.logger, !b.entered, fmt.Sprintf("State %s has already been entered", b.name))
	b.entered = true
}

// VerifyNotExited make sure that this state is only exited once
func (b *BaseState) VerifyNotExited() {
	Precondition(b.logger, !b.exited, fmt.Sprintf("State %s has already been exited", b.name))
	b.exited = true
}
