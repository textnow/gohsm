package hsm

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

// Name returns the state's name
func (p *UndefinedState) Name() string {
	return "NilState"
}

// OnEnter enters this state
func (p *UndefinedState) OnEnter(e Event) State {
	panic("OnEnter called on NilState - not cool!")
}

// OnExit exits this state
func (p *UndefinedState) OnExit(e Event) State {
	panic("OnExit called on NilState - not cool!")
}

// EventHandler checks to see if this state handles the specified event
func (p *UndefinedState) EventHandler(e Event) Transition {
	panic("EventHandler called on NilState - not cool!")
}

// ParentState gets this state's parent state
func (p *UndefinedState) ParentState() State {
	panic("ParentState called on NilState - not cool!")
}

// Entered returns true if OnEnter has been called
func (p *UndefinedState) Entered() bool {
	panic("'Entered called on NilState - not cool!")
}

// Exited returns true if OnExit has been called
func (p *UndefinedState) Exited() bool {
	panic("'Entered called on NilState - not cool!")
}

// NilState defines a nil State
var NilState = &UndefinedState{}

// IsNilState returns true if the passed in state is of type UndefinedState
func IsNilState(state State) bool {
	_, ok := state.(*UndefinedState)
	return ok
}
