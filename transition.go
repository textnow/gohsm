package hsm

import (
	"reflect"
)

// Action is a type representing a function to be executed on a transition.
type Action func(srv Service)

// NopAction is a no-op action
var NopAction = func(srv Service) {}

// Transition defines the interface that is implemented by the different types of transitions
type Transition interface {
	Execute(srv Service, fromState State) State
}

// ExternalTransition is a transition from one state to a different state
type ExternalTransition struct {
	event   Event
	toState State
	action  Action
}

// ExternalTransition Constructor
func NewExternalTransition(event Event, toState State, action Action) *ExternalTransition {
	return &ExternalTransition{
		event:   event,
		toState: toState,
		action:  action,
	}
}

// Execute calls OnExit() for the fromState and all parent states up to the nil parent state
// or the parentState that is shared by the new toState.  Action() is then called and
// OnEnter() is finally called on the new toState.
func (t *ExternalTransition) Execute(srv Service, fromState State) State {
	// Call OnExit until one of the following is true:
	//   - parentState is NilState
	//   - parentState is equal to toState's parentState
	parentState := fromState.OnExit(t.event)
	for !IsNilState(parentState) {
		if !IsNilState(t.toState.ParentState()) && parentState.Name() == t.toState.ParentState().Name() {
			break
		}
		parentState = parentState.OnExit(t.event)
	}

	// Execute action on transition
	t.action(srv)

	// Enter toState and return new currentState
	return t.toState.OnEnter(t.event)
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
func (t *InternalTransition) Execute(srv Service, fromState State) State {
	t.action(srv)

	return fromState
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
func (t *EndTransition) Execute(srv Service, fromState State) State {
	// Call OnExit
	parentState := fromState.OnExit(t.event)
	for !IsNilState(parentState) {
		parentState = parentState.OnExit(t.event)
	}

	// Execute action on transition
	t.action(srv)

	// All done - turn out the lights
	return NilState
}

// UndefinedTransition is used to define NilTransition
type UndefinedTransition struct{}

func (tr *UndefinedTransition) Execute(srv Service, fromState State) State {
	panic("'Execute called on NilTransition - not cool!")
}

var NilTransition = &UndefinedTransition{}

// IsNilTransition returns true if the passed in transition is of type UndefinedTransition
func IsNilTransition(transition Transition) bool {
	return reflect.TypeOf(transition).String() == "*hsm.UndefinedTransition"
}
