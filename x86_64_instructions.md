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

References:
- 2.2.2 Additional Encodings for Control and Debug Registers
- 3.1.1.3 Instruction Column in the Opcode Summary Table
- B.1.4.5 Segment Register (sreg) Field
- B.1.4.6 Special-Purpose Register (eee) Field
- http://wiki.osdev.org/X86-64_Instruction_Encoding#Registers
- https://www-user.tu-chemnitz.de/~heha/viewchm.php/hs/x86.chm/x64.htm

## Examples:

### ADD AL, imm8

- Opcode: `04 ib`
- Assembly: `ADD AL, 0x12`
- Machine Code: `04 12`

### ADD AX, imm16

- Opcode: `05 iw`
- Assembly: `ADD AX, 0x0102`
- Machine Code: `66 05 02 01`

### ADD EAX, imm32

- Opcode: `05 id`
- Assembly: `ADD EAX, 0x01020304`
- Machine Code: `05 04 03 02 01`

### ADD RAX, imm32

- Opcode: `REX.W + 05 id`
- Assembly: `ADD RAX, 0x01020304`
- Machine Code: `48 05 04 03 02 01`

### ADD r/m8, imm8

#### Opcode: `80 /0 ib`, Assembly: `ADD BYTE PTR [??], 0x12`

| Machine Code  | `80 00 12` | `80 01 12` | `80 02 12` | `80 03 12` |
| ------------- | ---------- | ---------- | ---------- | ---------- |
| `??` (32-bit) | `EAX`      | `ECX`      | `EDX`      | `EBX`      |
| `??` (64-bit) | `RAX`      | `RCX`      | `RDX`      | `RBX`      |

| Machine Code  | `80 04 12` | `80 05 12` | `80 06 12` | `80 07 12` |
| ------------- | ---------- | ---------- | ---------- | ---------- |
| `??` (32-bit) | (N/A) *    | (N/A) *    | `ESI`      | `EDI`      |
| `??` (64-bit) | (N/A) *    | (N/A) *    | `RSI`      | `RDI`      |

*:
- `ADD BYTE PTR [ESP], 0x12` = `80 04 24 12` (32-bit)
- `ADD BYTE PTR [RSP], 0x12` = `80 04 24 12` (64-bit)
- `ADD BYTE PTR [ESP], 0x12` = `67 80 04 24 12` (64-bit)
- `ADD BYTE PTR [EBP], 0x12` = `80 45 00 12` (32-bit)
- `ADD BYTE PTR [RBP], 0x12` = `80 45 00 12` (64-bit)
- `ADD BYTE PTR [EBP + 0x0], 0x12` = `67 80 45 00 12` (64-bit)

#### Opcode: `80 /0 ib`, Assembly: `ADD ??, 0x12`

| Machine Code | `80 c0 12` | `80 c1 12` | `80 c2 12` | `80 c3 12` |
| ------------ | ---------- | ---------- | ---------- | ---------- |
| `??`         | `AL` #     | `CL`       | `DL`       | `BL`       |

| Machine Code | `80 c4 12` | `80 c5 12` | `80 c6 12` | `80 c7 12` |
| ------------ | ---------- | ---------- | ---------- | ---------- |
| `??`         | `AH`       | `CH`       | `DH`       | `BH`       |

#:
- This opcode is redundant. Use the opcode `04 ib` (`04 12`) to achieve the same result.

#### Opcode: `REX + 80 /0 ib`, Assembly: `REX ADD BYTE PTR [??], 0x12`

| Machine Code  | `40 80 00 12` | `40 80 01 12` | `40 80 02 12` | `40 80 03 12` |
| ------------- | ------------- | ------------- | ------------- | ------------- |
| `??`          | `RAX`         | `RCX`         | `RDX`         | `RBX`         |

| Machine Code  | `40 80 04 12` | `40 80 05 12` | `40 80 06 12` | `40 80 07 12` |
| ------------- | ------------- | ------------- | ------------- | ------------- |
| `??`          | (N/A) *       | (N/A) *       | `RSI`         | `RDI`         |

*:
- `REX ADD BYTE PTR [RSP], 0x12` = `40 80 04 24 12`
- `REX ADD BYTE PTR [RBP], 0x12` = `40 80 45 00 12`

#:
- All the above opcodes are redundant. Use the opcode `80 /0 ib` (`80 ?? 12`) to achieve the same result.

#### Opcode: `REX + 80 /0 ib`, Assembly: `ADD BYTE PTR [??], 0x12`

| Machine Code  | `41 80 00 12` | `41 80 01 12` | `41 80 02 12` | `41 80 03 12` |
| ------------- | ------------- | ------------- | ------------- | ------------- |
| `??`          | `R8`          | `R9`          | `R10`         | `R11`         |

| Machine Code  | `41 80 04 12` | `41 80 05 12` | `41 80 06 12` | `41 80 07 12` |
| ------------- | ------------- | ------------- | ------------- | ------------- |
| `??`          | (N/A) *       | (N/A) *       | `R14`         | `R15`         |

#### Opcode: `REX + 80 /0 ib`, Assembly: `ADD ??, 0x12`

| Machine Code  | `41 80 c0 12` | `41 80 c1 12` | `41 80 c2 12` | `41 80 c3 12` |
| ------------- | ------------- | ------------- | ------------- | ------------- |
| `??`          | `R8B`         | `R9B`         | `R10B`        | `R11B`        |

| Machine Code  | `41 80 c4 12` | `41 80 c5 12` | `41 80 c6 12` | `41 80 c7 12` |
| ------------- | ------------- | ------------- | ------------- | ------------- |
| `??`          | `R12B`        | `R13B`        | `R14B`        | `R15B`        |

References:
- https://www.systutorials.com/72643/beginners-guide-x86-64-instruction-encoding/
