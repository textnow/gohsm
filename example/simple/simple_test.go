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
	sm   *hsm.StateMachineEngine

	startStateEngine *hsm.StateEngine
	endStateEngine   *hsm.StateEngine

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
		test.startStateEngine = nil
		test.endStateEngine = nil
	})

	s.BeforeScenario(func(s interface{}) {
		test.startStateEngine = nil
		test.endStateEngine = nil
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
	startState := states.NewStateA(t.a)

	t.sm = hsm.NewStateMachineEngine(logger, startState)
	t.startStateEngine = t.sm.CurrentStateEngine()
	t.endStateEngine = t.sm.CurrentStateEngine()
	return nil
}

func (t *simpleHSMTest) currentStateIs(state string) error {
	if t.sm.CurrentStateEngine().Name() != state {
		return fmt.Errorf("expected state %s, got state %s", state, t.sm.CurrentStateEngine().Name())
	}
	return nil
}

func (t *simpleHSMTest) theStateIsEntered(state string) error {
	// Search the end state and all parent states for the desired state to check that it has been entered
	for stateEngine := t.endStateEngine; stateEngine != nil; stateEngine = stateEngine.ParentStateEngine() {
		if stateEngine.Entered() {
			fmt.Printf("State %s entered = true\n", stateEngine.Name())
		} else {
			fmt.Printf("State %s entered = false\n", stateEngine.Name())
		}

		if stateEngine.Name() == state {
			if !stateEngine.Entered() {
				return fmt.Errorf("expected state %s to be entered", state)
			}
			return nil
		}
	}

	return fmt.Errorf("did not find state %s", state)
}

func (t *simpleHSMTest) theStateIsExited(state string) error {
	// Search the start state and all parent states for the desired state to check that it has been exited
	for stateEngine := t.startStateEngine; stateEngine != nil; stateEngine = stateEngine.ParentStateEngine() {
		if stateEngine.Name() == state {
			if !stateEngine.Exited() {
				return fmt.Errorf("expected state %s to be exited", state)
			}
			return nil
		}
	}

	return fmt.Errorf("did not find state %s", state)
}

func (t *simpleHSMTest) theStateIsNotEntered(state string) error {
	// Search the end state and all parent states for the desired state to check that it has not been entered
	for stateEngine := t.endStateEngine; stateEngine != nil; stateEngine = stateEngine.ParentStateEngine() {
		if stateEngine.Name() == state {
			if stateEngine.Entered() {
				return fmt.Errorf("expected state %s to NOT be entered", state)
			}
			return nil
		}
	}

	return nil
}

func (t *simpleHSMTest) theStateIsNotExited(state string) error {
	// Search the start state and all parent states for the desired state to check that it has not been exited
	for stateEngine := t.startStateEngine; stateEngine != nil; stateEngine = stateEngine.ParentStateEngine() {
		if stateEngine.Name() == state {
			if stateEngine.Exited() {
				return fmt.Errorf("expected state %s to NOT be exited", state)
			}
			return nil
		}
	}

	return nil
}

func (t *simpleHSMTest) theEventIsGenerated(event string) error {
	t.startStateEngine = t.sm.CurrentStateEngine()
	t.sm.HandleEvent(states.NewKeyPressEvent(event))
	t.endStateEngine = t.sm.CurrentStateEngine()

	return nil
}

func (t *simpleHSMTest) theActionIsExecuted(actionID int) error {
	if states.LastActionIdExecuted != actionID {
		return fmt.Errorf("expected action %d to executed, but it was not", actionID)
	}
	return nil
}
