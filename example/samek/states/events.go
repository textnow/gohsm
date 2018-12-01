package states

import (
	"strings"
)

type KeyPressEvent struct {
	input string
}

func NewKeyPressEvent(text string) *KeyPressEvent {
	return &KeyPressEvent{
		input: strings.Trim(text, "\n"),
	}
}

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
