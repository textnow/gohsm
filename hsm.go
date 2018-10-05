package hsm

import (
	"errors"
)

var (
	ErrAlreadyRegistered = errors.New("already registered")
)

type Handler interface {
	OnEnter()
	OnExit()
}

type Event interface {
	ID() string
}

type BaseEvent struct {
	name string
}
func (be *BaseEvent) ID() string {
	return be.name
}

type Action func()

var StartEvent = &BaseEvent{"Start"}
var EndEvent = &BaseEvent{"End"}

func EmptyAction() {
}

type EmptyHandler struct {
}
func (eh *EmptyHandler) OnEnter() {
}
func (eh *EmptyHandler) OnExit() {
}

type Transition interface {
	NextState() *State
}

type DirectTransition struct {
	next *State
}
func NewDirectTransition(next *State) *DirectTransition {
	return &DirectTransition{
		next: next,
	}
}
func (dt *DirectTransition) NextState() *State {
	return dt.next
}

type ConditionalTransition struct {
	next func() *State
}
func NewConditionalTransition(next func() *State) *ConditionalTransition {
	return &ConditionalTransition{
		next: next,
	}
}
func (ct *ConditionalTransition) NextState() *State {
	return ct.next()
}

var EndTransition = &DirectTransition{}


type eventHandler struct {
	action Action
	transition Transition
}

type State struct {
	name          string
	handler       Handler
	eventHandlers map[string]eventHandler

	subStateMachine *StateMachine
}

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

func (is *State) SetExternalTransition(e Event, a Action, t Transition) error {
	if _, ok := is.eventHandlers[e.ID()]; ok {
		return ErrAlreadyRegistered
	}

	is.eventHandlers[e.ID()] = eventHandler{
		action: a,
		transition: t,
	}
	return nil
}

func (is *State) SetInternalTransition(e Event, a Action) error {
	if _, ok := is.eventHandlers[e.ID()]; ok {
		return ErrAlreadyRegistered
	}

	is.eventHandlers[e.ID()] = eventHandler{
		action: a,
		transition: nil,
	}
	return nil
}

type StateMachine struct {
	start *State
	curr *State
}

func NewStateMachine(t Transition) *StateMachine {
	start := NewState("internal_start", nil)
	start.eventHandlers[StartEvent.ID()] = eventHandler{
		action: nil,
		transition: t,
	}

	return &StateMachine{
		start: start,
		curr: start,
	}
}

func (sm *StateMachine) currentState() *State {
	if sm.curr == nil {
		return nil
	}

	if sm.curr.subStateMachine != nil {
		return sm.curr.subStateMachine.currentState()
	} else {
		return sm.curr
	}
}

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
			sm.curr.subStateMachine.curr.handler.OnExit()
		}
		if sm.curr.subStateMachine != nil {
			// Set this state machine up to be re-entered
			sm.curr.subStateMachine.curr = sm.curr.subStateMachine.start
		}

		sm.curr.handler.OnExit()
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

		sm.curr.handler.OnEnter()

		if sm.curr.subStateMachine != nil {
			sm.curr.subStateMachine.handleEvent(StartEvent)
		}
	}

	return true
}

type StateMachineEngine struct {
	sm *StateMachine
}

func NewStateMachineEngine(sm *StateMachine) *StateMachineEngine {
	return &StateMachineEngine{
		sm: sm,
	}
}

func (sme *StateMachineEngine) dispatchEvent(e Event) {
	sme.sm.handleEvent(e)
}

func (sme *StateMachineEngine) Run(events <-chan Event) {
	sme.sm.handleEvent(StartEvent)

	for {
		e, ok := <-events
		if !ok {
			return
		}

		sme.dispatchEvent(e)
	}
}

