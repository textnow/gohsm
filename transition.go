package hsm

import (
	"github.com/Enflick/textnow-mono/bazel-textnow-mono/external/go_sdk/src/fmt"
	"go.uber.org/zap"
)

// Action is a type representing a function to be executed on a transition.
type Action func(logger *zap.Logger)

// NopAction is a no-op action
var NopAction = func(logger *zap.Logger) {}

// Transition defines the interface that is implemented by the different types of transitions
type Transition interface {
	Execute(logger *zap.Logger, fromState State) State
}

// ExternalTransition is a transition from one state to a different state
type ExternalTransition struct {
	event   Event
	toState State
	action  Action
}

// NewExternalTransition constructor
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
func (t *ExternalTransition) Execute(logger *zap.Logger, fromState State) State {
	// Call OnExit until one of the following is true:
	//   - parentState is NilState
	//   - parentState is equal to toState's parentState
	//   - infinite loop detected
	parentState := fromState.OnExit(t.event)
	loopCounter := 0
	for !IsNilState(parentState) {
		if !IsNilState(t.toState.ParentState()) && parentState.Name() == t.toState.ParentState().Name() {
			break
		}
		parentState = parentState.OnExit(t.event)
		loopCounter++
		Assertion(logger, loopCounter < 1000, fmt.Sprintf("ExternalTransition detected infinite loop"))
	}

	// Execute action on transition
	t.action(logger)

	// Enter toState and return new currentState
	return t.toState.OnEnter(t.event)
}

// InternalTransition is a transition within a state where only the action() is performed.
// OnEnter() and OnExit() are not called.
type InternalTransition struct {
	event  Event
	action Action
}

// NewInternalTransition constructor
func NewInternalTransition(event Event, action Action) *InternalTransition {
	return &InternalTransition{
		event:  event,
		action: action,
	}
}

// Execute calls the transition's action and returns the same fromState that started the transition
func (t *InternalTransition) Execute(logger *zap.Logger, fromState State) State {
	t.action(logger)

	return fromState
}

// EndTransition is the final transition that terminates the state machine
type EndTransition struct {
	event  Event
	action Action
}

// NewEndTransition constructor
func NewEndTransition(event Event, action Action) *EndTransition {
	return &EndTransition{
		event:  event,
		action: action,
	}
}

// Execute calls OnExit() for the fromState and all parent states up to the nil parent state
// Action() is then called and nil is returned for the new current state
func (t *EndTransition) Execute(logger *zap.Logger, fromState State) State {
	// Call OnExit
	parentState := fromState.OnExit(t.event)
	loopCounter := 0
	for !IsNilState(parentState) {
		parentState = parentState.OnExit(t.event)

		loopCounter++
		Assertion(logger, loopCounter < 1000, fmt.Sprintf("ExternalTransition detected infinite loop"))
	}

	// Execute action on transition
	t.action(logger)

	// All done - turn out the lights
	return NilState
}

// UndefinedTransition is used to define NilTransition
type UndefinedTransition struct{}

// Execute executes the transition
func (tr *UndefinedTransition) Execute(logger *zap.Logger, fromState State) State {
	panic("'Execute called on NilTransition - not cool!")
}

// NilTransition defines a nil transition
var NilTransition = &UndefinedTransition{}

// IsNilTransition returns true if the passed in transition is of type UndefinedTransition
func IsNilTransition(transition Transition) bool {
	_, ok := transition.(*UndefinedTransition)
	return ok
}
