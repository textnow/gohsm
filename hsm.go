package hsm

import (
	"context"
	"errors"
	"go.uber.org/zap"
)

var (
	// ErrAlreadyRegistered is returned if you attempt to register an event handler that is already registered.
	ErrAlreadyRegistered = errors.New("already registered")
)

// StateMachineEngine represents an instance of an executable state machine engine.
type StateMachineEngine struct {
	logger *zap.Logger

	t Transition

	curr State
}

// NewStateMachineEngine creates a new instance of the state machine engine with the specified state as the base.
func NewStateMachineEngine(logger *zap.Logger, t Transition) *StateMachineEngine {
	sme := &StateMachineEngine{
		logger: logger,
		t:      t,
		curr:   nil,
	}
	// This will ensure we are in the proper state starting from the beginning.
	sme.initialize()
	return sme

}

// CurrentState returns the currently active state.
func (sme *StateMachineEngine) CurrentState() State {
	return sme.curr
}

// HandleEvent attempts to synchronously process the supplied event.
// This is useful in testing contexts or when the Run() method of subscribing to a channel is not used.
func (sme *StateMachineEngine) HandleEvent(e Event) {
	handled := sme.handleEvent(e)
	if !handled {
		sme.logger.Debug("event not handled",
			zap.String("event_id", e.ID()),
		)
	} else {
		sme.logger.Debug("handled event",
			zap.String("current_state", sme.curr.Name()),
		)
	}
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
					zap.String("current_state", sme.curr.Name()),
				)

				handled := sme.handleEvent(e)
				if !handled {
					sme.logger.Debug("event not handled",
						zap.String("event_id", e.ID()),
					)
					continue
				}

				if sme.curr == EndState {
					sme.logger.Debug("current state nil, terminating run loop")
					return
				}

				sme.logger.Debug("handled event",
					zap.String("current_state", sme.curr.Name()),
				)
			case <- ctx.Done():
				sme.logger.Debug("received done on context")
				return
			}
		}
	}()
}

func (sme *StateMachineEngine) initialize() {
	start := sme.t.NextState()

	statesToEnter := append([]State{start}, sme.resolveLeaf(start)...)

	for _, state := range statesToEnter {
		state.handler().OnEnter(StartEvent)
	}

	sme.curr = statesToEnter[len(statesToEnter)-1]
	sme.logger.Debug("state machine initialized",
		zap.String("starting_state", sme.curr.Name()),
	)
}

// resolveLeaf finds all states from the current state to a leaf state and returns them ordered.
// This allows us to determine what we need to enter when partway through the state transition.
// Since we are going down the tree, we supply the StartEvent to determine how to go through the tree.
func (sme *StateMachineEngine) resolveLeaf(curr State) []State {
	start := curr
	var statesToEnter []State

	for {
		if _, ok := curr.(*LeafState); ok {
			break
		}

		h := curr.handlerForEvent(StartEvent)
		if h == nil {
			sme.logger.Fatal("state lacks handler for Start event",
				zap.String("state_id", curr.Name()),
			)
			return nil
		}

		curr = h.transition.NextState()
		statesToEnter = append(statesToEnter, curr)
	}

	// DEBUGGING
	var states []string
	for _, state := range statesToEnter {
		states = append(states, state.Name())
	}
	sme.logger.Debug("resolved leaf",
		zap.String("origin", start.Name()),
		zap.Strings("path", states),
	)

	return statesToEnter
}

// resolveTransition generates the lists of states that will be left, and the ones that will be entered,
// when transitioning from origin to dest. Both origin and dest must be valid states.
func (sme *StateMachineEngine) resolveTransition(origin State, dest State) ([]State, []State) {
	if origin == nil {
		sme.logger.Fatal("sme called with nil origin")
		return nil, nil
	} else if dest == nil {
		sme.logger.Fatal("sme called with nil dest",
			zap.String("origin_name", origin.Name()),
		)
		return nil, nil
	}

	// A self external transition
	if origin == dest {
		return []State{origin}, []State{dest}
	} else if origin.parent() == dest.parent() {
		return []State{origin}, []State{dest}
	} else if dest == EndState {
		return []State{origin}, []State{}
	}

	var originToRoot []State
	for i := origin; i != nil; i = i.parent() {
		sme.logger.Debug("originToRoot",
			zap.String("name", i.Name()),
		)
		originToRoot = append(originToRoot, i)
	}

	var destToRoot []State
	for i := dest; i != nil; i = i.parent() {
		sme.logger.Debug("destToRoot",
			zap.String("name", i.Name()),
		)
		destToRoot = append(destToRoot, i)
	}

	// We start at the origin. We examine the chain from the root to the dest
	// If we don't find the origin, proceed up one level and try again.
	// Once we find a common parent, we will know what we need to leave and enter to execute the transition.
	for ascIdx, asc := range originToRoot {
		sme.logger.Debug("asc",
			zap.String("name", asc.Name()),
			zap.Int("idx", ascIdx),
		)
		for descIdx, desc := range destToRoot {
			sme.logger.Debug("desc",
				zap.String("name", desc.Name()),
				zap.Int("idx", descIdx),
			)

			if asc == desc {
				sme.logger.Debug("found match")

				toExit := originToRoot[:ascIdx+1]
				var toEnter []State
				for i := descIdx; i >= 0; i-- {
					toEnter = append(toEnter, destToRoot[i])
				}

				// DEBUGGING
				var toExitStr []string
				for _, state := range toExit {
					toExitStr = append(toExitStr, state.Name())
				}
				var toEnterStr []string
				for _, state := range toEnter {
					toEnterStr = append(toEnterStr, state.Name())
				}
				sme.logger.Debug("resolved transition",
					zap.String("origin", origin.Name()),
					zap.String("destination", dest.Name()),
					zap.Strings("toExit", toExitStr),
					zap.Strings("toEnter", toEnterStr),
				)

				return toExit, toEnter
			}
		}
	}

	// No common parent.
	toExit := originToRoot
	var toEnter []State
	for i := len(destToRoot)-1; i >= 0; i-- {
		toEnter = append(toEnter, destToRoot[i])
	}

	// DEBUGGING
	var toExitStr []string
	for _, state := range toExit {
		toExitStr = append(toExitStr, state.Name())
	}
	var toEnterStr []string
	for _, state := range toEnter {
		toEnterStr = append(toEnterStr, state.Name())
	}
	sme.logger.Debug("resolved transition (no common parent)",
		zap.String("origin", origin.Name()),
		zap.String("destination", dest.Name()),
		zap.Strings("toExit", toExitStr),
		zap.Strings("toEnter", toEnterStr),
	)

	return toExit, toEnter
}

// handleEvent takes an event and determines how the state machine will handle it.
func (sme *StateMachineEngine) handleEvent(e Event) bool {
	var toLeave []State
	var toEnter []State

	for curr := sme.curr; curr != nil; curr = curr.parent() {
		eh := curr.handlerForEvent(e)

		if eh == nil {
			// If we don't have a handler at the current level we will go up a level.
			// If we find a handler one level higher we will, by definition, be leaving
			// the current state so we include the current state in our set to leave.
			toLeave = append(toLeave, curr)
			sme.logger.Debug("no registered handler for event, going up a level",
				zap.String("state_name", curr.Name()),
				zap.String("event_name", e.ID()),
			)
			continue
		}

		next := curr

		// If we have a transition this is an external transition.
		if eh.transition != nil {
			// We use the transition function to determine where we will go,
			// and then use the results from resolveTransition determine where
			// our target state will be.
			// We then use the resolveLeaf results to finish off the set of states we will be entering.
			// Remember! We must always end at a leaf state, we won't end in a composite state.
			next = eh.transition.NextState()

			currToLeave, currToEnter := sme.resolveTransition(curr, next)
			toLeave = append(toLeave, currToLeave...)
			toEnter = append(toEnter, currToEnter...)

			toEnter = append(toEnter, sme.resolveLeaf(next)...)

			// Our 'next' state isn't necessarily where we will end.
			// If toEnter exists we'll want to find the last (i.e. leaf) state to know where we will end.
			if len(toEnter) > 0 {
				next = toEnter[len(toEnter)-1]
			}

			// We only leave states if this is an external transition.
			for _, s := range toLeave {
				sme.logger.Debug("toExit",
					zap.String("name", s.Name()),
				)
				s.handler().OnExit(e)
			}
		}

		// We always execute the action.
		eh.action()

		// If we had an external transition we execute the onEnter methods of all the states we transition through.
		if eh.transition != nil {
			for _, s := range toEnter {
				sme.logger.Debug("toEnter",
					zap.String("name", s.Name()),
				)
				s.handler().OnEnter(e)
			}

			// We only need to update our current location if we're in an external transition.
			sme.curr = next
		}
		return true
	}

	return false
}
