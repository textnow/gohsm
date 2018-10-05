package hsm

// Event is an interface to an event.
type Event interface {
	ID() string
}

// BaseEvent is a no-op event.
type BaseEvent struct {
	name string
}

// ID returns the event ID.
func (be *BaseEvent) ID() string {
	return be.name
}

// StartEvent is an event used to start a state machine evaluation.
var StartEvent = &BaseEvent{"Start"}

// Handler represents something that will be triggered during a state transition.
type Handler interface {
	OnEnter(e Event)
	OnExit(e Event)
}

// EmptyHandler is a no-op handler.
type EmptyHandler struct{}

// OnEnter does nothing.
func (eh *EmptyHandler) OnEnter(Event) {}

// OnExit does nothing.
func (eh *EmptyHandler) OnExit(Event) {}

// Action is a type representing a function to be executed on a transition.
type Action func()
