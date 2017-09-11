SECTION .data           ; Section containing initialisation data

snippet db "KANGAROO"   ; Stored backwards because BigEndian

SECTION .bss            ; Section containing uninitilised data

SECTION .text           ; Section containing code

global _start           ; Linker needs to find the entry point

; Global functions
Print:
    ; Print the current 'snippet'
    mov eax, 4
    mov ebx, 1
    mov ecx, snippet
    mov edx, 9
    int 80H
    ret

; Main
_start:
    nop                 ; Keeps gdb happy?

    call Print

    mov ebx, snippet
    mov eax, 8          ; Length of string
DoMore:
    add byte [ebx], 32  ; Adds 32 to the number being pointed to by ebx
    inc ebx
    dec eax
    jnz DoMore

    call Print

    nop                 ; Keeps gdb happy?
