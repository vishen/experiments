
build_hello_go:
	go build -gcflags="-N -l" -o hello hello.go

build_hello_asm:
	nasm -f elf64 -o smaller_hello.o smaller_hello.asm
	ld -o smaller_hello smaller_hello.o

build_debugger:
	go build -o debugger debugger.go

inspect_binary:
	objdump -xS ./smaller_hello
