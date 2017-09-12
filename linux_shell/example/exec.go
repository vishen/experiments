package main

import (
	"context"
	"fmt"
	"os/exec"
)

func main() {

	ctx := context.Background()

	cmd := exec.CommandContext(ctx, "false")

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	fmt.Printf("Command: %#v\n", cmd)

	err := cmd.Wait()
	fmt.Printf("Command finished with error: %v", err)
}
