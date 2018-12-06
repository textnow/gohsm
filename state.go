package hsm

// State defines the interface that must be implemented by each State
type State interface {
	Name() string
	OnEnter(e Event) *StateEngine
	OnExit(e Event) *StateEngine
	EventHandler(e Event) *EventHandler
	StateEngine() *StateEngine
}
