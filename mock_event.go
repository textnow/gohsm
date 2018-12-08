package hsm

type MockEvent struct {
	id string
}

func NewMockEvent(id string) *MockEvent {
	return &MockEvent{id}
}

func (e *MockEvent) ID() string {
	return e.id
}

var mockStartEventId = "start"
var mockSkipEventId = "skip"
var mockEndEventId = "end"
