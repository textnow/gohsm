package hsm

// Event is an interface to an event.
type Event interface {
	ID() string
}

// BaseEvent is a no-op event.
type BaseEvent struct {
	Name string
}

// ID returns the event ID.
func (be *BaseEvent) ID() string {
	return be.Name
}

// StartEvent is an event used to start a state machine evaluation.
var StartEvent = &BaseEvent{"Start"}
