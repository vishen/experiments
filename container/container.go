package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func execCommand(binary string, args ...string) error {

	/*
		- PID namespaces isolate the process ID number space
		- CLONE_NEWUSER: Create new user namespace for this process
		- CLONE_NEWPID:
	*/

	uid := os.Getuid()
	gid := os.Getgid()

	cmd := exec.Command(binary, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUSER,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      uid,
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      gid,
				Size:        1,
			},
		},
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	return cmd.Wait()
}

func main() {

	flag.Parse()

	args := flag.Args()

	binary := args[0]
	binaryArgs := args[1:]

	if err := execCommand(binary, binaryArgs...); err != nil {
		log.Fatal(err)
	}
}
