package hsm

import (
	"context"
	"go.uber.org/zap"
)

type StateMachineEngine struct {
	logger             *zap.Logger
	currentStateEngine *StateEngine
}

func NewStateMachineEngine(logger *zap.Logger, startStateEngine *StateEngine) *StateMachineEngine {
	sme := &StateMachineEngine{
		logger:             logger,
		currentStateEngine: startStateEngine,
	}

	// This will ensure we are in the proper state starting from the beginning.
	sme.initialize()
	return sme
}

func (sme *StateMachineEngine) initialize() {
	sme.currentStateEngine = sme.currentStateEngine.GetState().OnEnter(StartEvent)
	sme.logger.Debug("state machine initialized",
		zap.String("starting_state", sme.currentStateEngine.Name()),
	)
}

func (sme *StateMachineEngine) handleEvent(e Event) bool {
	// Find an event handler (if none found then skip the event)
	eventHandler := sme.currentStateEngine.GetState().GetEventHandler(e)
	parentStateEngine := sme.currentStateEngine.GetParentStateEngine()
	for eventHandler == nil {
		if parentStateEngine == nil {
			// Skip event handling
			return false
		}

		eventHandler = parentStateEngine.GetState().GetEventHandler(e)
		parentStateEngine = parentStateEngine.GetParentStateEngine()
	}

	// Handle the event and update the current state
	sme.currentStateEngine = eventHandler.transition.execute(sme.currentStateEngine)

	return true
}

// Run sets up the state machine engine, primes the state machine with its start event
// and then continues to read from the supplied channel until its closed.
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
