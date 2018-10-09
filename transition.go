package hsm

// Transition represents a possible way of going from one state to another.
// It contains the logic used to determine what the next step will be.
type Transition interface {
	NextState() State
}

// DirectTransition represents a straightforward transition where the next state is known a priori.
type DirectTransition struct {
	next State
}

// NewDirectTransition creates a new direct transition using the supplied state as the destination state.
func NewDirectTransition(next State) *DirectTransition {
	return &DirectTransition{
		next: next,
	}
}

// NextState returns the saved next state.
func (dt *DirectTransition) NextState() State {
	return dt.next
}

// ConditionalEvaluator is a function that will return the desired state using its own evaluation criteria.
type ConditionalEvaluator func() State

// ConditionalTransition represents a transition where the outcome state depends on a method evaluation.
type ConditionalTransition struct {
	next ConditionalEvaluator
}

// NewConditionalTransition creates a new conditional transition using the supplied evaluator method.
func NewConditionalTransition(next ConditionalEvaluator) *ConditionalTransition {
	return &ConditionalTransition{
		next: next,
	}
}

// NextState returns the calculated next state.
func (ct *ConditionalTransition) NextState() State {
	return ct.next()
}

// EndTransition is a convenience variable representing a transition to an end state.
var EndTransition = &DirectTransition{
	next: EndState,
}
