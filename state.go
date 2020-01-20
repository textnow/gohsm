package hsm

import (
	"go.uber.org/zap"
)

// InternalState interface used by the StateMachine Framework
type InternalState interface {
	Name() string
	ParentState() State
	VerifyNotEntered()
	VerifyNotExited()
	Entered() bool
	Exited() bool
	Logger() *zap.Logger
}

// ExternalState interface used by the StateMachine Framework and must be implemented by each State
type ExternalState interface {
	OnEnter(e Event) State
	OnExit(e Event) State
	EventHandler(e Event) Transition
}

// State interface that must be implemented by all states in a StateMachine
// States can use the BaseState for the InternalState implementation
type State interface {
	InternalState
	ExternalState
}
