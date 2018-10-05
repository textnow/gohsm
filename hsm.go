package hsm

import (
	"errors"
)

var (
	// ErrAlreadyRegistered is returned if you attempt to register an event handler that is already registered.
	ErrAlreadyRegistered = errors.New("already registered")
)

// eventHandler is an internal representation of an event handler.
// If transition is nil it represents an internal transition.
type eventHandler struct {
	action     Action
	transition Transition
}

// State represents a single state in the state machine. It may have a sub state machine.
type State struct {
	name          string
	handler       Handler
	eventHandlers map[string]eventHandler

	subStateMachine *StateMachine
}

// NewState creates a new named state that will trigger the supplied handler when entered/left.
func NewState(name string, handler Handler) *State {
	if handler == nil {
		handler = &EmptyHandler{}
	}
	return &State{
		name:          name,
		handler:       handler,
		eventHandlers: map[string]eventHandler{},
	}
}

// NewStateWithSubStateMachine creates a new named state that will trigger the suppled handler, along with
// a child state machine that will trigger when this state is active.
func NewStateWithSubStateMachine(name string, handler Handler, subStateMachine *StateMachine) *State {
	if handler == nil {
		handler = &EmptyHandler{}
	}
	return &State{
		name:            name,
		handler:         handler,
		eventHandlers:   map[string]eventHandler{},
		subStateMachine: subStateMachine,
	}
}

// SetExternalTransition adds a transition to the state that will cause this state to be left when the specified event occurs.
func (is *State) SetExternalTransition(e Event, a Action, t Transition) error {
	if _, ok := is.eventHandlers[e.ID()]; ok {
		return ErrAlreadyRegistered
	}

	is.eventHandlers[e.ID()] = eventHandler{
		action:     a,
		transition: t,
	}
	return nil
}

// SetInternalTransition adds a transition to the state that will cause this action to be triggered when the specified event occurs.
func (is *State) SetInternalTransition(e Event, a Action) error {
	if _, ok := is.eventHandlers[e.ID()]; ok {
		return ErrAlreadyRegistered
	}

	is.eventHandlers[e.ID()] = eventHandler{
		action:     a,
		transition: nil,
	}
	return nil
}

// StateMachine is an instance of a state machine. It maintains the transitions required for states.
type StateMachine struct {
	start *State
	curr  *State
}

// NewStateMachine creates a new state machine. The supplied transition is executed when the state machine is first started (its owning state is entered).
func NewStateMachine(t Transition) *StateMachine {
	start := NewState("internal_start", nil)
	start.eventHandlers[StartEvent.ID()] = eventHandler{
		action:     nil,
		transition: t,
	}

	return &StateMachine{
		start: start,
		curr:  start,
	}
}

// currentState is an internal handler that returns the current state (or state of a sub state machine if set).
func (sm *StateMachine) currentState() *State {
	if sm.curr == nil {
		return nil
	}

	if sm.curr.subStateMachine != nil {
		return sm.curr.subStateMachine.currentState()
	}
	return sm.curr
}

// handleEvent handles the supplied event.
// If the state machine doesn't have this event registered, and any sub state machines don't have this event registered, false will be returned.
// If the state is registered it will be handled and return true.
func (sm *StateMachine) handleEvent(e Event) bool {
	// Do not attempt to handle the event if we're done (i.e. no current state left).
	if sm.curr == nil {
		return false
	}

	if sm.curr.subStateMachine != nil {
		if sm.curr.subStateMachine.handleEvent(e) {
			return true
		}
	}

	ev, ok := sm.curr.eventHandlers[e.ID()]

	// The current state doesn't have a handler for this event, so return 'unhandled'
	if !ok {
		return false
	}

	// This means we have an external transition.
	if ev.transition != nil {
		if sm.curr.subStateMachine != nil && sm.curr.subStateMachine.curr != nil {
			sm.curr.subStateMachine.curr.handler.OnExit(e)
		}
		if sm.curr.subStateMachine != nil {
			// Set this state machine up to be re-entered
			sm.curr.subStateMachine.curr = sm.curr.subStateMachine.start
		}

		sm.curr.handler.OnExit(e)
	}

	if ev.action != nil {
		ev.action()
	}

	if ev.transition != nil {
		next := ev.transition.NextState()
		sm.curr = next

		if next == nil {
			return true
		}

		sm.curr.handler.OnEnter(e)

		if sm.curr.subStateMachine != nil {
			sm.curr.subStateMachine.handleEvent(StartEvent)
		}
	}

	return true
}

// StateMachineEngine represents an instance of an executable state machine engine.
type StateMachineEngine struct {
	sm *StateMachine
}

// NewStateMachineEngine creates a new instance of the state machine engine with the specified state as the base.
func NewStateMachineEngine(sm *StateMachine) *StateMachineEngine {
	return &StateMachineEngine{
		sm: sm,
	}
}

// Run sets up the state machine engine, primes the state machine with its start event
// and then continues to read from the supplied channel until its closed.
func (sme *StateMachineEngine) Run(events <-chan Event) {
	sme.sm.handleEvent(StartEvent)

	for {
		e, ok := <-events
		if !ok {
			return
		}

		sme.sm.handleEvent(e)
	}
}
