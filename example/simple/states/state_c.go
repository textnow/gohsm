package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
	"go.uber.org/zap"
)

// StateC represents State C
type StateC struct {
	hsm.BaseState
}

// NewStateC constructor
func NewStateC(logger *zap.Logger, parentState hsm.State) *StateC {
	hsm.Precondition(logger, parentState != nil, fmt.Sprintf("NewStateC: parentState cannot be nil"))

	return &StateC{
		BaseState: *hsm.NewBaseState("C", parentState, logger),
	}
}

// OnEnter enters this state and returns the new current state
func (s *StateC) OnEnter(event hsm.Event) hsm.State {
	s.VerifyNotEntered()
	s.Logger().Debug("->C;")
	return s
}

// OnExit exits this state and returns the parentState or nil if this state does not have a parent
func (s *StateC) OnExit(event hsm.Event) hsm.State {
	s.VerifyNotExited()
	s.Logger().Debug("<-C;")
	return s.ParentState()
}

// EventHandler returns the transition associated with the event or nil if this state does not handle the event
func (s *StateC) EventHandler(event hsm.Event) hsm.Transition {
	switch event.ID() {
	case ex.ID():
		return hsm.NewExternalTransition(event, NewStateC(s.Logger(), s.ParentState()), action6)
	case ey.ID():
		return hsm.NewInternalTransition(event, action7)
	default:
		return nil
	}
}

func action6(logger *zap.Logger) {
	logger.Debug("Action6")
	LastActionIDExecuted = 6
}

func action7(logger *zap.Logger) {
	logger.Debug("Action7")
	LastActionIDExecuted = 7
}
