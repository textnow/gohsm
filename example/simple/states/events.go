package states

import (
	"strings"
)

// KeyPressEvent represents a keyboard entered event
type KeyPressEvent struct {
	input string
}

// NewKeyPressEvent constructor
func NewKeyPressEvent(text string) *KeyPressEvent {
	return &KeyPressEvent{
		input: strings.Trim(text, "\n"),
	}
}

// ID returns the events identifier
func (kpe *KeyPressEvent) ID() string {
	return kpe.input
}

var ea = &KeyPressEvent{"a"}
var eb = &KeyPressEvent{"b"}
var ec = &KeyPressEvent{"c"}
var ed = &KeyPressEvent{"d"}
var ee = &KeyPressEvent{"e"}
var ey = &KeyPressEvent{"y"}
var ex = &KeyPressEvent{"x"}

// LastActionIDExecuted helps with testing
var LastActionIDExecuted = 0
