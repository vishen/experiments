SECTION .data           ; Section containing initialisation data

SECTION .bss            ; Section containing uninitilised data

SECTION .text           ; Section containing code

global _start           ; Linker needs to find the entry point

_start:
    nop                 ; Keeps gdb happy?

    mov     ax, -42

    mov     ecx, eax    ; ecx becomes 65494 because it is moving without considering the signed bit
    movsx   edx, ax     ; edx becomes -42 because it is considerings the signed bit

    nop                 ; Keeps gdb happy?
