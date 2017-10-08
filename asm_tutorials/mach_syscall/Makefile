build:
	nasm -f macho64 syscall.asm
	ld syscall.o

run: build
	./a.out

objdump:
	gobjdump -Sx ./a.out
