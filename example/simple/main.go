package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/Enflick/gohsm/example/simple/states"
	"os"
	"time"

	"github.com/Enflick/gohsm"
	"go.uber.org/zap"
)

func handleInput(events chan hsm.Event) {
	for {
		// Wait for the output to stabilize before proceeding
		time.Sleep(time.Millisecond * 100)
		fmt.Print("\nEnter text: ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		fmt.Println(text)

		if text == "done\n" {
			break
		}

		events <- states.NewKeyPressEvent(text)
	}
}

func main() {
	logger, _ := zap.NewDevelopment()
	startState := states.NewStateA(logger, true)
	stateMachineEngine := hsm.NewStateMachine(logger, startState, hsm.StartEvent)

	events := make(chan hsm.Event)
	stateMachineEngine.Run(context.TODO(), events)

	handleInput(events)
	logger.Debug("Done\n")
}
