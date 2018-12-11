package main

import (
	"fmt"
	"github.com/Enflick/gohsm/example/simple/states"
	"os"
	"testing"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/Enflick/gohsm"
	"go.uber.org/zap"
)

func TestMain(m *testing.M) {
	if status := m.Run(); status != 0 {
		os.Exit(status)
	}

	// Run behaviour driven narrow integration tests that intersect with this service.
	status := godog.RunWithOptions(
		"godog",
		func(s *godog.Suite) {
			SimpleHSMContext(s)
		},
		godog.Options{
			Strict: true,
			Format: "progress",
			Paths:  []string{"./features/simple.feature"},
		},
	)
	os.Exit(status)
}

// simpleHSMTest is used for behavioural HSM testing using the simple HSM structure.
type simpleHSMTest struct {
	name string
	sm   *hsm.StateMachine

	startState hsm.State
	endState   hsm.State

	a bool
}

// simpleHSMContext defines the behaviour based testing context.
func SimpleHSMContext(s *godog.Suite) {
	test := simpleHSMTest{
		a: false,
	}

	// We use the same API across scenarios, this reduces isolation to catch
	// long-chain errors (pointer issues, etc.)
	s.BeforeSuite(func() {
		test.startState = hsm.NilState
		test.endState = hsm.NilState
	})

	s.BeforeScenario(func(s interface{}) {
		test.startState = hsm.NilState
		test.endState = hsm.NilState
		test.a = false
		states.LastActionIdExecuted = 0

		switch v := s.(type) {
		case *gherkin.Scenario:
			test.name = v.Name
		case *gherkin.ScenarioOutline:
			test.name = v.Name
		}
	})

	s.Step(`^test field A is true$`, test.setFieldAToTrue)
	s.Step(`^the simple hierarchical state machine is setup$`, test.theSimpleHSMIsInitialized)
	s.Step(`^the current state should be "([^"]*)"$`, test.currentStateIs)
	s.Step(`^event "([^"]*)" is generated$`, test.theEventIsGenerated)
	s.Step(`^the action with ID (\d+) should be executed$`, test.theActionIsExecuted)
	s.Step(`^the state "([^"]*)" should be entered$`, test.theStateIsEntered)
	s.Step(`^the state "([^"]*)" should be exited`, test.theStateIsExited)
	s.Step(`^the state "([^"]*)" should not be entered$`, test.theStateIsNotEntered)
	s.Step(`^the state "([^"]*)" should not be exited`, test.theStateIsNotExited)

}

func (t *simpleHSMTest) setFieldAToTrue() error {
	t.a = true
	return nil
}

func (t *simpleHSMTest) theSimpleHSMIsInitialized() error {
	logger, _ := zap.NewDevelopment()
	srv := states.NewSimpleService(hsm.NewDefaultService(logger), "test")
	startState := states.NewStateA(srv, t.a)

	t.sm = hsm.NewStateMachine(srv, startState)
	t.startState = t.sm.CurrentState()
	t.endState = t.sm.CurrentState()
	return nil
}

func (t *simpleHSMTest) currentStateIs(state string) error {
	if t.sm.CurrentState().Name() != state {
		return fmt.Errorf("expected state %s, got state %s", state, t.sm.CurrentState().Name())
	}
	return nil
}

func (t *simpleHSMTest) theStateIsEntered(stateName string) error {
	// Search the end state and all parent states for the desired state to check that it has been entered
	for state := t.endState; !hsm.IsNilState(state); state = state.ParentState() {
		if state.Entered() {
			fmt.Printf("State %s entered = true\n", state.Name())
		} else {
			fmt.Printf("State %s entered = false\n", state.Name())
		}

		if state.Name() == stateName {
			if !state.Entered() {
				return fmt.Errorf("expected state %s to be entered", state)
			}
			return nil
		}
	}

	return fmt.Errorf("did not find state %s", stateName)
}

func (t *simpleHSMTest) theStateIsExited(stateName string) error {
	// Search the start state and all parent states for the desired state to check that it has been exited
	for state := t.startState; !hsm.IsNilState(state); state = state.ParentState() {
		if state.Name() == stateName {
			if !state.Exited() {
				return fmt.Errorf("expected state %s to be exited", stateName)
			}
			return nil
		}
	}

	return fmt.Errorf("did not find state %s", stateName)
}

func (t *simpleHSMTest) theStateIsNotEntered(stateName string) error {
	// Search the end state and all parent states for the desired state to check that it has not been entered
	for state := t.endState; !hsm.IsNilState(state); state = state.ParentState() {
		if state.Name() == stateName {
			if state.Entered() {
				return fmt.Errorf("expected state %s to NOT be entered", stateName)
			}
			return nil
		}
	}

	return nil
}

func (t *simpleHSMTest) theStateIsNotExited(stateName string) error {
	// Search the start state and all parent states for the desired state to check that it has not been exited
	for state := t.startState; !hsm.IsNilState(state); state = state.ParentState() {
		if state.Name() == stateName {
			if state.Exited() {
				return fmt.Errorf("expected state %s to NOT be exited", stateName)
			}
			return nil
		}
	}

	return nil
}

func (t *simpleHSMTest) theEventIsGenerated(event string) error {
	t.startState = t.sm.CurrentState()
	t.sm.HandleEvent(states.NewKeyPressEvent(event))
	t.endState = t.sm.CurrentState()

	return nil
}

func (t *simpleHSMTest) theActionIsExecuted(actionID int) error {
	if states.LastActionIdExecuted != actionID {
		return fmt.Errorf("expected action %d to executed, but it was not", actionID)
	}
	return nil
}
