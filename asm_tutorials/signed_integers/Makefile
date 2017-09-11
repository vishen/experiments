all: build run

NAME=signed_integers

build:
	nasm -f elf -g -F stabs ${NAME}.asm
	ld -o ${NAME} ${NAME}.o

run:
	./${NAME}
