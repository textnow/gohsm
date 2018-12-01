package hsm

// Action is a type representing a function to be executed on a transition.
type Action func()

// NopAction is a no-op action
var NopAction = func() {}

type Transition interface {
	execute(fromStateEngine *StateEngine) *StateEngine
}

type ExternalTransition struct {
	event         Event
	toStateEngine *StateEngine
	action        Action
}

func NewExternalTransition(event Event, toStateEngine *StateEngine, action Action) *ExternalTransition {
	return &ExternalTransition{
		event:         event,
		toStateEngine: toStateEngine,
		action:        action,
	}
}

func (t *ExternalTransition) execute(fromStateEngine *StateEngine) *StateEngine {
	// Call OnExit starting with the fromStateEngine and all parentStateEngines
	// until parentStateEngine is nil or equal to the toState's parentStateEngine
	parentStateEngine := fromStateEngine.GetState().OnExit(t.event)
	for parentStateEngine != nil &&
		(t.toStateEngine.GetParentStateEngine() == nil ||
			parentStateEngine.Name() != t.toStateEngine.GetParentStateEngine().Name()) {
		parentStateEngine = parentStateEngine.GetState().OnExit(t.event)
	}

	// Execute action on transition
	t.action()

	// Enter toState and return new currentState
	return t.toStateEngine.GetState().OnEnter(t.event)
}

type InternalTransition struct {
	event  Event
	action Action
}

func NewInternalTransition(event Event, action Action) *InternalTransition {
	return &InternalTransition{
		event:  event,
		action: action,
	}
}

func (t *InternalTransition) execute(fromStateEngine *StateEngine) *StateEngine {
	// Execute action on transition only (OnEnter and OnExit are not called
	// when executing an internal transition
	t.action()

	return fromStateEngine
}

type DaveEndTransition struct {
	event  Event
	action Action
}

func NewEndTransition(event Event, action Action) *DaveEndTransition {
	return &DaveEndTransition{
		event:  event,
		action: action,
	}
}

func (t *DaveEndTransition) execute(fromStateEngine *StateEngine) *StateEngine {
	// Call OnExit starting with the fromStateEngine and all parentStateEngines
	// until parentStateEngine is nil (final cleanup before state machine exit)
	parentStateEngine := fromStateEngine.GetState().OnExit(t.event)
	for parentStateEngine != nil {
		parentStateEngine = parentStateEngine.GetState().OnExit(t.event)
	}

	// Execute action on transition
	t.action()

	// All done - turn out the lights
	return nil
}
