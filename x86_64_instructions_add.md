# x86_64 Instructions (ADD)

## Instruction Encoding

| Opcode             | Instruction        | Op<br>/En | 64-bit<br>Mode | Compat<br>/Leg<br>Mode | Description |
| ------------------ | ------------------ | --------- | -------------- | ---------------------- | ----------- |
| `04 ib`            | `ADD AL, imm8`     | `I`  | Valid | Valid | Add `imm8` to `AL`.                              |
| `05 iw`            | `ADD AX, imm16`    | `I`  | Valid | Valid | Add `imm16` to `AX`.                             |
| `05 id`            | `ADD EAX, imm32`   | `I`  | Valid | Valid | Add `imm32` to `EAX`.                            |
| `REX.W + 05 id`    | `ADD RAX, imm32`   | `I`  | Valid | N.E.  | Add `imm32` sign-extended to 64-bits to `RAX`.   |
| `80 /0 ib`         | `ADD r/m8, imm8`   | `MI` | Valid | Valid | Add `imm8` to `r/m8`.                            |
| `REX + 80 /0 ib`   | `ADD r/m8, imm8`   | `MI` | Valid | N.E.  | Add sign-extended `imm8` to `r/m64`.             |
| `81 /0 iw`         | `ADD r/m16, imm16` | `MI` | Valid | Valid | Add `imm16` to `r/m16`.                          |
| `81 /0 id`         | `ADD r/m32, imm32` | `MI` | Valid | Valid | Add `imm32` to `r/m32`.                          |
| `REX.W + 81 /0 id` | `ADD r/m64, imm32` | `MI` | Valid | N.E.  | Add `imm32` sign-extended to 64-bits to `r/m64`. |
| `83 /0 ib`         | `ADD r/m16, imm8`  | `MI` | Valid | Valid | Add sign-extended `imm8` to `r/m16`.             |
| `83 /0 ib`         | `ADD r/m32, imm8`  | `MI` | Valid | Valid | Add sign-extended `imm8` to `r/m32`.             |
| `REX.W + 83 /0 ib` | `ADD r/m64, imm8`  | `MI` | Valid | N.E.  | Add sign-extended `imm8` to `r/m64`.             |
| `00 /r`            | `ADD r/m8, r8`     | `MR` | Valid | Valid | Add `r8` to `r/m8`.                              |
| `REX + 00 /r`      | `ADD r/m8, r8`     | `MR` | Valid | N.E.  | Add `r8` to `r/m8`.                              |
| `01 /r`            | `ADD r/m16, r16`   | `MR` | Valid | Valid | Add `r16` to `r/m16`.                            |
| `01 /r`            | `ADD r/m32, r32`   | `MR` | Valid | Valid | Add `r32` to `r/m32`.                            |
| `REX.W + 01 /r`    | `ADD r/m64, r64`   | `MR` | Valid | N.E.  | Add `r64` to `r/m64`.                            |
| `02 /r`            | `ADD r8, r/m8`     | `RM` | Valid | Valid | Add `r/m8` to `r8`.                              |
| `REX + 02 /r`      | `ADD r8, r/m8`     | `RM` | Valid | N.E.  | Add `r/m8` to `r8`.                              |
| `03 /r`            | `ADD r16, r/m16`   | `RM` | Valid | Valid | Add `r/m16` to `r16`.                            |
| `03 /r`            | `ADD r32, r/m32`   | `RM` | Valid | Valid | Add `r/m32` to `r32`.                            |
| `REX.W + 03 /r`    | `ADD r64, r/m64`   | `RM` | Valid | N.E.  | Add `r/m64` to `r64`.                            |

## Instruction Operand Encoding

| Op<br>/En | Operand 1             | Operand 2       |
| --------- | --------------------- | --------------- |
| `RM`      | `ModRM:reg` (r, w)    | `ModRM:r/m` (r) |
| `MR`      | `ModRM:r/m` (r, w)    | `ModRM:reg` (r) |
| `MI`      | `ModRM:r/m` (r, w)    | `imm8`          |
| `I`       | `AL`/`AX`/`EAX`/`RAX` | `imm8`          |

## Examples

### Assemble

#### `AL`/`AX`/`EAX`/`RAX` Shortcut (rN ← immN)

| Opcode          | Instruction      | Assembly              | Arch.  | Encoding            | Remarks |
| --------------- | ---------------- | --------------------- | ------ | ------------------- | ------- |
| `04 ib`         | `ADD AL, imm8`   | `ADD AL, 0x12`        | -      | `04 12`             |         |
| `05 iw`         | `ADD AX, imm16`  | `ADD AX, 0x0102`      | -      | `66 05 02 01`       | *       |
| `05 id`         | `ADD EAX, imm32` | `ADD EAX, 0x01020304` | -      | `05 04 03 02 01`    |         |
| `REX.W + 05 id` | `ADD RAX, imm32` | `ADD RAX, 0x01020304` | 64-bit | `48 05 04 03 02 01` |         |

Remarks:
- Is there no "r64 ← imm64" ???

#### Opcode `80 /0 ib` Instruction `ADD r/m8, imm8` (r8 ← imm8)

| Assembly       | Encoding   |
| -------------- | ---------- |
| `ADD CL, 0x12` | `80 c1 12` |
| `ADD DL, 0x12` | `80 c2 12` |
| `ADD BL, 0x12` | `80 c3 12` |
| `ADD AH, 0x12` | `80 c4 12` |
| `ADD CH, 0x12` | `80 c5 12` |
| `ADD DH, 0x12` | `80 c6 12` |
| `ADD BH, 0x12` | `80 c7 12` |

    ┌───────────┐ ┌────────────────────┐ ┌───────────┐
    │  Opcode   │ │       ModR/M       │ │ Immediate │
    ├───────────┤ ├─────┬────────┬─────┤ ├───────────┤
    │  Opcode   │ │ Mod │ Opcode │ R/M │ │ Immediate │
    ├───────────┤ ├─────┼────────┼─────┤ ├───────────┤
    │ 1000 0000 │ │ 11  │ 00 0   │ bbb │ │ 0001 0010 │
    ├───────────┤ ├─────┴────────┴─────┤ ├───────────┤
    │    8    0 │ │        c         ? │ │    1    2 │
    └───────────┘ └────────────────────┘ └───────────┘

#### Opcode `80 /0 ib` Instruction `ADD r/m8, imm8` ([r32] ← imm8, 32-bit)

| Assembly                   | Encoding      | Remarks |
| -------------------------- | ------------- | ------- |
| `ADD BYTE PTR [EAX], 0x12` | `80 00 12`    |         |
| `ADD BYTE PTR [ECX], 0x12` | `80 01 12`    |         |
| `ADD BYTE PTR [EDX], 0x12` | `80 02 12`    |         |
| `ADD BYTE PTR [EBX], 0x12` | `80 03 12`    |         |
| `ADD BYTE PTR [ESP], 0x12` | `80 04 24 12` | *       |
| `ADD BYTE PTR [EBP], 0x12` | `80 45 00 12` | *       |
| `ADD BYTE PTR [ESI], 0x12` | `80 06 12`    |         |
| `ADD BYTE PTR [EDI], 0x12` | `80 07 12`    |         |

#### Opcode `80 /0 ib` Instruction `ADD r/m8, imm8` ([r32] ← imm8, 64-bit)

| Assembly                   | Encoding         | Remarks |
| -------------------------- | ---------------- | ------- |
| `ADD BYTE PTR [EAX], 0x12` | `67 80 00 12`    |         |
| `ADD BYTE PTR [ECX], 0x12` | `67 80 01 12`    |         |
| `ADD BYTE PTR [EDX], 0x12` | `67 80 02 12`    |         |
| `ADD BYTE PTR [EBX], 0x12` | `67 80 03 12`    |         |
| `ADD BYTE PTR [ESP], 0x12` | `67 80 04 24 12` | *       |
| `ADD BYTE PTR [EBP], 0x12` | `67 80 45 00 12` | *       |
| `ADD BYTE PTR [ESI], 0x12` | `67 80 06 12`    |         |
| `ADD BYTE PTR [EDI], 0x12` | `67 80 07 12`    |         |

#### Opcode `80 /0 ib` Instruction `ADD r/m8, imm8` ([r64] ← imm8, 64-bit)

| Assembly                   | Encoding      | Remarks |
| -------------------------- | ------------- | ------- |
| `ADD BYTE PTR [RAX], 0x12` | `80 00 12`    |         |
| `ADD BYTE PTR [RCX], 0x12` | `80 01 12`    |         |
| `ADD BYTE PTR [RDX], 0x12` | `80 02 12`    |         |
| `ADD BYTE PTR [RBX], 0x12` | `80 03 12`    |         |
| `ADD BYTE PTR [RSP], 0x12` | `80 04 24 12` | *       |
| `ADD BYTE PTR [RBP], 0x12` | `80 45 00 12` | *       |
| `ADD BYTE PTR [RSI], 0x12` | `80 06 12`    |         |
| `ADD BYTE PTR [RDI], 0x12` | `80 07 12`    |         |

#### Opcode `REX + 80 /0 ib` Instruction `ADD r/m8, imm8` (r64 ← imm8, 64-bit)

| Assembly         | Encoding      |
| ---------------- | ------------- |

| `ADD SPL, 0x12`  | `40 80 c4 12` |
| `ADD BPL, 0x12`  | `40 80 c5 12` |
| `ADD SIL, 0x12`  | `40 80 c6 12` |
| `ADD DIL, 0x12`  | `40 80 c7 12` |
| `ADD R8B, 0x12`  | `41 80 c0 12` |
| `ADD R9B, 0x12`  | `41 80 c1 12` |
| `ADD R10B, 0x12` | `41 80 c2 12` |
| `ADD R11B, 0x12` | `41 80 c3 12` |
| `ADD R12B, 0x12` | `41 80 c4 12` |
| `ADD R13B, 0x12` | `41 80 c5 12` |
| `ADD R14B, 0x12` | `41 80 c6 12` |
| `ADD R15B, 0x12` | `41 80 c7 12` |

    ┌──────────────────────┐ ┌───────────┐ ┌────────────────────┐ ┌───────────┐
    │      REX Prefix      │ │  Opcode   │ │       ModR/M       │ │ Immediate │
    ├──────┬───┬───┬───┬───┤ ├───────────┤ ├─────┬────────┬─────┤ ├───────────┤
    │  -   │ W │ R │ X │ B │ │  Opcode   │ │ Mod │ Opcode │ R/M │ │ Immediate │
    ├──────┼───┼───┼───┼───┤ ├───────────┤ ├─────┼────────┼─────┤ ├───────────┤
    │ 0100 │ 0 │ 0 │ 0 │ b │ │ 1000 0000 │ │ 11  │ 00 0   │ bbb │ │ 0001 0010 │
    ├──────┴───┴───┴───┴───┤ ├───────────┤ ├─────┴────────┴─────┤ ├───────────┤
    │    4               ? │ │    8    0 │ │        c         ? │ │    1    2 │
    └──────────────────────┘ └───────────┘ └────────────────────┘ └───────────┘

References:
- Figure 2-5. Register-Register Addressing (No Memory Operand); REX.X Not Used

#### Opcode `REX + 80 /0 ib` Instruction `ADD r/m8, imm8` ([r64] ← imm8, 64-bit)

| Assembly                   | Encoding         | Remarks |
| -------------------------- | ---------------- | ------- |
| `ADD BYTE PTR [R8], 0x12`  | `41 80 00 12`    |         |
| `ADD BYTE PTR [R9], 0x12`  | `41 80 01 12`    |         |
| `ADD BYTE PTR [R10], 0x12` | `41 80 02 12`    |         |
| `ADD BYTE PTR [R11], 0x12` | `41 80 03 12`    |         |
| `ADD BYTE PTR [R12], 0x12` | `41 80 04 24 12` | *       |
| `ADD BYTE PTR [R13], 0x12` | `41 80 45 00 12` | *       |
| `ADD BYTE PTR [R14], 0x12` | `41 80 06 12`    |         |
| `ADD BYTE PTR [R15], 0x12` | `41 80 07 12`    |         |

#### Opcode `81 /0 iw` Instruction `ADD r/m16, imm16` (r16 ← imm16)

| Assembly         | Encoding         |
| ---------------- | ---------------- |
| `ADD CX, 0x0102` | `66 81 c1 02 01` |
| `ADD DX, 0x0102` | `66 81 c2 02 01` |
| `ADD BX, 0x0102` | `66 81 c3 02 01` |
| `ADD SP, 0x0102` | `66 81 c4 02 01` |
| `ADD BP, 0x0102` | `66 81 c5 02 01` |
| `ADD SI, 0x0102` | `66 81 c6 02 01` |
| `ADD DI, 0x0102` | `66 81 c7 02 01` |

#### Opcode `81 /0 id` Instruction `ADD r/m32, imm32` (r32 ← imm32)

| Assembly              | Encoding            |
| --------------------- | ------------------- |
| `ADD ECX, 0x01020304` | `81 c1 04 03 02 01` |
| `ADD EDX, 0x01020304` | `81 c2 04 03 02 01` |
| `ADD EBX, 0x01020304` | `81 c3 04 03 02 01` |
| `ADD ESP, 0x01020304` | `81 c4 04 03 02 01` |
| `ADD EBP, 0x01020304` | `81 c5 04 03 02 01` |
| `ADD ESI, 0x01020304` | `81 c6 04 03 02 01` |
| `ADD EDI, 0x01020304` | `81 c7 04 03 02 01` |

#### Opcode `REX.W + 81 /0 id` Instruction `ADD r/m64, imm32` (r64 ← imm32, 64-bit)

| Assembly              | Encoding               |
| --------------------- | ---------------------- |
| `ADD RCX, 0x01020304` | `48 81 c1 04 03 02 01` |
| `ADD RDX, 0x01020304` | `48 81 c2 04 03 02 01` |
| `ADD RBX, 0x01020304` | `48 81 c3 04 03 02 01` |
| `ADD RSP, 0x01020304` | `48 81 c4 04 03 02 01` |
| `ADD RBP, 0x01020304` | `48 81 c5 04 03 02 01` |
| `ADD RSI, 0x01020304` | `48 81 c6 04 03 02 01` |
| `ADD RDI, 0x01020304` | `48 81 c7 04 03 02 01` |

#### Opcode `83 /0 ib` Instruction `ADD r/m16, imm8` (r16 ← imm8)

| Assembly       | Encoding      |
| -------------- | ------------- |
| `ADD AX, 0x12` | `66 83 c0 12` |
| `ADD CX, 0x12` | `66 83 c1 12` |
| `ADD DX, 0x12` | `66 83 c2 12` |
| `ADD BX, 0x12` | `66 83 c3 12` |
| `ADD SP, 0x12` | `66 83 c4 12` |
| `ADD BP, 0x12` | `66 83 c5 12` |
| `ADD SI, 0x12` | `66 83 c6 12` |
| `ADD DI, 0x12` | `66 83 c7 12` |

#### Opcode `83 /0 ib` Instruction `ADD r/m32, imm8` (r32 ← imm8)

| Assembly        | Encoding   |
| --------------- | ---------- |
| `ADD EAX, 0x12` | `83 c0 12` |
| `ADD ECX, 0x12` | `83 c1 12` |
| `ADD EDX, 0x12` | `83 c2 12` |
| `ADD EBX, 0x12` | `83 c3 12` |
| `ADD ESP, 0x12` | `83 c4 12` |
| `ADD EBP, 0x12` | `83 c5 12` |
| `ADD ESI, 0x12` | `83 c6 12` |
| `ADD EDI, 0x12` | `83 c7 12` |

#### Opcode `REX.W + 83 /0 ib` Instruction `ADD r/m64, imm8` (r64 ← imm8, 64-bit)

| Assembly        | Encoding      |
| --------------- | ------------- |
| `ADD RAX, 0x12` | `48 83 c0 12` |
| `ADD RCX, 0x12` | `48 83 c1 12` |
| `ADD RDX, 0x12` | `48 83 c2 12` |
| `ADD RBX, 0x12` | `48 83 c3 12` |
| `ADD RSP, 0x12` | `48 83 c4 12` |
| `ADD RBP, 0x12` | `48 83 c5 12` |
| `ADD RSI, 0x12` | `48 83 c6 12` |
| `ADD RDI, 0x12` | `48 83 c7 12` |

#### Opcode `00 /r` Instruction `ADD r/m8, r8` (r8 ← r8)

| Assembly     | Encoding |
| ------------ | -------- |
| `ADD AL, AL` | `00 c0`  |
| `ADD AL, CL` | `00 c8`  |
| `ADD AL, DL` | `00 d0`  |
| `ADD AL, BL` | `00 d8`  |
| `ADD AL, AH` | `00 e0`  |
| `ADD AL, CH` | `00 e8`  |
| `ADD AL, DH` | `00 f0`  |
| `ADD AL, BH` | `00 f8`  |

#### Opcode `REX + 00 /r` Instruction `ADD r/m8, r8` (r8 ← r8, 64-bit)

| Assembly        | Encoding   |
| --------------- | ---------- |
| `ADD R8B, AL`   | `41 00 c0` |
| `ADD R8B, CL`   | `41 00 c8` |
| `ADD R8B, DL`   | `41 00 d0` |
| `ADD R8B, BL`   | `41 00 d8` |
| `ADD R8B, SPL`  | `41 00 e0` |
| `ADD R8B, BPL`  | `41 00 e8` |
| `ADD R8B, SIL`  | `41 00 f0` |
| `ADD R8B, DIL`  | `41 00 f8` |
| `ADD R8B, R8B`  | `45 00 c0` |
| `ADD R8B, R9B`  | `45 00 c8` |
| `ADD R8B, R10B` | `45 00 d0` |
| `ADD R8B, R11B` | `45 00 d8` |
| `ADD R8B, R12B` | `45 00 e0` |
| `ADD R8B, R13B` | `45 00 e8` |
| `ADD R8B, R14B` | `45 00 f0` |
| `ADD R8B, R15B` | `45 00 f8` |

#### Opcode `01 /r` Instruction `ADD r/m16, r16` (r16 ← r16)

| Assembly     | Encoding   |
| ------------ | ---------- |
| `ADD AX, AX` | `66 01 c0` |

#### Opcode `01 /r` Instruction `ADD r/m32, r32` (r32 ← r32)

| Assembly       | Encoding |
| -------------- | -------- |
| `ADD EAX, EAX` | `01 c0`  |

#### Opcode `REX.W + 01 /r` Instruction `ADD r/m64, r64` (r64 ← r64, 64-bit)

| Assembly       | Encoding   |
| -------------- | ---------- |
| `ADD RAX, RAX` | `48 01 c0` |

#### Opcode `02 /r` Instruction `ADD r8, r/m8` (r8 ← [r64], 64-bit)

| Assembly        | Encoding |
| --------------- | -------- |
| `ADD AL, [RAX]` | `02 00`  |

#### Opcode `REX + 02 /r` Instruction `ADD r8, r/m8` (r8 ← [r64], 64-bit)

| Assembly         | Encoding   |
| ---------------- | ---------- |
| `ADD R8B, [RAX]` | `44 02 00` |

#### Opcode `03 /r` Instruction `ADD r16, r/m16` (r16 ← [r64], 64-bit)

| Assembly        | Encoding   |
| --------------- | ---------- |
| `ADD AX, [RAX]` | `66 03 00` |

#### Opcode `03 /r` Instruction `ADD r32, r/m32` (r32 ← [r64], 64-bit)

| Assembly         | Encoding |
| ---------------- | -------- |
| `ADD EAX, [RAX]` | `03 00`  |

#### Opcode `REX.W + 03 /r` Instruction `ADD r64, r/m64` (r64 ← [r64], 64-bit)

| Assembly         | Encoding   |
| ---------------- | ---------- |
| `ADD RAX, [RAX]` | `48 03 00` |

0:  67 80 00 12             add    BYTE PTR [eax],0x12
4:  67 80 40 01 12          add    BYTE PTR [eax+0x1],0x12
9:  67 80 40 08 12          add    BYTE PTR [eax+0x8],0x12
e:  67 80 40 10 12          add    BYTE PTR [eax+0x10],0x12
13: 67 80 40 11 12          add    BYTE PTR [eax+0x11],0x12
18: 67 80 40 21 12          add    BYTE PTR [eax+0x21],0x12
1d: 67 80 40 31 12          add    BYTE PTR [eax+0x31],0x12
22: 67 80 40 41 12          add    BYTE PTR [eax+0x41],0x12
27: 67 80 40 51 12          add    BYTE PTR [eax+0x51],0x12
2c: 67 80 40 61 12          add    BYTE PTR [eax+0x61],0x12
31: 67 80 40 71 12          add    BYTE PTR [eax+0x71],0x12
36: 67 80 80 81 00 00 00    add    BYTE PTR [eax+0x81],0x12
3d: 12
3e: 67 80 80 01 01 00 00    add    BYTE PTR [eax+0x101],0x12
45: 12
46: 80 00 12                add    BYTE PTR [rax],0x12
49: 80 40 01 12             add    BYTE PTR [rax+0x1],0x12
4d: 80 40 08 12             add    BYTE PTR [rax+0x8],0x12
51: 80 40 10 12             add    BYTE PTR [rax+0x10],0x12
55: 80 40 11 12             add    BYTE PTR [rax+0x11],0x12
59: 80 40 21 12             add    BYTE PTR [rax+0x21],0x12
5d: 80 40 31 12             add    BYTE PTR [rax+0x31],0x12
61: 80 40 41 12             add    BYTE PTR [rax+0x41],0x12
65: 80 40 51 12             add    BYTE PTR [rax+0x51],0x12
69: 80 40 61 12             add    BYTE PTR [rax+0x61],0x12
6d: 80 40 71 12             add    BYTE PTR [rax+0x71],0x12
71: 80 80 81 00 00 00 12    add    BYTE PTR [rax+0x81],0x12
78: 80 80 01 01 00 00 12    add    BYTE PTR [rax+0x101],0x12

### Disassemble

| Encoding            | Architecture | Assembly                                   | Remarks |
| ------------------- | ------------ | ------------------------------------------ | ------- |
| `04 12`             | -            | `ADD AL, 0x12`                             |         |
| `66 05 02 01`       | -            | `ADD AX, 0x0102`                           |         |
| `05 04 03 02 01`    | -            | `ADD RAX, 0x01020304`                      |         |
| `48 05 04 03 02 01` | 64-bit       | `ADD RAX, 0x01020304`                      |         |

| Encoding            | Architecture | Assembly                                   | Remarks |
| ------------------- | ------------ | ------------------------------------------ | ------- |
| `80 c0 12`          | -            | `ADD AL, 0x12`                             | *       |
| `80 c1 12`          | -            | `ADD CL, 0x12`                             |         |
| `80 c2 12`          | -            | `ADD DL, 0x12`                             |         |
| `80 c3 12`          | -            | `ADD BL, 0x12`                             |         |
| `80 c4 12`          | -            | `ADD AH, 0x12`                             |         |
| `80 c5 12`          | -            | `ADD CH, 0x12`                             |         |
| `80 c6 12`          | -            | `ADD DH, 0x12`                             |         |
| `80 c7 12`          | -            | `ADD BH, 0x12`                             |         |

| Encoding            | Architecture | Assembly                                   | Remarks |
| ------------------- | ------------ | ------------------------------------------ | ------- |
| `80 00 12`          | 32-bit       | `ADD BYTE PTR [EAX], 0x12`                 |         |

| Encoding            | Architecture | Assembly                                   | Remarks |
| ------------------- | ------------ | ------------------------------------------ | ------- |
| `80 00 12`          | 64-bit       | `ADD BYTE PTR [RAX], 0x12`                 |         |
| `80 01 12`          | (same)       | `ADD BYTE PTR [RCX], 0x12`                 |         |
| `80 02 12`          | (same)       | `ADD BYTE PTR [RDX], 0x12`                 |         |
| `80 03 12`          | (same)       | `ADD BYTE PTR [RBX], 0x12`                 |         |
| `80 04 12`          | (same)       | `.BYTE 0x80; ADD AL, 0x12`                 | *       |
| `80 05 12`          | (same)       | `.BYTE 0x80; .BYTE 0x5; .BYTE 0x12`        | *       |
| `80 06 12`          | (same)       | `ADD BYTE PTR [RSI], 0x12`                 |         |
| `80 07 12`          | (same)       | `ADD BYTE PTR [RDI], 0x12`                 |         |

| Encoding            | Architecture | Assembly                                   | Remarks |
| ------------------- | ------------ | ------------------------------------------ | ------- |
| `41 80 c0 12`       | 64-bit       | `ADD R8B, 0x12`                            |         |
| `41 80 c1 12`       | (same)       | `ADD R9B, 0x12`                            |         |
| `41 80 c2 12`       | (same)       | `ADD R10B, 0x12`                           |         |
| `41 80 c3 12`       | (same)       | `ADD R11B, 0x12`                           |         |
| `41 80 c4 12`       | (same)       | `ADD R12B, 0x12`                           |         |
| `41 80 c5 12`       | (same)       | `ADD R13B, 0x12`                           |         |
| `41 80 c6 12`       | (same)       | `ADD R14B, 0x12`                           |         |
| `41 80 c7 12`       | (same)       | `ADD R15B, 0x12`                           |         |

| Encoding            | Architecture | Assembly                                   | Remarks |
| ------------------- | ------------ | ------------------------------------------ | ------- |
| `41 80 00 12`       | 64-bit       | `ADD BYTE PTR [R8], 0x12`                  |         |
| `41 80 01 12`       | (same)       | `ADD BYTE PTR [R9], 0x12`                  |         |
| `41 80 02 12`       | (same)       | `ADD BYTE PTR [R10], 0x12`                 |         |
| `41 80 03 12`       | (same)       | `ADD BYTE PTR [R11], 0x12`                 |         |
| `41 80 04 12`       | (same)       | `REX.B; .BYTE 0x80; ADD AL, 0x12`          | *       |
| `41 80 05 12`       | (same)       | `REX.B; .BYTE 0x80; .BYTE 0x5; .BYTE 0x12` | *       |
| `41 80 06 12`       | (same)       | `ADD BYTE PTR [R14], 0x12`                 |         |
| `41 80 07 12`       | (same)       | `ADD BYTE PTR [R15], 0x12`                 |         |
