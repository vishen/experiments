SECTION .data           ; Section containing initialisation data

SECTION .bss            ; Section containing uninitilised data
    buff resb 1

SECTION .text           ; Section containing code
    global _start       ; Linker needs to find the entry point

_start:
    nop                 ; Keeps gdb happy?

read:
    mov eax, 3          ; Specify sys_read call
    mov ebx, 0          ; stdin
    mov ecx, buff       ; Pass the address of the buffer to read to
    mov edx, 1          ; Read 1 character from stdin
    int 80h

    cmp eax, 0          ; sys_read return value; 0 indicates EOF
    je exit             ; Exit if EOF

    cmp byte [buff], 61h    ; 61h (hex) -> 'a'
    jb write                ; if below 'a' we don't need to do anything

    cmp byte [buff], 7ah    ; 7ah -> 'z'
    ja write                ; if above 'z' we don't need to do anything

    sub byte [buff], 20h    ; Convert lowercase to uppercase

write:
    mov eax, 4          ; sys_write call
    mov ebx, 1          ; stdout
    mov ecx, buff       ; Pass address of char to write
    mov edx, 1          ; Pass number of chars to write
    int 80h
    jmp read

exit:
    mov eax, 1          ; sys_exit call
    mov ebx, 0          ; return value
    int 80h
