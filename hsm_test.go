package hsm

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
	"testing"
)

type resolveLeafTest struct {
	name string
	tree State
	expectedResults []State
}

var leafState = &LeafState{}
var leafState2 = &LeafState{}
var compositeState = NewCompositeState("test", nil, []State{leafState})
var superCompositeState = NewCompositeState("test2", nil, []State{compositeState})
var badCompositeState = &CompositeState{LeafState: LeafState{eventHandlers: map[string]*eventHandler{}}}

var parentState = NewCompositeState("test3", nil, []State{leafState, leafState2})

var resolveLeafTests = []resolveLeafTest{
	{
		name: "basic test, base is leaf",
		tree: &LeafState{},
		expectedResults: nil,
	},
	{
		name: "basic test, leaf with single parent",
		tree: compositeState,
		expectedResults: []State{leafState},
	},
	{
		name: "basic test, leaf with two parents",
		tree: superCompositeState,
		expectedResults: []State{compositeState, leafState},
	},
	{
		name: "invalid case, composite state with no handlers",
		tree: badCompositeState,
		expectedResults: nil,
	},
}

func TestResolveLeaf(t *testing.T) {
	for _, tt := range resolveLeafTests {
		t.Run(tt.name, func(t *testing.T) {
			logger := zaptest.NewLogger(t)
			sme := &StateMachineEngine{
				logger: logger,
			}

			leaves := sme.resolveLeaf(tt.tree)
			assert.Equal(t, tt.expectedResults, leaves)
		})
	}
}

type resolveTransitionTest struct {
	name string
	origin State
	dest State
	toLeave []State
	toEnter []State
}

var s21 = &LeafState{name: "s21"}
var s2 = NewCompositeState("s2", nil, []State{s21})
var s11 = &LeafState{name: "s11"}
var s1 = NewCompositeState("s1", nil, []State{s11})
var s0 = NewCompositeState("s0", nil, []State{s1, s2})
var s3 = &LeafState{name: "s3"}

var resolveTransitionTests = []resolveTransitionTest{
	{
		name: "invalid transition (origin)",
		origin: nil,
		dest: leafState,
		toLeave: nil,
		toEnter: nil,
	},
	{
		name: "invalid transition (dest)",
		origin: leafState,
		dest: nil,
		toLeave: nil,
		toEnter: nil,
	},
	{
		name: "self transition",
		origin: leafState,
		dest: leafState,
		toLeave: []State{leafState},
		toEnter: []State{leafState},
	},
	{
		name: "common parent",
		origin: leafState,
		dest: leafState2,
		toLeave: []State{leafState},
		toEnter: []State{leafState2},
	},
	{
		name: "transition to end state",
		origin: leafState,
		dest: EndState,
		toLeave: []State{leafState},
		toEnter: []State{},
	},
	{
		name: "s0 is common parent",
		origin: s11,
		dest: s21,
		toLeave: []State{s11, s1, s0},
		toEnter: []State{s0, s2, s21},
	},
	{
		name: "no common parent",
		origin: s11,
		dest: s3,
		toLeave: []State{s11, s1, s0},
		toEnter: []State{s3},
	},
}

func TestResolveTransition(t *testing.T) {
	for _, tt := range resolveTransitionTests {
		t.Run(tt.name, func(t *testing.T) {
			logger := zaptest.NewLogger(t)
			sme := &StateMachineEngine{
				logger: logger,
			}

			toLeave, toEnter := sme.resolveTransition(tt.origin, tt.dest)
			assert.Equal(t, tt.toLeave, toLeave)
			assert.Equal(t, tt.toEnter, toEnter)
		})
	}
}
