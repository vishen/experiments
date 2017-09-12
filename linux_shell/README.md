# Shell
A shell is an interface that allows you to interact with the kernel of an operating system.

## How does a shell work
```
- [1] Start shell
- [2] Wait for users input
- [3] Parse users input
- [4] Execute command and return result
- Go back to [2]
```

The shell is the parent process, this is the main thread of out program which will wait for user input. The shell should `fork` any commands given by the user.

## Fork
Fork creates a copy of the current process, but at the same time execution of the parent process is not halted. The copy is known as the child and has a unique `pid`. Both processes are now running the exact same code, with their own stacks.

## Exec
Exec replaces the current process with the command specified after exec. This command does not create a new PID.

This brings us to the exec family of functions. Namely, it has the following functions:
```
- execl
- execv
- execle
- execve
- execlp
- execvp
```
