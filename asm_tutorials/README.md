# ASM Learnding

## GDB
Basics of running a binary in GDB


```
$ gdb <binary>
(gdb) info files
(gdb) info proc mappings
(gdb) info break
(gdb) break <filename>:<line number>
(gdb) run
(gdb) info reg
(gdb) x <memory location>
(gdb) x/<repeats:int><format:(x=hex|i=machine code|u=unsigned int)><size:(b=bytes|h=two bytes|w=four bytes|g=eight bytes> <memory location>
(gdb) x/1xb <addr>
(gdb) x/1ub <addr>
(gdb) x/1xg <addr>
(gdb) info frame
(gdb) n
```


