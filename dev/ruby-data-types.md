# Ruby 2.4 Data Types

- Types
  - Numeric
    - Integer
      - Fixnum (deprecated in Ruby 2.4)
      - Bignum (deprecated in Ruby 2.4)
    - Float
    - Rational (the `r` suffix was added in Ruby 2.1)
    - Complex (the `i` suffix was added in Ruby 2.1)
  - String
  - NilClass
  - Array
  - Hash
  - Range
  - Symbol

## Class Hierarchy

- [BasicObject](https://ruby-doc.org/core/BasicObject.html)
  - [Object](https://ruby-doc.org/core/Object.html)
    - [TrueClass](https://ruby-doc.org/core/TrueClass.html)
    - [FalseClass](https://ruby-doc.org/core/FalseClass.html)
    - [Numeric](https://ruby-doc.org/core/Numeric.html)
      - [Integer](https://ruby-doc.org/core/Integer.html)
        - [Fixnum](https://ruby-doc.org/core-2.3.4/Fixnum.html) (deprecated in Ruby 2.4)
        - [Bignum](https://ruby-doc.org/core-2.3.4/Bignum.html) (deprecated in Ruby 2.4)
      - [Float](https://ruby-doc.org/core/Float.html)
      - [Rational](https://ruby-doc.org/core/Rational.html) (the `r` suffix was added in Ruby 2.1)
      - [Complex](https://ruby-doc.org/core/Complex.html) (the `i` suffix was added in Ruby 2.1)
    - [String](https://ruby-doc.org/core/String.html)
    - [Regexp](https://ruby-doc.org/core/Regexp.html)
    - [NilClass](https://ruby-doc.org/core/NilClass.html)
    - [Array](https://ruby-doc.org/core/Array.html)
    - [Hash](https://ruby-doc.org/core/Hash.html)
    - [Range](https://ruby-doc.org/core/Range.html)
    - [Symbol](https://ruby-doc.org/core/Symbol.html)

Note:
Since Ruby 2.1.0, Rational literals are available, like these: `0r`, `2r/3`, `2/3r`.

Note:
Since Ruby 2.1.0, Complex literals are available, like these: `0i`, `1+0i`, `1+1i`.

Note:
Since Ruby 2.4.0, Fixnum and Bignum are unified into Integer.

Note:
It seems that there does not exist an official Ruby Language Specification.
Therefore, the above information was gathered from various sources, including:

- [Ruby Programming - Wikibooks](https://en.wikibooks.org/wiki/Ruby_Programming)
- [Ruby Programming/Data types - Wikibooks](https://en.wikibooks.org/wiki/Ruby_Programming/Data_types)
- Running the following commands:

      git clone https://github.com/ruby/ruby
      cd ruby
      ruby sample/exyacc.rb < parse.y
