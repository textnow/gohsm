package hsm

// EventHandler contains the transition that is ready for execution
type EventHandler struct {
	transition Transition
}

// EventHandler constructor
func NewEventHandler(transition Transition) *EventHandler {
	return &EventHandler{
		transition: transition,
	}
}

// Execute performs the transition and returns the new current state
func (eventHandler *EventHandler) Execute(fromStateEngine *StateEngine) *StateEngine {
	return eventHandler.transition.Execute(fromStateEngine)
}
