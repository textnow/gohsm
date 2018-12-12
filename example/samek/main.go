package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/Enflick/gohsm"
	"github.com/Enflick/gohsm/example/samek/states"
	"go.uber.org/zap"
	"os"
	"time"
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
	srv := hsm.NewDefaultService(logger)

	startState := states.NewS0State(srv)
	stateMachineEngine := hsm.NewStateMachine(srv, startState, hsm.StartEvent)

	events := make(chan hsm.Event)
	stateMachineEngine.Run(context.TODO(), events)

	handleInput(events)
	logger.Debug("Done\n")
}
