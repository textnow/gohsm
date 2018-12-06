/*
HSM provides the framework for Hierarchical State Machine implementations.

Related Documents:
    - State Driven Development: https://confluence.enflick.com/x/3YwuAg
    - Go State Machine Framework: https://confluence.enflick.com/x/9AIgAw

Included in this framework are the following components:

  - StateMachineEngine:
    Engine that controls the state machine event processing

  - StateEngine:
    State implementation required by the StateMachineEngine to support event processing.  Operations that are
    common across all states are located in the StateEngine

  - State:
    Interface that must be implemented by all States in the StateMachine

  - Transition:
    Interface that is implemented by each of the different types of transitions:

      - ExternalTransition:
        Transition from current state to a different state.  On execution the following takes place:
          1. OnExit is called on the current state and all parent states up to the parent state that owns
             the new state (or the parent state is nil)
          2. action() associated with the the transition is called
          3. OnEnter() is called on the new state which may call OnEnter() for a sub-state.  The final
             new current state is returned by the OnEnter() call

      - InternalTransition:
        Transition within the current state.  On execution the following takes place:
          1. action() associated with the the transition is called

      - EndTransition:
        Transition from current state that terminates the state machine.  On execution the following takes place:
          1. OnExit is called on the current state and all parent states until there are no more parent states
          2. action() associated with the the transition is called

  - EventHandler:
    Wrapper around a transition.  For each event, the StateMachineEngine searches for an EventHandler starting with
    the current state and then proceeded with each parent state.  If found, then the transition contained in the
    EventHandler is executed.  Otherwise the event is dropped.

  - Event:
    An event represents something that has happened (login, logout, newCall, networkChange, etc.) that might drive
    a change in the state machine
*/
package hsm

import (
	"context"
	"go.uber.org/zap"
)

// StateMachineEngine manages event processing as implemented by each State
type StateMachineEngine struct {
	logger             *zap.Logger
	currentStateEngine *StateEngine
}

// StateMachineEngine constructor
func NewStateMachineEngine(logger *zap.Logger, startState State) *StateMachineEngine {
	sme := &StateMachineEngine{
		logger:             logger,
		currentStateEngine: startState.StateEngine(),
	}

	// This will ensure we are in the proper state starting from the beginning.
	sme.initialize()
	return sme
}

func (sme *StateMachineEngine) initialize() {
	sme.currentStateEngine = sme.currentStateEngine.OnEnter(StartEvent)
	sme.logger.Debug("state machine initialized",
		zap.String("starting_state", sme.currentStateEngine.Name()),
	)
}

func (sme *StateMachineEngine) handleEvent(e Event) bool {
	// Find an event handler (if none found then skip the event)
	eventHandler := sme.currentStateEngine.EventHandler(e)
	parentStateEngine := sme.currentStateEngine.ParentStateEngine()
	for eventHandler == nil {
		if parentStateEngine == nil {
			// Skip event handling
			return false
		}

		eventHandler = parentStateEngine.EventHandler(e)
		parentStateEngine = parentStateEngine.ParentStateEngine()
	}

	// Handle the event and update the current state
	sme.currentStateEngine = eventHandler.transition.Execute(sme.currentStateEngine)

	return true
}

// Run starts the StateMachineEngine and processes incoming events until the StateMachineEngine
// terminates (new currentState is nil after processing a transition) or the "done" event is received
func (sme *StateMachineEngine) Run(ctx context.Context, events <-chan Event) {
	go func() {
		for {
			select {
			case e, ok := <-events:
				if !ok {
					return
				}

				sme.logger.Debug("handling event",
					zap.String("event_id", e.ID()),
					zap.String("current_state", sme.currentStateEngine.Name()),
				)

				handled := sme.handleEvent(e)
				if !handled {
					sme.logger.Debug("event not handled",
						zap.String("event_id", e.ID()),
					)
					continue
				}

				if sme.currentStateEngine == nil {
					sme.logger.Debug("current state nil, terminating run loop")
					return
				}

				sme.logger.Debug("handled event",
					zap.String("current_state", sme.currentStateEngine.Name()),
				)
			case <-ctx.Done():
				sme.logger.Debug("received done on context")
				return
			}
		}
	}()
}
