# Bootloader

The BIOS is expected to run the bootloader!

The bootloader is expected to be stored at a particular location, so that the BIOS or MBR can find it, load it, and run it. If the boot device is a hard disk (or something else emulating a hard disk) then the bootloader is expected to be stored as the very first "block" of the partition.

The main function of these bootloaders is to find the kernel, wherever it is on the media, load it, and run it. Additionally, the bootloaders need to set up a known environment for the kernel (which often includes switching to Protected Mode). The bootloaders also might collect some system data for the kernel to use (some data is much easier to get while the system is still in Real Mode).

Every OS is expected to have its own bootloader, stored on the media with the kernel.

There are a number of bootloaders / bootsectors that can boot Linux, such as GRUB 2 and syslinux. The Linux kernel has a Boot protocol [7] which specifies the requirements for a bootloader to implement Linux support. [6]

## BIOS / UEFI
Every computer with a motherboard includes a special chip referred to as the BIOS (Memory Basic Input/Output System) or UEFI. [1]

When an x86-based computer is turned on (assuming that it turns on successfully), it begins a complex path to get to the stage where control is transferred to your kernel's "main()" routine. The exact sequence of steps depends on what kind of device the computer decides to boot from, and whether it uses the legacy BIOS boot method, or the new UEFI method. The new UEFI method completely changes the entire boot process, and is covered in another article. The intention is that the UEFI boot process will someday completely replace the BIOS method in all new computers, but it may not succeed. [2]

In reality, both legacy motherboards and UEFI-based motherboards come with BIOS ROMs, which contain firmware that performs the initial power-on configuration of the system before loading some third-party code into memory and jumping to it. The differences between legacy BIOS firmware and UEFI BIOS firmware are where they find that code, how they prepare the system before jumping to it, and what convenience functions they provide for the code to call while running. [3]

### Functions of BIOS / UEFI
BIOS Functions:
```
- POST: Test computer hardware insuring hardware is properly functioning before starting process of loading Operating System.
- Bootstrap Loader: Process of locating the operating system. If capable Operating system located BIOS will pass the control to it
```

### Initialisation

On a legacy system, BIOS performs all the usual platform initialization (memory controller configuration, PCI bus configuration and BAR mapping, graphics card initialization, etc.), but then drops into a backwards-compatible real mode environment. [3]

UEFI firmware performs those same steps, but also prepares a protected mode environment with flat segmentation and for x86-64 CPUs, a long mode environment with identity-mapped paging. The A20 gate is enabled as well. [3]

If BIOS initialisation occurs, since it is then in real mode, the bootloader must then perform the following:
```
- Enable A20 gate
- Configure GDT and IDT (?)
- Switch to protected mode (?)
- For x86-64 CPUS; configure paging and switch to long mode (?)
```

The very first thing that happens when the machine is turned on, is that the CPU starts trying to run a program at the very end of the 4Gb memory area. At that location must be some ROM which contains a BIOS initialization program. The initialization code can be large, and the ROM may be as large as 256Kb in size. An OS programmer cannot modify or control this stage of the process in any way. [2]

RAM detection -> the BIOS must use some RAM in order to perform its functions. To use the RAM, the BIOS must first detect the type and quantity of RAM chips installed. This can only be done (using chipset-specific methods) while the CPU is running code that is stored in ROM, specifically. Once the RAM has been detected, the BIOS may perform a simple memory test on it, and then the BIOS loads data and code into several memory areas in RAM: the BDA, the EBDA, and the 64K "BIOS area" at physical address `0xF0000` to `0xFFFFF`. The BIOS also sets up a tiny stack somewhere in memory, and sets up the Real Mode IVT from physical address `0` to `0x3FF`. Some of the physical memory between address `0xA0000` and `0xFFFFF` is then set to "read only" mode using chipset-specific methods. [2]

Hardware detection/Initialization -> the BIOS detects, enumerates, configures, and initializes every bus, and almost every piece of hardware on the system, using values that the BIOS chooses. It stores a great deal of information about all of this hardware, for the OS to later parse. If the BIOS finds any ROM chips on any hardware, they are mapped (not loaded) into physical memory at addresses that the BIOS chooses. It is important to note that many BIOSes do a rather bad job of configuring, sometimes. An OS may well need to reconfigure the MTRRs, or the PCI bus, or the mapping of some PCI devices. The BIOS is supposed to always bring all the hardware into a functional state, but that state may not be optimal, or even technically "legal" according to the specs. [2]

"Boot sequence" -> at this point, the BIOS is done with its initialization. Now it tries to transfer control to the next stage of the bootloader process; so the BIOS must choose the "boot device". There is a list stored in CMOS, called the "boot sequence", that tells the BIOS which devices to test, and in what order, to see if they exist and are bootable. The BIOS may try to boot from a floppy disk, hard disk "C:", a USB flash memory device, a CD, a network, or something else. All of these devices can have some type of "bootsector", and there is a flag that the BIOS can check to see if the bootsector is valid. The BIOS will transfer control to the first valid bootsector that it finds, as it searches through the boot sequence. If the BIOS never finds a valid bootsector, it will lock up with an error message. [2]

The BIOS transfers 512 bytes of data from each device that exists, into physical memory starting at address `0x7c00`. If the last two bytes transferred are `0x55`, and then `0xAA`, then the BIOS considers this to be a valid bootsector, and starts running the code that now begins at `0x7c00`. [2]

## A20 Gate

The A20 Address Line is the physical representation of the 21st bit (number 20, counting from 0) of any memory access. When the IBM-AT (Intel 286) was introduced, it was able to access up to sixteen megabytes of memory (instead of the 1 MByte of the 8086). But to remain compatible with the 8086, a quirk in the 8086 architecture (memory wraparound) had to be duplicated in the AT. To achieve this, the A20 line on the address bus was disabled by default. [4]

The wraparound was caused by the fact the 8086 could only access 1 megabyte of memory, but because of the segmented memory model they could effectively address up to 1 megabyte and 64 kilobytes (minus 16 bytes). Because there are twenty address lines on the 8086 (A0 through A19), any address above the 1 megabyte mark wraps around to zero. For some reason a few short-sighted programmers decided to write programs that actually used this wraparound (rather than directly addressing the memory at its normal location at the bottom of memory). Therefore in order to support these 8086-era programs on the new processors, this wraparound had to be emulated on the IBM AT and its compatibles; this was originally achieved by way of a latch that by default set the A20 line to zero. Later the 486 added the logic into the processor and introduced the A20M pin to control it. [4]

For an operating system developer (or Bootloader developer) this means the A20 line has to be enabled so that all memory can be accessed. This started off as a simple hack but as simpler methods were added to do it, it became harder to program code that would definitely enable it and even harder to program code that would definitely disable it. [4]

Example: [5]

## Resources
```
- [1] http://flint.cs.yale.edu/feng/cos/resources/BIOS/
- [2] http://wiki.osdev.org/System_Initialization_(x86)
- [3] http://wiki.osdev.org/UEFI
- [4] http://wiki.osdev.org/A20_Line
- [5] https://www.quora.com/What-is-the-A20-gate-in-a-CPU
- [6] https://0xax.gitbooks.io/linux-insides/Booting/linux-bootstrap-1.html
- [7] https://github.com/torvalds/linux/blob/16f73eb02d7e1765ccab3d2018e0bd98eb93d973/Documentation/x86/boot.txt
```

# TODOS

Replace all the `(?)`!