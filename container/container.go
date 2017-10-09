package main

// https://lk4d4.darth.io/posts/unpriv4/

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

const (
	containerCloneFlags = syscall.CLONE_NEWUSER | syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET
)

func run(args ...string) error {

	/*
		- PID namespaces isolate the process ID number space
		- CLONE_NEWUSER: Create new user namespace for this process
		- CLONE_NEWPID:
		- CLONE_NEWUTS: Lets you edit the hostname
		- CLONE_NEWNS: Lets you mount /proc/
		- CLONE_NEWNT: Lets you create a network namespace, defaults to no namespace
	*/

	uid := os.Getuid()
	gid := os.Getgid()

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, args...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: containerCloneFlags,
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

	return cmd.Run()
}

func child(args ...string) error {

	// setGgroups()

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

	// Cleanup mounts
	defer func() {
		syscall.Unmount("proc", 0)

	}()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Error running command: %s", err)
	}

	return nil
}

func setCgroups() {

	// Set CGroups, this
	/*
		// https://wiki.archlinux.org/index.php/cgroups
		// https://github.com/docker/libcontainer/tree/master/cgroups/fs
		// https://0xax.gitbooks.io/linux-insides/Cgroups/cgroups1.html
		// http://manpages.ubuntu.com/manpages/zesty/man7/cgroups.7.html
			func cg() {
			cgroups := "/sys/fs/cgroup/"
			pids := filepath.Join(cgroups, "pids")
			os.Mkdir(filepath.Join(pids, "liz"), 0755)
			must(ioutil.WriteFile(filepath.Join(pids, "liz/pids.max"), []byte("20"), 0700))
			// Removes the new cgroup in place after the container exits
			must(ioutil.WriteFile(filepath.Join(pids, "liz/notify_on_release"), []byte("1"), 0700))
			must(ioutil.WriteFile(filepath.Join(pids, "liz/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
		}
	*/
	/*
		cgroups := "/sys/fs/cgroup/"
		pids := filepath.Join(cgroups, "pids")
		path := filepath.Join(pids, "containerr")

		if err := os.MkdirAll(path, 0755); err != nil {
			return err
		}

		if err := ioutil.WriteFile(path+"/pids.max", []byte("1"), 0700); err != nil {
			return err
		}

		if err := ioutil.WriteFile(path+"/notify_on_release", []byte("1"), 0700); err != nil {
			return err
		}

		if err := ioutil.WriteFile(path+"/cgroup.procs", []byte(strconv.Itoa(os.Getpid())), 0700); err != nil {
			return err
		}
	*/
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
