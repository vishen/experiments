package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func run(args ...string) error {

	/*
		- PID namespaces isolate the process ID number space
		- CLONE_NEWUSER: Create new user namespace for this process
		- CLONE_NEWPID:
		- CLONE_NEWUTS: Lets you edit the hostname
	*/

	uid := os.Getuid()
	gid := os.Getgid()

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, args...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUSER | syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
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

	/*if err := cmd.Start(); err != nil {
		return err
	}*/

	return cmd.Run()
}

func child(args ...string) error {

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := syscall.Sethostname([]byte("containerr")); err != nil {
		return fmt.Errorf("Couldn't set hostname of child: %s", err)
	}

	if err := os.Chdir("/"); err != nil {
		return fmt.Errorf("Unable to chdir to '/': %s", err)
	}

	if err := syscall.Mount("proc", "proc", "proc", 0, ""); err != nil {
		return fmt.Errorf("Error mounting /proc: %s", err)
	}

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Error running command: %s", err)
	}

	// Set CGroups
	/*
			func cg() {
			cgroups := "/sys/fs/cgroup/"
			pids := filepath.Join(cgroups, "pids")
			os.Mkdir(filepath.Join(pids, "liz"), 0755)
			must(ioutil.WriteFile(filepath.Join(pids, "liz/pids.max"), []byte("20"), 0700))
			// Removes the new cgroup in place after the container exits
			must(ioutil.WriteFile(filepath.Join(pids, "liz/notify_on_release"), []byte("1"), 0700))
			must(ioutil.WriteFile(filepath.Join(pids, "liz/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
		}
			cgroups := "/sys/fs/cgroup/"
			pids := filepath.Join(cgroups, "pids")
			os.Mkdir(filepath.Join(pids, "containerr"), 0755)
	*/

	syscall.Unmount("proc", 0)

	return nil
}

func main() {

	flag.Parse()
	args := flag.Args()

	switch args[0] {
	case "child":
		if err := child(args[1:]...); err != nil {
			log.Fatalf("Error executing child '%v': %s", args[1:], err)
		}
	case "run":
		if err := run(args[1:]...); err != nil {
			log.Fatalf("Error executing run '%v': %s", args[1:], err)
		}
	default:
		log.Fatalf("Unknown command: %s", args[0])
	}

}
