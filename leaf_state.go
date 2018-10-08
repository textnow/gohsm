package hsm

// LeafState represents a single state in the state machine. It is a terminal state.
type LeafState struct {
	name          string
	h             Handler
	eventHandlers map[string]*eventHandler

	parentState State
}

// NewLeafState creates a new named state that will trigger the supplied handler when entered/left.
func NewLeafState(name string, handler Handler) *LeafState {
	if handler == nil {
		handler = &EmptyHandler{}
	}
	return &LeafState{
		name:          name,
		h:             handler,
		eventHandlers: map[string]*eventHandler{},
	}
}

// SetExternalTransition adds a transition to the state that will cause this state to be left when the specified event occurs.
func (ls *LeafState) SetExternalTransition(e Event, a Action, t Transition) error {
	if _, ok := ls.eventHandlers[e.ID()]; ok {
		return ErrAlreadyRegistered
	}

	ls.eventHandlers[e.ID()] = &eventHandler{
		action:     a,
		transition: t,
	}
	return nil
}

// SetInternalTransition adds a transition to the state that will cause this action to be triggered when the specified event occurs.
func (ls *LeafState) SetInternalTransition(e Event, a Action) error {
	if _, ok := ls.eventHandlers[e.ID()]; ok {
		return ErrAlreadyRegistered
	}

	ls.eventHandlers[e.ID()] = &eventHandler{
		action:     a,
		transition: nil,
	}
	return nil
}

// Name is the name of this leaf state.
func (ls *LeafState) Name() string {
	return ls.name
}

func (ls *LeafState) setParent(s State) {
	ls.parentState = s
}

func (ls *LeafState) parent() State {
	return ls.parentState
}

func (ls *LeafState) handler() Handler {
	return ls.h
}

func (ls *LeafState) handlerForEvent(e Event) *eventHandler {
	if eh, ok := ls.eventHandlers[e.ID()]; ok {
		return eh
	}
	return nil
}
