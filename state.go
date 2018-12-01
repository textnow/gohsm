package hsm

// Each state in the state machine must implement this interface
type State interface {
	Name() string
	OnEnter(e Event) *StateEngine
	OnExit(e Event) *StateEngine
	GetEventHandler(e Event) *EventHandler
	GetStateEngine() *StateEngine
}
