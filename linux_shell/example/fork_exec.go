package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {

	/*binary, lookErr := exec.LookPath("ls")
	  if lookErr != nil {
	      panic(lookErr)
	  }*/

	workingDir := "/"
	environ := os.Environ()
	files := []uintptr{uintptr(syscall.Stdin), uintptr(syscall.Stdout), uintptr(syscall.Stderr)}

	procAttr := syscall.ProcAttr{Dir: workingDir, Env: environ, Files: files}

	childPid, err := syscall.ForkExec("/bin/ls", []string{"-a", "-h", "-l"}, &procAttr)
	if err != nil {
		panic(err)
	}

	pid := os.Getpid()
	parentPid := os.Getppid()

	time.Sleep(5 * time.Second)
	fmt.Printf("pid=%d, childPid=%d, parentPid=%d\n", pid, childPid, parentPid)
}
