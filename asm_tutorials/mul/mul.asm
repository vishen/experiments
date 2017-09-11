SECTION .data           ; Section containing initialisation data

SECTION .bss            ; Section containing uninitilised data

SECTION .text           ; Section containing code

global _start           ; Linker needs to find the entry point

_start:
    nop                 ; Keeps gdb happy?

    mov eax, 11
    mov ebx, 12
    mul ebx             ; stores result in 'eax'

    mov eax, 0FFFFFFFFh
    mov ebx, 03B72H
    mul ebx             ; stores one part of the result in 'eax' and the other in 'ebx'

    nop                 ; Keeps gdb happy?
