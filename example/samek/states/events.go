package states

import (
	"strings"
)

// KeyPressEvent represents an keyboard entered event
type KeyPressEvent struct {
	input string
}

// NewKeyPressEvent constructor
func NewKeyPressEvent(text string) *KeyPressEvent {
	return &KeyPressEvent{
		input: strings.Trim(text, "\n"),
	}
}

// ID returns the event's id
func (kpe *KeyPressEvent) ID() string {
	return kpe.input
}

var ea = &KeyPressEvent{"a"}
var eb = &KeyPressEvent{"b"}
var ec = &KeyPressEvent{"c"}
var ed = &KeyPressEvent{"d"}
var ee = &KeyPressEvent{"e"}
var ef = &KeyPressEvent{"f"}
var eg = &KeyPressEvent{"g"}
var eh = &KeyPressEvent{"h"}
