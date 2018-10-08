package hsm

// CompositeState represents a state that will manage itself and states below it.
type CompositeState struct {
	LeafState

	childStates []State
}

// NewCompositeState creates a new named state that will have states below it in the hierarchy.
// The children will all be direct children of this state. The EndState constant does not need to be listed.
// The first child in the array will be the default destination state when this state is entered.
func NewCompositeState(name string, handler Handler, children []State) *CompositeState {
	t := NewDirectTransition(children[0])
	return NewCompositeStateWithTransition(name, handler, children, t)
}

// NewCompositeStateWithTransition creates a new named state that will have states below it in the hierarchy.
// The children will all be direct children of this state. The EndState constant does not need to be listed.
// The supplied transition will be use to determine which of the children are selected as the default entering state.
func NewCompositeStateWithTransition(name string, handler Handler, children []State, initialTransition Transition) *CompositeState {
	cs := &CompositeState{
		childStates: children,
	}
	cs.name = name
	cs.h = handler
	cs.eventHandlers = map[string]*eventHandler{
		StartEvent.ID(): {
			action:     NopAction,
			transition: initialTransition,
		},
	}

	// TODO: we can assert that the children don't already exist to a parent.
	// This will help safeguard against programmer errors.
	for _, child := range children {
		child.setParent(cs)
	}

	return cs
}
