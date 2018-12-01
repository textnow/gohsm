package hsm

// Wrapper for a transition between states
type EventHandler struct {
	transition Transition
}

func NewEventHandler(transition Transition) *EventHandler {
	return &EventHandler{
		transition: transition,
	}
}

func (eventHandler *EventHandler) GetTransition() Transition {
	return eventHandler.transition
}
