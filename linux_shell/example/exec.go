package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
)

func exit() {
	fmt.Println("Exiting goshell...")
	os.Exit(0)
}

func main() {

	/* TODO(vishen):
	- Handle pipes
	- Add custom commands; cd
	*/

	// Capture any interupts
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Listen for the interrupt and exit
	go func() {
		<-c
		exit()
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("goshell> ")

		command, err := reader.ReadString('\n')

		// Happens when you hit ctrl+d, which sends an EOF
		if command == "" {
			break
		}

		if err != nil {
			fmt.Printf("Error with command '%s': %s\n", command, err)
			continue
		}

		// Remove the '\n' from the end
		command = command[0 : len(command)-1]

		if command == "exit" || command == "quit" {
			break
		}

		commandSplit := strings.Split(command, " ")

		ctx := context.Background()

		cmd := exec.CommandContext(ctx, commandSplit[0], commandSplit[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr

		if err := cmd.Start(); err != nil {
			fmt.Printf("Error starting command '%s': %s\n", command, err)
		}

		if err := cmd.Wait(); err != nil {
			fmt.Printf("Command finished with error: %v\n", err)
		}
	}

	exit()
}
