all: build run

build:
	nasm -f elf -g -F stabs eatsyscall.asm
	ld -o eatsyscall eatsyscall.o

run:
	./eatsyscall