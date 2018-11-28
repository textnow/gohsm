package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/rmrobinson-textnow/gohsm"
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
			simpleHSMContext(s)
		},
		godog.Options{
			Strict: true,
			Format: "progress",
			Paths:  []string{"./features/simple.feature"},
		},
	)
	os.Exit(status)
}

func (t *simpleHSMTest) action1() {
	t.actionsExecuted[1] = true
}
func (t *simpleHSMTest) action2() {
	t.actionsExecuted[2] = true
}
func (t *simpleHSMTest) action3() {
	t.actionsExecuted[3] = true
}
func (t *simpleHSMTest) action4() {
	t.actionsExecuted[4] = true
}
func (t *simpleHSMTest) action5() {
	t.actionsExecuted[5] = true
}
func (t *simpleHSMTest) action6() {
	t.actionsExecuted[6] = true
}
func (t *simpleHSMTest) action7() {
	t.actionsExecuted[7] = true
}

type testState struct {
	entered bool
	exited  bool
}

func (s *testState) OnEnter(event hsm.Event) {
	s.entered = true
}
func (s *testState) OnExit(event hsm.Event) {
	s.exited = true
}

// simpleHSMTest is used for behavioural HSM testing using the simple HSM structure.
type simpleHSMTest struct {
	name string
	sm   *hsm.StateMachineEngine

	actionsExecuted map[int]bool
	states          map[string]*testState

	a bool
}

// simpleHSMContext defines the behaviour based testing context.
func simpleHSMContext(s *godog.Suite) {
	test := simpleHSMTest{
		states: map[string]*testState{},
		a:      false,
	}

	sd := &testState{}
	sc := &testState{}
	sb := &testState{}
	sa := &testState{}

	// We use the same API across scenarios, this reduces isolation to catch
	// long-chain errors (pointer issues, etc.)
	s.BeforeSuite(func() {
		test.states["D"] = sd
		test.states["C"] = sc
		test.states["B"] = sb
		test.states["A"] = sa
	})

	s.BeforeScenario(func(s interface{}) {
		test.actionsExecuted = map[int]bool{}
		test.resetStateEntriesAndExits()
		test.a = false

		switch v := s.(type) {
		case *gherkin.Scenario:
			test.name = v.Name
		case *gherkin.ScenarioOutline:
			test.name = v.Name
		}
	})

	s.Step(`^the simple hierarchical state machine is setup$`, test.theSimpleHSMIsInitialized)
	s.Step(`^the initial state is "([^"]*)"$`, test.setInitialState)
	s.Step(`^test field A is true$`, test.setFieldAToTrue)
	s.Step(`^the current state should be "([^"]*)"$`, test.currentStateIs)
	s.Step(`^event "([^"]*)" is generated$`, test.theEventIsGenerated)
	s.Step(`^the action with ID (\d+) should be executed$`, test.theActionIsExecuted)
	s.Step(`^the state "([^"]*)" should be entered$`, test.theStateIsEntered)
	s.Step(`^the state "([^"]*)" should be exited`, test.theStateIsExited)
	s.Step(`^the state "([^"]*)" should not be entered$`, test.theStateIsNotEntered)
	s.Step(`^the state "([^"]*)" should not be exited`, test.theStateIsNotExited)
	s.Step(`^the entries and exits are reset$`, test.resetStateEntriesAndExits)

}

func (t *simpleHSMTest) setFieldAToTrue() error {
	t.a = true
	return nil
}

func (t *simpleHSMTest) setInitialState(state string) error {
	if state == "C" {
		t.sm.HandleEvent(&keyPressEvent{"a"})
	}
	t.resetStateEntriesAndExits()
	return nil
}

func (t *simpleHSMTest) theSimpleHSMIsInitialized() error {
	ea := &keyPressEvent{"a"}
	eb := &keyPressEvent{"b"}
	ec := &keyPressEvent{"c"}
	ed := &keyPressEvent{"d"}
	ee := &keyPressEvent{"e"}
	ex := &keyPressEvent{"x"}
	ey := &keyPressEvent{"y"}

	c := hsm.NewLeafState("C", t.states["C"])
	c.SetExternalTransition(ex, t.action6, hsm.NewDirectTransition(c))
	c.SetInternalTransition(ey, t.action7)

	b := hsm.NewLeafState("B", t.states["B"])
	b.SetExternalTransition(ea, t.action1, hsm.NewDirectTransition(c))

	d := hsm.NewLeafState("D", t.states["D"])
	d.SetExternalTransition(ee, t.action5, hsm.EndTransition)

	a := hsm.NewCompositeStateWithTransition("A", t.states["A"], []hsm.State{
		b,
		c,
	}, hsm.NewConditionalTransition(func() hsm.State {
		if t.a {
			return b
		}
		return c
	}))

	a.SetInternalTransition(eb, t.action2)
	a.SetExternalTransition(ec, t.action3, hsm.NewDirectTransition(a))
	a.SetExternalTransition(ed, t.action4, hsm.NewDirectTransition(d))

	logger, _ := zap.NewDevelopment()
	t.sm = hsm.NewStateMachineEngine(logger, hsm.NewDirectTransition(a))
	return nil
}

func (t *simpleHSMTest) resetStateEntriesAndExits() error {
	for stateID, state := range t.states {
		state.entered = false
		state.exited = false
		t.states[stateID] = state
	}
	return nil
}

func (t *simpleHSMTest) currentStateIs(state string) error {
	if t.sm.CurrentState().Name() != state {
		return fmt.Errorf("expected state %s, got state %s", state, t.sm.CurrentState().Name())
	}
	return nil
}

func (t *simpleHSMTest) theStateIsEntered(state string) error {
	if !t.states[state].entered {
		return fmt.Errorf("expected state %s to be entered", state)
	}
	return nil
}

func (t *simpleHSMTest) theStateIsExited(state string) error {
	if !t.states[state].exited {
		return fmt.Errorf("expected state %s to be exited", state)
	}
	return nil
}

func (t *simpleHSMTest) theStateIsNotEntered(state string) error {
	if t.states[state].entered {
		return fmt.Errorf("expected state %s to not be entered", state)
	}
	return nil
}

func (t *simpleHSMTest) theStateIsNotExited(state string) error {
	if t.states[state].exited {
		return fmt.Errorf("expected state %s to not be exited", state)
	}
	return nil
}

func (t *simpleHSMTest) theEventIsGenerated(event string) error {
	t.sm.HandleEvent(&keyPressEvent{event})
	return nil
}

func (t *simpleHSMTest) theActionIsExecuted(actionID int) error {
	if t.actionsExecuted[actionID] != true {
		return fmt.Errorf("expected action %d to executed, but it was not", actionID)
	}
	return nil
}
