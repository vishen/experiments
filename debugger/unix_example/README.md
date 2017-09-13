# Debugger
A very basic debugger written in Go. Currently you can only step through a linux binary inspecting current register state at each point.

Only works on linux as it uses Ptrace, and Ptrace support on mac is shit.

## Running
```
$ make build_debugger
$ ./debugger <path/to/binary>
```

## Commands
```
- next, step -> proceeds to next instruction
- state, print -> prints registers to screen
```

## TODO
```
- Breakpoints
- Continue
- Incorporate elf debugging info?
```
