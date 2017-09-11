SECTION .data           ; Section containing initialisation data

SECTION .bss            ; Section containing uninitilised data

SECTION .text           ; Section containing code

global _start           ; Linker needs to find the entry point

_start:
    nop                 ; Keeps gdb happy?

    nop                 ; Keeps gdb happy?
