SECTION .data           ; Section containing initialisation data

EatMsg: db "Hello, Jonathan...",10,10,10  ; 10 here means new line I believe
EatLen: equ $-EatMsg

SECTION .bss            ; Section containing uninitilised data

SECTION .text           ; Section containing code

global _start           ; Linker needs to find the entry point

_start:
    nop                 ; Keeps gdb happily?

    mov eax, 4          ; Specify sys_write system call
    mov ebx, 1          ; Specify file descriptor 1: stdout
    mov ecx, EatMsg     ; Pass offset of the message
    mov edx, EatLen     ; Pass the length of the message
    int 80H             ; Make syscall to ouput text to stdout

    mov eax, 1          ; Specify Exit system call
    mov ebx, 0          ; Return an exit code of 0
    int 80H             ; Make syscall to terminate program
