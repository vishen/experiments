all: build run

NAME=template

build:
	nasm -f elf -g -F stabs ${NAME}.asm
	ld -o ${NAME} ${NAME}.o

run:
	./${NAME}
