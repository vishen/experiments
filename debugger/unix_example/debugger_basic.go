package main

/*
	It start new process which is traced so it stops before executing first instruction and sends signal to the parent process. Parent process waits for such signal and issues logs like log.Printf("State: %v\n", err). Afterwards process is restarted and parent waits for its termination.
*/

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {

	fmt.Printf("Run %v\n", os.Args[1:])

	cmd := exec.Command(os.Args[1], os.Args[2:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{Ptrace: true}

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Wait()

	log.Printf("State: %v\n", err)
	log.Println("Restarting...")

	cpid := cmd.Process.Pid
	// Get parent group id?
	pgid, err := syscall.Getpgid(cpid)
	if err != nil {
		log.Panic("Error: syscall.Getpdig(1): ", err)
	}

	/*
		PTRACE_O_TRACECLONE:
			-> Stop the tracee at the next clone(2) and automatically start tracing the newly cloned process...
	*/
	if syscall.PtraceSetOptions(cpid, syscall.PTRACE_O_TRACECLONE); err != nil {
		log.Fatal("Error: syscall.PtraceSetOptions(1): ", err)
	}

	// Stops tracee (command passed) after execution of single instruction
	if syscall.PtraceSingleStep(cpid); err != nil {
		log.Panic("Error: syscall.PtraceSingleStep(1): ", err)
	}

	steps := 1

	for {
		var ws syscall.WaitStatus
		/* 	wait4 suspends execution of the current process until a child (as specified by pid) has exited,
		or until a signal is delivered whose action is to terminate the current process or to call a signal handling function. */
		wpid, err := syscall.Wait4(-1*pgid, &ws, syscall.WALL, nil)
		if err != nil {
			log.Fatal("Error: syscall.Wait4(1): ", err)
		}

		if wpid == cpid && ws.Exited() {
			break
		}

		if !ws.Exited() {
			if err := syscall.PtraceSingleStep(wpid); err != nil {
				log.Fatal("Error: syscall.PtraceSingleStep(2): ", err)
			}
			steps += 1
		}
	}
	log.Printf("Steps: %d\n", steps)

}
