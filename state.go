package hsm

import (
	"reflect"
)

// State interface that must be implemented by all states in a StateMachine
type State interface {
	Name() string
	OnEnter(e Event) State
	OnExit(e Event) State
	EventHandler(e Event) Transition
	ParentState() State
	Entered() bool
	Exited() bool
}

// UndefinedState is used to define the NilState
type UndefinedState struct{}

func (p *UndefinedState) Name() string {
	return "NilState"
}

func (p *UndefinedState) OnEnter(e Event) State {
	panic("OnEnter called on NilState - not cool!")
}

func (p *UndefinedState) OnExit(e Event) State {
	panic("OnExit called on NilState - not cool!")
}

func (p *UndefinedState) EventHandler(e Event) Transition {
	panic("EventHandler called on NilState - not cool!")
}

func (p *UndefinedState) ParentState() State {
	panic("ParentState called on NilState - not cool!")
}

func (p *UndefinedState) Entered() bool {
	panic("'Entered called on NilState - not cool!")
}

func (p *UndefinedState) Exited() bool {
	panic("'Entered called on NilState - not cool!")
}

var NilState = &UndefinedState{}

// IsNilState returns true if the passed in state is of type UndefinedState
func IsNilState(state State) bool {
	return reflect.TypeOf(state).String() == "*hsm.UndefinedState"
}
