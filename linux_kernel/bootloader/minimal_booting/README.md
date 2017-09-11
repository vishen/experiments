# Minimal bootloader

Will boot and print "!" on the screen

## Running
```
$ nasm -f bin boot.nasm && qemu-system-x86_64 boot
$ gobjdump -D -b binary -mi386 -Maddr16,data16,intel boot  # On macosx
```
