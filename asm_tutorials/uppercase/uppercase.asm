SECTION .data
SECTION .bss
    BUFFLEN equ 1024            ; Length of buffer
    buff: resb BUFFLEN          ; Text buffer

SECTION .text
    global _start

_start:
    nop

; Read from file into a buffer
read:
    mov eax, 3                  ; sys_read
    mov ebx, 0                  ; stdin
    mov ecx, buff               ; Pass offset of the buffer to read to
    mov edx, BUFFLEN            ; How manay bytes to read
    int 80h
    mov esi, eax                ; Keep the return value from 'sys_read' (eax) -> number of bytes read
    cmp eax, 0                  ; Check if we have EOF
    je exit

; Set-up registers for buffer step
    mov ecx, esi                ; Place the number of bytes read into ecx
    mov ebp, buff               ; Place address of buffer into edp
    dec ebp                     ; Adjust count to offset

; Go through the buffer and convert lowercase to uppercase
scan_and_replace:
    cmp byte [ebp + ecx], 61h   ; Is it below 'a'
    jb continue
    cmp byte[ebp + ecx], 7ah    ; Is it above 'z'
    ja continue
    sub byte [ebp + ecx], 20h   ; At this point we have a lowercase letter and we need to uppercase it
continue:
    dec ecx                     ; Decrement counter by 1
    jnz scan_and_replace        ; Continue until counter is zero

; Write processed buffer to stdout
write:
    mov eax, 4                  ; sys_write
    mov ebx, 1                  ; stdout
    mov ecx, buff               ; Offset of buffer to write from
    mov edx, esi                ; We want to write the number of characters read from the file! This should be saved from the read: step
    int 80h
    jmp read

exit:
    mov eax, 1                  ; sys_exit
    mov ebx, 0                  ; Exit code
    int 80h
