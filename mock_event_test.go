package hsm

// MockEvent is mocks an Event for testing
type MockEvent struct {
	id string
}

// NewMockEvent constructor
func NewMockEvent(id string) *MockEvent {
	return &MockEvent{id}
}

// ID getter
func (e *MockEvent) ID() string {
	return e.id
}

var mockStartEventID = "start"
var mockSkipEventID = "skip"
var mockEndEventID = "end"
