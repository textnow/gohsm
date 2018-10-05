package main

import (
	"fmt"
	"time"

	"github.com/rmrobinson-textnow/gohsm"
)

func action1() {
	fmt.Printf("Action1\n")
}
func action2() {
	fmt.Printf("Action2\n")
}
func action3() {
	fmt.Printf("Action3\n")
}
func action4() {
	fmt.Printf("Action4\n")
}
func action5() {
	fmt.Printf("Action5\n")
}
func action6() {
	fmt.Printf("Action6\n")
}
func action7() {
	fmt.Printf("Action7\n")
}

type eventA struct {
}

func (a *eventA) ID() string {
	return "a"
}

type eventB struct {
}

func (b *eventB) ID() string {
	return "b"
}

type eventC struct {
}

func (c *eventC) ID() string {
	return "c"
}

type eventD struct {
}

func (d *eventD) ID() string {
	return "d"
}

type eventE struct {
}

func (e *eventE) ID() string {
	return "e"
}

type eventX struct {
}

func (x *eventX) ID() string {
	return "x"
}

type eventY struct {
}

func (y *eventY) ID() string {
	return "y"
}

type stateA struct {
	A bool
}

func (a *stateA) OnEnter(hsm.Event) {
	fmt.Printf("->A\n")
}
func (a *stateA) OnExit(hsm.Event) {
	fmt.Printf("<-A\n")
}

type stateB struct {
}

func (b *stateB) OnEnter(hsm.Event) {
	fmt.Printf("->B\n")
}
func (b *stateB) OnExit(hsm.Event) {
	fmt.Printf("<-B\n")
}

type stateC struct {
}

func (c *stateC) OnEnter(hsm.Event) {
	fmt.Printf("->C\n")
}
func (c *stateC) OnExit(hsm.Event) {
	fmt.Printf("<-C\n")
}

type stateD struct {
}

func (d *stateD) OnEnter(hsm.Event) {
	fmt.Printf("->D\n")
}
func (d *stateD) OnExit(hsm.Event) {
	fmt.Printf("<-D\n")
}

func main() {
	sd := &stateD{}
	sc := &stateC{}
	sb := &stateB{}
	sa := &stateA{
		A: true,
	}

	ea := &eventA{}
	eb := &eventB{}
	ec := &eventC{}
	ed := &eventD{}
	ee := &eventE{}
	ex := &eventX{}
	ey := &eventY{}

	c := hsm.NewState("C", sc)
	c.SetExternalTransition(ex, action6, hsm.NewDirectTransition(c))
	c.SetInternalTransition(ey, action7)

	b := hsm.NewState("B", sb)
	b.SetExternalTransition(ea, action1, hsm.NewDirectTransition(c))

	smaChild := hsm.NewStateMachine(hsm.NewConditionalTransition(func() *hsm.State {
		if sa.A {
			return b
		} else {
			return c
		}
	}))

	d := hsm.NewState("D", sd)
	d.SetExternalTransition(ee, action5, hsm.EndTransition)

	a := hsm.NewStateWithSubStateMachine("A", sa, smaChild)
	a.SetInternalTransition(eb, action2)
	a.SetExternalTransition(ec, action3, hsm.NewDirectTransition(a))
	a.SetExternalTransition(ed, action4, hsm.NewDirectTransition(d))

	sma := hsm.NewStateMachine(hsm.NewDirectTransition(a))

	engine := hsm.NewStateMachineEngine(sma)

	events := make(chan hsm.Event, 20)
	events <- ea
	events <- eb
	events <- ec
	events <- ea
	events <- ex
	events <- ey

	go func() {
		time.Sleep(time.Second)
		close(events)
	}()

	engine.Run(events)
	fmt.Printf("Done\n")
}
