package hsm

// StateEngine is an internal representation of a State that is used by the StateMachineEngine
type StateEngine struct {
	state             State
	parentStateEngine *StateEngine
}

// StateEngine constructor
func NewStateEngine(state State, parentStateEngine *StateEngine) *StateEngine {
	stateEngine := &StateEngine{
		state:             state,
		parentStateEngine: parentStateEngine,
	}

	return stateEngine
}

// Name returns the state's name
func (a *StateEngine) Name() string {
	return a.state.Name()
}

// ParentStateEngine returns this StateEngine's parentStateEngine or nil if there is not parent  state
func (a *StateEngine) ParentStateEngine() *StateEngine {
	return a.parentStateEngine
}

// OnEnter returns the new current StateEngine
func (a *StateEngine) OnEnter(event Event) *StateEngine {
	return a.state.OnEnter(event)
}

// OnExit allows the state to perform exit house keeping and returns the state's parentStateEngine
func (a *StateEngine) OnExit(event Event) *StateEngine {
	return a.state.OnExit(event)
}

// EventHandler returns a handler for the current event or nil if the state does not process the event
func (a *StateEngine) EventHandler(event Event) *EventHandler {
	return a.state.EventHandler(event)
}
