# Ruby

## Basics

Assignment: `answer = 42`

## Fixnum

### Arithmetic Operations

Name                | Method | Example | Result
------------------- | ------ | ------- | ------
Addition            | +      | 11 + 5  | 16
Subtraction         | -      | 11 - 5  | 6
Multiplication      | *      | 11 * 5  | 55
Euclidean Division  | /      | 11 / 5  | 2
Euclidean Modulo    | %      | 11 % 5  | 1
Exponentiation      | **     | 11 ** 5 | 161051

### Bit Operations

Name                   | Method | Example     | Result
---------------------- | ------ | ----------- | ------
Negate                 | ~      | ~11         | -12
Bitwise OR             | &#124; | 11 &#124; 5 | 15
Bitwise AND            | &      | 11 & 5      | 1
Bitwise XOR            | ^      | 11 ^ 5      | 14
Arithmetic Left Shift  | <<     | 11 << 3     | 88
Arithmetic Right Shift | >>     | 11 >> 3     | 1

### Euclidean Division and Modulo

Division | Result | Modulo    | Result | Expression
-------- | ------ | --------- | ------ | -----------------------
11 / 5   | 2      | 11 % 5    | 1      | 11 = 5 * 2 + 1
11 / -5  | -3     | 11 % -5   | -4     | 11 = (-5) * (-3) + (-4)
-11 / 5  | -3     | -11 % 5   | 4      | -11 = 5 * (-3) + 4
-11 / -5 | 2      | -11 % -5  | -1     | -11 = (-5) * 2 + (-1)

### Visualizing Euclidean Division and Modulo

|        |     |     |         |     |     |
| ------ | --- | --- | ------- | --- | --- |
| 5 * 3  | 15  |     | -5 * 3  | -15 |     |
|        | 11  |     |         | -11 |     |
| 5 * 2  | 10  | *   | -5 * 2  | -10 | *   |
| 5 * 1  | 5   |     | -5 * 1  | -5  |     |
| 5 * 0  | 0   |     | -5 * 0  | 0   |     |
| 5 * -1 | -5  |     | -5 * -1 | 5   |     |
| 5 * -2 | -10 |     | -5 * -2 | 10  |     |
|        | -11 |     |         | 11  |     |
| 5 * -3 | -15 | *   | -5 * -3 | 15  | *   |

Observation:
The one which is "equal to or just less than" the dividend is selected as the "product",
then the quotient and the remainder follow.

## Strings

Single quote (no escape sequence): `'Hello\tworld'` == `"Hello\\tworld"`

Double quote (escape sequence available): `"Hello\tworld"`

### Methods

Name          | Method  | Example           | Result
------------- | ------- | ----------------- | ------------
Length        | length  | "Kenji".length    | 5
Concatenation | +       | "Kenji" + 65.to_s | "Kenji65"
Append        | <<      | "Kenji" << 65     | "KenjiA"
Copy          | *       | "Kenji" * 2       | "KenjiKenji"
Format        | %       | "%05d" % 123      | "00123"
Reverse       | reverse | "Kenji".reverse   | "ijneK"
Lines | lines
Bytes | bytes
Characters | chars
Find String? | [] | s['another string']
Containment | include? | poem.include? "my hand"
Lowercase | downcase |
Remove Characters | delete |

## Arrays

Empty Array: `[]`

Non-empty Array: `[12, 47, 35]`

### Methods

Name                    | Method | Example             | Result
----------------------- | ------ | ------------------- | ------------
Length                  | length | [12, 47, 35].length | 3
Index                   | []     | [12, 47, 35][2]     | 35
Maximum                 | max    | [12, 47, 35].max    | 47
Sort (returns new copy) | sort   | [12, 47, 35].sort   | [12, 35, 47]
Sort (in place)         | sort!  | [12, 47, 35].sort!  | [12, 35, 47]
| join

## Hashes

Empty Hash: `{}`

Non-empty Hash: `{"Hello"=>1, "World"=>2}`

### Methods

Name    | Method | Example                           | Result
------- | ------ | --------------------------------- | ------
Length  | length | {"Hello"=>1, "World"=>2}.length   | 2
Index   | []     | {"Hello"=>1, "World"=>2}["Hello"] | 1
Key Set | keys   | {"Hello"=>1, "World"=>2}.keys     | ["Hello", "World"]

## Symbols

:hp

:mp
