# gohsm

## Introduction
HSM provides the framework for Hierarchical State Machine implementations.

## Related Documents:
 - [Introduction to Hierarchical State Machines](https://barrgroup.com/Embedded-Systems/How-To/Introduction-Hierarchical-State-Machines)
 - [Yet Another Hierarchical State Machine](https://accu.org/index.php/journals/252)
 - [State Diagram](https://en.wikipedia.org/wiki/State_diagram)
 - gohsm Object Model: go_state_machine_framework.png

## Framework Overview
Included in this framework are the following components:

  - **StateMachine**:
    Machine that controls the event processing

  - **State**:
    Interface that must be implemented by all States in the StateMachine

  - **Transition**:
    Interface that is implemented by each of the different types of transitions:

      - **ExternalTransition**:
        Transition from current state to a different state.  On execution the following takes place:
          1. OnExit is called on the current state and all parent states up to the parent state that owns
             the new state (or the parent state is nil)
          2. action() associated with the the transition is called
          3. OnEnter() is called on the new state which may call OnEnter() for a sub-state.  The final
             new current state is returned by the OnEnter() call

      - **InternalTransition**:
        Transition within the current state.  On execution the following takes place:
          1. action() associated with the the transition is called

      - **EndTransition**:
        Transition from current state that terminates the state machine.  On execution the following takes place:
          1. OnExit is called on the current state and all parent states until there are no more parent states
          2. action() associated with the the transition is called

  - **Event**:
    An event represents something that has happened (login, logout, newCall, networkChange, etc.) that might drive
    a change in the state machine

## How To Use
1. Define the set of events that will be processed.
2. Define all of the states that are possible.  For each state, implement the methods required by the State interface
4. Create the state machine, and hand it the starting state
5. Call Run on the state machine

## Example Usage
Two sample HSMs have been implemented in the example/ directory.
