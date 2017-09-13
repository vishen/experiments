package main

/*
	This doesn't work because Apple requires your program to be a trusted entity when using `ptrace`
	functionality
*/

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	fmt.Printf("Run %v\n", os.Args[1:])

	cmd := exec.Command(os.Args[1], os.Args[2:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Ptrace: true,
	}

	cmd.Start()

	fmt.Println(cmd.Process.Pid)

	err := cmd.Wait()

	fmt.Printf("State: %v\n", err)
	fmt.Printf("Restarting...\n")

	// Look at below for darwin Ptrace capabilities + getting registers
	// https://github.com/derekparker/delve/blob/b6fe5aebaf1ea0d0b5b5f03525fa023caa81489c/pkg/proc/native/ptrace_darwin.go
	// https://github.com/derekparker/delve/blob/b6fe5aebaf1ea0d0b5b5f03525fa023caa81489c/pkg/proc/native/registers_darwin_amd64.go

	pid := cmd.Process.Pid

	if err := syscall.PtraceAttach(pid); err != nil {
		fmt.Printf("Error with PtraceAttach(..): %s\n", err)
	}
}
