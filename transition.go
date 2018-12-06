package hsm

// Action is a type representing a function to be executed on a transition.
type Action func()

// NopAction is a no-op action
var NopAction = func() {}

// Transition defines the interface that is implemented by the different types of transitions
type Transition interface {
	Execute(fromStateEngine *StateEngine) *StateEngine
}

// ExternalTransition is a transition from one state to a different state
type ExternalTransition struct {
	event         Event
	toStateEngine *StateEngine
	action        Action
}

// ExternalTransition Constructor
func NewExternalTransition(event Event, toStateEngine *StateEngine, action Action) *ExternalTransition {
	return &ExternalTransition{
		event:         event,
		toStateEngine: toStateEngine,
		action:        action,
	}
}

// Execute calls OnExit() for the fromState and all parent states up to the nil parent state
// or the parentState that is shared by the new toState.  Action() is then called and
// OnEnter() is finally called on the new toState.
func (t *ExternalTransition) Execute(fromStateEngine *StateEngine) *StateEngine {
	// Call OnExit
	parentStateEngine := fromStateEngine.OnExit(t.event)
	for parentStateEngine != nil &&
		(t.toStateEngine.ParentStateEngine() == nil ||
			parentStateEngine.Name() != t.toStateEngine.ParentStateEngine().Name()) {
		parentStateEngine = parentStateEngine.OnExit(t.event)
	}

	// Execute action on transition
	t.action()

	// Enter toState and return new currentState
	return t.toStateEngine.OnEnter(t.event)
}

// InternalTransition is a transition within a state where only the action() is performed.
// OnEnter() and OnExit() are not called.
type InternalTransition struct {
	event  Event
	action Action
}

// InternalTransition constructor
func NewInternalTransition(event Event, action Action) *InternalTransition {
	return &InternalTransition{
		event:  event,
		action: action,
	}
}

// Execute calls the transition's action and returns the same fromState that started the transition
func (t *InternalTransition) Execute(fromStateEngine *StateEngine) *StateEngine {
	t.action()

	return fromStateEngine
}

// EndTransition is the final transition that terminates the state machine
type EndTransition struct {
	event  Event
	action Action
}

// EndTransition constructor
func NewEndTransition(event Event, action Action) *EndTransition {
	return &EndTransition{
		event:  event,
		action: action,
	}
}

// Execute calls OnExit() for the fromState and all parent states up to the nil parent state
// Action() is then called and nil is returned for the new current state
func (t *EndTransition) Execute(fromStateEngine *StateEngine) *StateEngine {
	// Call OnExit
	parentStateEngine := fromStateEngine.OnExit(t.event)
	for parentStateEngine != nil {
		parentStateEngine = parentStateEngine.OnExit(t.event)
	}

	// Execute action on transition
	t.action()

	// All done - turn out the lights
	return nil
}
