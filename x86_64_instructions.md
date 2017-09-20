# x86_64 Instructions

## Instruction Format
    ┌─────────────┬────────┬────────┬─────┬──────────────┬───────────┐
    │ Instruction │ Opcode │ ModR/M │ SIB │ Displacement │ Immediate │
    │ Prefixes    │        │        │     │              │           │
    └─────────────┴────────┴────────┴─────┴──────────────┴───────────┘

- Instruction Prefixes:
  - Prefixes of 1 byte each (optional).
  - The `REX` prefix is optional, but if used must be immediately before the opcode.
- Opcode:
  - 1-, 2-, or 3-byte opcode.
- ModR/M:
  - 1 byte (if required).
  - Format:

        7    6 5      3 2    0
        ┌─────┬────────┬─────┐
        │ Mod │ Reg/   │ R/M │
        │     │ Opcode │     │
        └─────┴────────┴─────┘
- SIB (Scale-Index-Base):
  - 1 byte (if required).
  - Format:

        7      6 5     3 2     0
        ┌───────┬───────┬──────┐
        │ Scale │ Index │ Base │
        └───────┴───────┴──────┘
- Displacement:
  - Address displacement of 1, 2, or 4 bytes or none.
  - Some rare instructions can take an 8-byte immediate or 8-byte displacement.
- Immediate:
  - Immediate data of 1, 2, or 4 bytes or none.
  - Some rare instructions can take an 8-byte immediate or 8-byte displacement.

References:
- 2.1 INSTRUCTION FORMAT FOR PROTECTED MODE, REAL-ADDRESS MODE, AND VIRTUAL-8086 MODE
- 2.2 IA-32E MODE

## Registers
    ┌───────┬──────┬──────┬──────┬─────┬──────┬─────┬─────┬───────┬───────┬───────┬────────┬────────┐
    │ rex,  │ 8-   │ 16-  │ 32-  │ 64- │ 16-  │ 80- │ 64- │ 128-  │ 256-  │ 512-  │ Con-   │ Debug  │
    │ reg   │ bit  │ bit  │ bit  │ bit │ bit  │ bit │ bit │ bit   │ bit   │ bit   │ trol   │        │
    │       │ GP   │ GP   │ GP   │ GP  │ Sreg │ x87 │ MMX │ XMM   │ YMM   │ ZMM   │        │        │
    ├───────┼──────┼──────┼──────┼─────┼──────┼─────┼─────┼───────┼───────┼───────┼────────┼────────┤
    │ 0,000 │ AL   │ AX   │ EAX  │ RAX │ ES   │ ST0 │ MM0 │ XMM0  │ YMM0  │ ZMM0  │ CR0    │ DR0    │
    │ 0,001 │ CL   │ CX   │ ECX  │ RCX │ CS   │ ST1 │ MM1 │ XMM1  │ YMM1  │ ZMM1  │ CR1 *  │ DR1    │
    │ 0,010 │ DL   │ DX   │ EDX  │ RDX │ SS   │ ST2 │ MM2 │ XMM2  │ YMM2  │ ZMM2  │ CR2    │ DR2    │
    │ 0,011 │ BL   │ BX   │ EBX  │ RBX │ DS   │ ST3 │ MM3 │ XMM3  │ YMM3  │ ZMM3  │ CR3    │ DR3    │
    ├───────┼──────┼──────┼──────┼─────┼──────┼─────┼─────┼───────┼───────┼───────┼────────┼────────┤
    │ 0,100 │ AH   │ SP   │ ESP  │ RSP │ FS   │ ST4 │ MM4 │ XMM4  │ YMM4  │ ZMM4  │ CR4    │ DR4 *  │
    │ 0,101 │ CH   │ BP   │ EBP  │ RBP │ GS   │ ST5 │ MM5 │ XMM5  │ YMM5  │ ZMM5  │ CR5 *  │ DR5 *  │
    │ 0,110 │ DH   │ SI   │ ESI  │ RSI │ -    │ ST6 │ MM6 │ XMM6  │ YMM6  │ ZMM6  │ CR6 *  │ DR6    │
    │ 0,111 │ BH   │ DI   │ EDI  │ RDI │ -    │ ST7 │ MM7 │ XMM7  │ YMM7  │ ZMM7  │ CR7 *  │ DR7    │
    ├───────┼──────┼──────┼──────┼─────┼──────┼─────┼─────┼───────┼───────┼───────┼────────┼────────┤
    │ 1,000 │ R8L  │ R8W  │ R8D  │ R8  │ ES   │ -   │ MM0 │ XMM8  │ YMM8  │ ZMM8  │ CR8 *  │ DR8 *  │
    │ 1,001 │ R9L  │ R9W  │ R9D  │ R9  │ CS   │ -   │ MM1 │ XMM9  │ YMM9  │ ZMM9  │ CR9 *  │ DR9 *  │
    │ 1,010 │ R10L │ R10W │ R10D │ R10 │ SS   │ -   │ MM2 │ XMM10 │ YMM10 │ ZMM10 │ CR10 * │ DR10 * │
    │ 1,011 │ R11L │ R11W │ R11D │ R11 │ DS   │ -   │ MM3 │ XMM11 │ YMM11 │ ZMM11 │ CR11 * │ DR11 * │
    ├───────┼──────┼──────┼──────┼─────┼──────┼─────┼─────┼───────┼───────┼───────┼────────┼────────┤
    │ 1,100 │ R12L │ R12W │ R12D │ R12 │ FS   │ -   │ MM4 │ XMM12 │ YMM12 │ ZMM12 │ CR12 * │ DR12 * │
    │ 1,101 │ R13L │ R13W │ R13D │ R13 │ GS   │ -   │ MM5 │ XMM13 │ YMM13 │ ZMM13 │ CR13 * │ DR13 * │
    │ 1,110 │ R14L │ R14W │ R14D │ R14 │ -    │ -   │ MM6 │ XMM14 │ YMM14 │ ZMM14 │ CR14 * │ DR14 * │
    │ 1,111 │ R15L │ R15W │ R15D │ R15 │ -    │ -   │ MM7 │ XMM15 │ YMM15 │ ZMM15 │ CR15 * │ DR15 * │
    └───────┴──────┴──────┴──────┴─────┴──────┴─────┴─────┴───────┴───────┴───────┴────────┴────────┘

Note:
- In the assembly below, `R8L` ~ `R15L` are known as `R8B` ~ `R15B` respectively.

References:
- 2.2.2 Additional Encodings for Control and Debug Registers
- 3.1.1.3 Instruction Column in the Opcode Summary Table
- B.1.4.5 Segment Register (sreg) Field
- B.1.4.6 Special-Purpose Register (eee) Field
- http://wiki.osdev.org/X86-64_Instruction_Encoding#Registers
- https://www-user.tu-chemnitz.de/~heha/viewchm.php/hs/x86.chm/x64.htm

## REX Prefix

    7     4  3   2   1   0
    ┌──────┬───┬───┬───┬───┐
    │ 0100 │ W │ R │ X │ B │
    └──────┴───┴───┴───┴───┘

References:
- 2.2.1.2 More on REX Prefix Fields

## Examples

References:
- https://www.systutorials.com/72643/beginners-guide-x86-64-instruction-encoding/
