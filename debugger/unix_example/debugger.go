package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func step(pid int) error {
	return syscall.PtraceSingleStep(pid)
}

func cont(pid int) error {
	return syscall.PtraceCont(pid, 0)
}

func getRegisters(pid int) (syscall.PtraceRegs, error) {
	var regs syscall.PtraceRegs
	err := syscall.PtraceGetRegs(pid, &regs)

	if err != nil {
		return regs, err
	}

	return regs, nil
}

func setPC(pid int, pc uint64) error {
	regs, err := getRegisters(pid)
	if err != nil {
		return err
	}

	regs.SetPC(pc)

	if err := syscall.PtraceSetRegs(pid, &regs); err != nil {
		return err
	}

	return nil
}

func getPC(pid int) (uint64, error) {
	regs, err := getRegisters(pid)
	if err != nil {
		return 0, err
	}

	return regs.PC(), nil
}

func setBreakpoint(pid int, breakpoint uintptr) ([]byte, error) {
	// The INT 3 instruction generates a special one byte opcode (CC) that is intended for calling the debug exception handler
	// We need to get and return the current instruction while overwriting with 0xCC
	original := make([]byte, 1)
	if _, err := syscall.PtracePeekData(pid, breakpoint, original); err != nil {
		return nil, err
	}

	if _, err := syscall.PtracePokeData(pid, breakpoint, []byte{0xCC}); err != nil {
		return nil, err
	}

	return original, nil
}

func cleanBreakpoint(pid int, breakpoint uintptr, original []byte) error {
	if _, err := syscall.PtracePokeData(pid, breakpoint, original); err != nil {
		return err
	}

	return nil
}

func printState(pid int) error {
	regs, err := getRegisters(pid)
	if err != nil {
		return err
	}

	// fmt.Printf("%#v\n", regs)
	// orax := regs.Orig_rax
	/*
		switch orax {
				case 1: // sys_write
					printPeekData(pid, regs.Rsi, regs.Rdx)
				case 2: // sys_open
					printPeekData(pid, regs.Rdi, 256)
				case 3: // sys_close
					fmt.Printf("Closing: %d\n", regs.Rdi)
				}
	*/

	fmt.Printf("rax=%d rdi=%d rsi=%d rdx=%d rbx=%d\n", regs.Rax, regs.Rdi, regs.Rsi, regs.Rdx, regs.Rbx)

	return nil
}

func main() {
	cmd := exec.Command(os.Args[1], os.Args[2:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Ptrace: true,
	}

	if err := cmd.Start(); err != nil {
		log.Fatalf("Error: cmd.Start(1): %s\n", err)
	}

	err := cmd.Wait()

	cpid := cmd.Process.Pid

	fmt.Printf("Current state for %d: %v\n", cpid, err)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("godebugger> ")

		command, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error: reader.ReadString(1): %s\n", err)
		}

		command = command[:len(command)-1]

		if command == "state" || command == "print" {
			if err := printState(cpid); err != nil {
				log.Printf("Error: printState(1): %s\n", err)
			}
		} else if command == "step" || command == "next" {
			if err := step(cpid); err != nil {
				log.Printf("Error: step(1): %s\n", err)
			}
		} else {
			log.Printf("Uknown command '%s'\n", command)
		}
	}
}
