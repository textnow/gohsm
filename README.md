# gohsm

## Introduction

This library provides a Golang implementation of a Hierarchical State Machine (HSM). This library uses several sources for implementation specification, including:
 - [Miro Samek](https://barrgroup.com/Embedded-Systems/How-To/Introduction-Hierarchical-State-Machines)
 - [Stefan Heinzmann](https://accu.org/index.php/journals/252)

Two sample HSMs have been implemented in the example/ directory.

The HSM implemented here contains 3 key types:
 1. the State interface
 2. the Transition interface
 3. the StateMachineEngine

The StateMachineEngine represents a single instance of a running state machine; it consists of a run loop which consumes events from the supplied 'events' channel and passes them to the HSM for processing.

All states in a State Machine must implement the State interface.  A state can be a
Composite State (contains sub-states) or a Leaf State (does not contain sub-states). 

This library does not currently implement guard conditions on transitions.

## How To Use

1. Define the set of events that will be processed.
2. Define all of the states that are possible.  For each state, implement the methods required by the State interface
4. Create the state machine engine, and hand it the starting state
5. Call Run on the state machine engine
