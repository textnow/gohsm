package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rmrobinson-textnow/gohsm"
	"go.uber.org/zap"
)

func action1() {
	fmt.Printf("\nAction1\n")
}
func action2() {
	fmt.Printf("\nAction2\n")
}
func action3() {
	fmt.Printf("\nAction3\n")
}
func action4() {
	fmt.Printf("\nAction4\n")
}
func action5() {
	fmt.Printf("\nAction5\n")
}
func action6() {
	fmt.Printf("\nAction6\n")
}
func action7() {
	fmt.Printf("\nAction7\n")
}

type stateA struct {
	A bool
}

func (a *stateA) OnEnter(hsm.Event) {
	fmt.Printf("->A;")
}
func (a *stateA) OnExit(hsm.Event) {
	fmt.Printf("<-A;")
}

type stateB struct {
}

func (b *stateB) OnEnter(hsm.Event) {
	fmt.Printf("->B;")
}
func (b *stateB) OnExit(hsm.Event) {
	fmt.Printf("<-B;")
}

type stateC struct {
	hsm.LeafState
}

func (c *stateC) OnEnter(hsm.Event) {
	fmt.Printf("->C;")
}
func (c *stateC) OnExit(hsm.Event) {
	fmt.Printf("<-C;")
}

type stateD struct {
}

func (d *stateD) OnEnter(hsm.Event) {
	fmt.Printf("->D;")
}
func (d *stateD) OnExit(hsm.Event) {
	fmt.Printf("<-D;")
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
	sd := &stateD{}
	sc := &stateC{}
	sb := &stateB{}
	sa := &stateA{
		A: true,
	}

	ea := &keyPressEvent{"a"}
	eb := &keyPressEvent{"b"}
	ec := &keyPressEvent{"c"}
	ed := &keyPressEvent{"d"}
	ee := &keyPressEvent{"e"}
	ex := &keyPressEvent{"x"}
	ey := &keyPressEvent{"y"}

	c := hsm.NewLeafState("C", sc)
	c.SetExternalTransition(ex, action6, hsm.NewDirectTransition(c))
	c.SetInternalTransition(ey, action7)

	b := hsm.NewLeafState("B", sb)
	b.SetExternalTransition(ea, action1, hsm.NewDirectTransition(c))

	d := hsm.NewLeafState("D", sd)
	d.SetExternalTransition(ee, action5, hsm.EndTransition)

	a := hsm.NewCompositeStateWithTransition("A", sa, []hsm.State{
		b,
		c,
	}, hsm.NewConditionalTransition(func() hsm.State {
		if sa.A {
			return b
		}
		return c
	}))

	a.SetInternalTransition(eb, action2)
	a.SetExternalTransition(ec, action3, hsm.NewDirectTransition(a))
	a.SetExternalTransition(ed, action4, hsm.NewDirectTransition(d))

	root := hsm.NewCompositeState("", &hsm.EmptyHandler{}, []hsm.State{
		a,
		d,
	})

	engine := hsm.NewStateMachineEngine(zap.L(), hsm.NewDirectTransition(root))

	events := make(chan hsm.Event)
	go handleInput(events)
	engine.Run(events)
	fmt.Printf("Done\n")
}
