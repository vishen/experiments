package main

/*
	Requires libseccomp v2.2 installed if using "github.com/seccomp/libseccomp-golang" from master (super annoying as `sudo apt-get install libseccomp2-dev` installs libseccomp v2.1 )
	BUILD: `go build -tags "seccomp" -o strace strace.go`
*/

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	sec "github.com/seccomp/libseccomp-golang"
)

func main() {
	fmt.Printf("Run %v\n", os.Args[1:])

	cmd := exec.Command(os.Args[1], os.Args[2:]...)

	/*
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
	*/

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Ptrace: true,
	}

	cmd.Start()

	fmt.Println(cmd.Process.Pid)

	err := cmd.Wait()

	fmt.Printf("State: %v\n", err)
	fmt.Printf("Restarting...\n")

	pid := cmd.Process.Pid

	print := true

	for {

		if print {
			var regs syscall.PtraceRegs
			if err := syscall.PtraceGetRegs(pid, &regs); err != nil {
				fmt.Printf("Error getting Ptrace Regs: %s\n", err)
				break
			}

			printPeekData := func(pid int, memoryAddress uint64, size uint64) {
				out := make([]byte, size)
				if _, err := syscall.PtracePeekData(pid, uintptr(memoryAddress), out); err != nil {
					fmt.Printf("Error PtracePeekData(...): %s\n", err)
				}

				firstNullByte := 0
				for i, ch := range out {
					if ch == 0 {
						firstNullByte = i
						break
					}
				}
				fmt.Println("Data at memory location: ", string(out[0:firstNullByte]))
			}

			orax := regs.Orig_rax

			name, _ := sec.ScmpSyscall(orax).GetName()
			fmt.Printf("(%s) origRax=%d rax=%d rdi=%d rsi=%d rdx=%d rbx=%d\n", name, orax, regs.Rax, regs.Rdi, regs.Rsi, regs.Rdx, regs.Rbx)

			// https://github.com/torvalds/linux/blob/master/arch/x86/entry/syscalls/syscall_64.tbl
			// http://blog.rchapman.org/posts/Linux_System_Call_Table_for_x86_64/
			switch orax {
			case 1: // sys_write
				printPeekData(pid, regs.Rsi, regs.Rdx)
			case 2: // sys_open
				printPeekData(pid, regs.Rdi, 256)
			case 3: // sys_close
				fmt.Printf("Closing: %d\n", regs.Rdi)
			}
		}
		err = syscall.PtraceSyscall(pid, 0)
		if err != nil {
			panic(err)
		}

		_, err = syscall.Wait4(pid, nil, 0, nil)
		if err != nil {
			panic(err)
		}

		print = !print
	}
}
