SECTION .data
msg: db "Hello, World Small",10
msgLen: equ $-msg

SECTION .bss

SECTION .text
	global _start

_start:
	mov rax, 1
	mov rdi, 1
	mov rsi, msg
	mov rdx, msgLen
	syscall

	mov rax, 60
	mov rdi, 0
	syscall
