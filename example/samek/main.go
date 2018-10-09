package main

import (
	"bufio"
	"fmt"
	"github.com/rmrobinson-textnow/gohsm"
	"go.uber.org/zap"
	"os"
	"strings"
	"time"
)

type stateS0 struct {
}

func (a *stateS0) OnEnter(hsm.Event) {
	fmt.Printf("->S0;")
}
func (a *stateS0) OnExit(hsm.Event) {
	fmt.Printf("<-S0;")
}

type stateS1 struct {
}

func (b *stateS1) OnEnter(hsm.Event) {
	fmt.Printf("->S1;")
}
func (b *stateS1) OnExit(hsm.Event) {
	fmt.Printf("<-S1;")
}

type stateS11 struct {
}

func (c *stateS11) OnEnter(hsm.Event) {
	fmt.Printf("->S11;")
}
func (c *stateS11) OnExit(hsm.Event) {
	fmt.Printf("<-S11;")
}

type stateS2 struct {
}

func (d *stateS2) OnEnter(hsm.Event) {
	fmt.Printf("->S2;")
}
func (d *stateS2) OnExit(hsm.Event) {
	fmt.Printf("<-S2;")
}

type stateS21 struct {
}

func (d *stateS21) OnEnter(hsm.Event) {
	fmt.Printf("->S21;")
}
func (d *stateS21) OnExit(hsm.Event) {
	fmt.Printf("<-S21;")
}

type stateS211 struct {
}

func (d *stateS211) OnEnter(hsm.Event) {
	fmt.Printf("->S211;")
}
func (d *stateS211) OnExit(hsm.Event) {
	fmt.Printf("<-S211;")
}

type keyPressEvent struct {
	input string
}

func (kpe *keyPressEvent) ID() string {
	return kpe.input
}

func handleInput(events chan hsm.Event) {
	for {
		// Wait for the output to stabilize before proceeding
		time.Sleep(time.Millisecond * 100)
		fmt.Print("\nEnter text: ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		fmt.Println(text)

		events <- &keyPressEvent{
			input: strings.Trim(text, "\n"),
		}
	}
}

func main() {
	eventA := &keyPressEvent{"a"}
	eventB := &keyPressEvent{"b"}
	eventC := &keyPressEvent{"c"}
	eventD := &keyPressEvent{"d"}
	eventE := &keyPressEvent{"e"}
	eventF := &keyPressEvent{"f"}
	eventG := &keyPressEvent{"g"}
	eventH := &keyPressEvent{"h"}

	s211 := hsm.NewLeafState("s211", &stateS211{})
	s21 := hsm.NewCompositeState("s21", &stateS21{}, []hsm.State{
		s211,
	})
	s2 := hsm.NewCompositeState("s2", &stateS2{}, []hsm.State{
		s21,
	})

	s11 := hsm.NewLeafState("s11", &stateS11{})
	s1 := hsm.NewCompositeState("s1", &stateS1{}, []hsm.State{
		s11,
	})

	s0 := hsm.NewCompositeState("s0", &stateS0{}, []hsm.State{
		s1,
		s2,
	})

	s211.SetExternalTransition(eventD, func() { fmt.Printf("\nEventD\n") }, hsm.NewDirectTransition(s21))
	s211.SetExternalTransition(eventG, func() { fmt.Printf("\nEventG\n") }, hsm.NewDirectTransition(s0))

	s21.SetExternalTransition(eventB, func() { fmt.Printf("\nEventB\n") }, hsm.NewDirectTransition(s211))
	s21.SetExternalTransition(eventH, func() { fmt.Printf("\nEventH\n") }, hsm.NewDirectTransition(s21))

	s2.SetExternalTransition(eventC, func() { fmt.Printf("\nEventC\n") }, hsm.NewDirectTransition(s1))
	s2.SetExternalTransition(eventF, func() { fmt.Printf("\nEventF\n") }, hsm.NewDirectTransition(s11))

	s11.SetExternalTransition(eventG, func() { fmt.Printf("\nEventG\n") }, hsm.NewDirectTransition(s211))
	s11.SetInternalTransition(eventH, func() { fmt.Printf("\nEventH\n") })

	s1.SetExternalTransition(eventA, func() { fmt.Printf("\nEventA\n") }, hsm.NewDirectTransition(s1))
	s1.SetExternalTransition(eventB, func() { fmt.Printf("\nEventB\n") }, hsm.NewDirectTransition(s11))
	s1.SetExternalTransition(eventC, func() { fmt.Printf("\nEventC\n") }, hsm.NewDirectTransition(s2))
	s1.SetExternalTransition(eventD, func() { fmt.Printf("\nEventD\n") }, hsm.NewDirectTransition(s0))
	s1.SetExternalTransition(eventF, func() { fmt.Printf("\nEventF\n") }, hsm.NewDirectTransition(s211))

	s0.SetExternalTransition(eventE, func() { fmt.Printf("\nEventE\n") }, hsm.NewDirectTransition(s211))

	sme := hsm.NewStateMachineEngine(zap.L(), hsm.NewDirectTransition(s0))

	events := make(chan hsm.Event)
	go handleInput(events)
	sme.Run(events)
}
