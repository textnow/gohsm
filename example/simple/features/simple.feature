Feature: transition states
  In order to be successful
  As a hierarchical state machine
  I need to be able to execute common state transitions

  Scenario: Can be initialized
    Given test field A is true
    And the simple hierarchical state machine is setup
    Then the state "A" should be entered
    And the state "B" should be entered
    And the current state should be "B"

  Scenario: Can perform an external transition in an internal state
    Given test field A is true
    And the simple hierarchical state machine is setup
    When event "a" is generated
    Then the current state should be "C"
    And the action with ID 1 should be executed
    And the state "C" should be entered
    And the state "B" should be exited

  Scenario: Can perform an internal transition on an outer state
    Given the simple hierarchical state machine is setup
    When event "b" is generated
    Then the action with ID 2 should be executed
    And the state "C" should not be exited
    And the current state should be "C"

  Scenario: Can perform an external transition on an outer state
    Given the simple hierarchical state machine is setup
    When event "c" is generated
    Then the state "C" should be exited
    And the state "A" should be exited
    And the action with ID 3 should be executed
    And the state "A" should be entered
    And the state "C" should be entered
    And the current state should be "C"

  Scenario: Can perform an external transition a state returning to the same state
    Given the simple hierarchical state machine is setup
    When event "x" is generated
    Then the state "C" should be exited
    Then the action with ID 6 should be executed
    Then the state "C" should be entered
    Then the current state should be "C"

  Scenario: Can perform an internal transition on an inner state
    Given the simple hierarchical state machine is setup
    And the current state should be "C"
    When event "y" is generated
    Then the action with ID 7 should be executed
    And the state "C" should not be exited
    And the current state should be "C"

  Scenario: Can perform an external transition on an outer leaving an inner state
    Given test field A is true
    And the simple hierarchical state machine is setup
    When event "d" is generated
    Then the state "B" should be exited
    And the state "A" should be exited
    And the action with ID 4 should be executed
    And the state "D" should be entered
    And the current state should be "D"