package hsm

type StateEngine struct {
	state             *State
	parentStateEngine *StateEngine
}

func NewStateEngine(state State, parentStateEngine *StateEngine) *StateEngine {
	stateEngine := &StateEngine{
		state:             &state,
		parentStateEngine: parentStateEngine,
	}

	return stateEngine
}

func (a *StateEngine) Name() string {
	return a.GetState().Name()
}

func (a *StateEngine) GetState() State {
	return *a.state
}

func (a *StateEngine) GetParentStateEngine() *StateEngine {
	return a.parentStateEngine
}
