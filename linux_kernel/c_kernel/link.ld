/*
*  link.ld
*/
OUTPUT_FORMAT(elf32-i386)
ENTRY(start)
SECTIONS
 {
   . = 0x100000;
   .text : { *(.text) } /* Get all the .text input sections from all input files */
   .data : { *(.data) }
   .bss  : { *(.bss)  }
 }
