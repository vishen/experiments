all: build run

NAME=uppercase

build:
	nasm -f elf -g -F stabs ${NAME}.asm
	ld -o ${NAME} ${NAME}.o

run:
	./${NAME} < t1.txt > output.txt
