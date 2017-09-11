all: build run

NAME=mul

build:
	nasm -f elf -g -F stabs ${NAME}.asm
	ld -o ${NAME} ${NAME}.o

run:
	./${NAME}
