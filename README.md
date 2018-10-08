# gohsm

This library provides a Golang implementation of a Hierarchical State Machine (HSM). This library uses several sources for implementation specification, including:
 - [Miro Samek](https://barrgroup.com/Embedded-Systems/How-To/Introduction-Hierarchical-State-Machines)
 - [Stefan Heinzmann](https://accu.org/index.php/journals/252)

Two sample HSMs have been implemented in the example/ directory.

The HSM implemented here contains 3 key types:
 1. the State interface
 2. the Transition interface
 3. the StateMachineEngine

The StateMachineEngine represents a single instance of a running state machine; it consists of a run loop which consumes events from the supplied 'events' channel and passes them to the HSM for processing.

The State interface is implemented by two concrete implementations of a State:
 1. the LeafState
 2. the CompositeState

A LeafState instance represents a state with no sub-states; a CompositeState instance represents a state with sub-states. Both LeafStates and CompositeStates can have transitions into and out of them. During operation (when not processing an event) the current state of the StateMachineEngine is always a LeafState; CompositeStates are processed along the way but are never 'resting' states. As events are received the set of registered transitions on the current state are examined; if a registered transition is found the HSM will (potentially) traverse up the hierarchy of the state machine to find the common parent and then descend the hierarchy to the target of the transition. Along the way the parent states left will have their OnExit methods invoked; as we descend to the target we will invoke the OnEnter methods of each state found. As the HSM is always resting in a LeafState we will continue down until a LeafState is reached (calling any OnEnters necessary).

Transitions represent different ways to get from one state to another. The default transition is a direct transition, which simply goes from the origin state to the destination state. If more complex logic is required, a conditional transition can be used. Conditional transitions will invoke a method supplied during initialization to determine which state to transition to.

This library does not currently implement guard conditions on transitions.
