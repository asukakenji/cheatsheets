# Swift 4 Data Types

- Types
  - Named Types
    - Classes
    - Structures
    - Enumerations
    - Protocols
  - Compound Types
    - Function Types
    - Tuple Types

## Structures

- Examples:
  - Structures:
    - [`Bool`](https://developer.apple.com/documentation/swift/bool)
    - [`Int`](https://developer.apple.com/documentation/swift/int)
    - [`Double`](https://developer.apple.com/documentation/swift/double)
    - [`Float`](https://developer.apple.com/documentation/swift/float)
    - [`String`](https://developer.apple.com/documentation/swift/string)
    - [`Character`](https://developer.apple.com/documentation/swift/character)
  - Generic Structures:
    - [`Range`](https://developer.apple.com/documentation/swift/range)
    - [`ClosedRange`](https://developer.apple.com/documentation/swift/closedrange)
    - [`Array`](https://developer.apple.com/documentation/swift/array)
    - [`Dictionary`](https://developer.apple.com/documentation/swift/dictionary)
    - [`Set`](https://developer.apple.com/documentation/swift/set)

## Enumerations

- Examples:
  - Enumerations:
    - [`Unicode`](https://developer.apple.com/documentation/swift/unicode)
  - Generic Enumerations:
    - [`Optional`](https://developer.apple.com/documentation/swift/optional)

## Protocols

- Examples:
  - Protocols:
    - [`Error`](https://developer.apple.com/documentation/swift/error)
    - [`OptionSet`](https://developer.apple.com/documentation/swift/optionset)

## Function Types

- Examples:
  - `() -> Void` (is an alias of `() -> ()`)
  - `(Void) -> Void` (is an alias of `(()) -> ()`, different from above)
  - `(Int) -> Int`
  - `(Int) -> (Int) -> Int` (the same as `(Int) -> ((Int) -> Int)`)

## Tuple Types

- Examples:
  - `(Int, Int)`
  - `(Int, String)`
  - `(Int, Int, Int)`
