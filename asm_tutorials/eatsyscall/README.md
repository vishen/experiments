# Eat System Call
Basic example of system calls in a 32 bit linux system

## Running
```
$ nasm -f elf -g -F stabs eatsyscall.asm
$ ld -o eatsyscall eatsyscall.o
$ ./eatsyscall
```

OR

```
$ make
```


## Notes
Nasm:
```
- -f elf    # Specifies that the .o will be generated in elf format
- -g        # Specifies that debug information included in .o file
- -F stabs  # Debug information should be in the "stabs" format, other format is dwarf
```
