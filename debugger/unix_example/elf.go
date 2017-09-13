package main

import (
	"debug/elf"
	"debug/gosym"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Printf("Elf debug %v\n", os.Args[1:])

	exe, err := elf.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Error: elf.Open(1): %s\n", err)
	}

	fmt.Printf("Exe: %#v\n", exe)

	for _, section := range exe.Sections {
		fmt.Println(section)
		/*data, err := section.Data()
		if err != nil {
			fmt.Printf("Error: Section.Data(1): %s\n", err)
		}
		fmt.Println(string(data))
		fmt.Println()*/
	}

	for _, prog := range exe.Progs {
		fmt.Println(prog)
	}

	symbols, err := exe.Symbols()
	if err != nil {
		fmt.Printf("Error: exe.Symbols(1): %s\n", err)
	} else {
		fmt.Printf("Symbols: %#v\n", symbols)
	}

	fmt.Printf("ELF Machine: %s\n", exe.Machine.String())
	fmt.Printf("ELF Arch: %s\n", exe.Class.String())
	fmt.Printf("ElF Data: %s\n", exe.Data)
	return

	var pclndat []byte
	if sec := exe.Section(".gopclntab"); sec != nil {
		pclndat, err = sec.Data()
		if err != nil {
			log.Fatalf("Cannot read .gopclntab section: %v\n", err)
		}
	}

	var symTabRaw []byte
	if sec := exe.Section(".gosymtab"); sec != nil {
		symTabRaw, err = sec.Data()
		if err != nil {
			log.Fatalf("Cannot read .gosymtab section: %v\n", err)
		}
	}

	pcln := gosym.NewLineTable(pclndat, exe.Section(".text").Addr)
	symTab, err := gosym.NewTable(symTabRaw, pcln)
	if err != nil {
		log.Fatal("Cannot create symbol table: %v\n", err)
	}

	sym := symTab.LookupFunc("main.main")
	filename, lineno, _ := symTab.PCToLine(sym.Entry)

	log.Printf("filename: %v\n", filename)
	log.Printf("lineno: %v\n", lineno)

}
