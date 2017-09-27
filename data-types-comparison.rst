=====================
Data Types Comparison
=====================

+-------------------+------------------+-------------+-------------+----------------+-------------+
|                   | C                | Java        | C#          | Go             | Swift       |
+===================+==================+=============+=============+================+=============+
| Boolean           | ``bool``         | ``boolean`` | ``bool``    | ``bool``       | ``Bool``    |
|                   |                  |             |             |                |             |
|                   |                  |             | ``Boolean`` | ``Bool``       |             |
+-------------------+------------------+-------------+-------------+----------------+-------------+
| Signed            | ``int8_t``       | ``byte``    | ``sbyte``   | ``int8``       | ``Int8``    |
| 8-bit integer     |                  |             |             |                |             |
|                   |                  |             | ``SByte``   | ``Int8``       |             |
+-------------------+------------------+-------------+-------------+----------------+-------------+
| Unsigned          | ``uint8_t``      | ``boolean`` | ``byte``    | ``uint8``      | ``UInt8``   |
| 8-bit integer     |                  |             |             |                |             |
|                   |                  |             | ``Byte``    | ``Uint8``      |             |
|                   |                  |             |             |                |             |
|                   |                  |             |             | ``byte``       |             |
+-------------------+------------------+-------------+-------------+----------------+-------------+
| Signed            | ``int16_t``      | ``short``   | ``short``   | ``int16``      | ``Int16``   |
| 16-bit integer    |                  |             |             |                |             |
|                   |                  |             | ``Int16``   | ``Int16``      |             |
+-------------------+------------------+-------------+-------------+----------------+-------------+
| Unsigned          | ``uint16_t``     | ``char``    | ``ushort``  | ``uint16``     | ``UInt16``  |
| 16-bit integer    |                  |             |             |                |             |
|                   |                  |             | ``UInt16``  | ``Uint16``     |             |
|                   |                  |             |             |                |             |
|                   |                  |             | ``char``    |                |             |
|                   |                  |             |             |                |             |
|                   |                  |             | ``Char``    |                |             |
+-------------------+------------------+-------------+-------------+----------------+-------------+
| Signed            | ``int32_t``      | ``short``   | ``int``     | ``int32``      | ``Int32``   |
| 32-bit integer    |                  |             |             |                |             |
|                   |                  |             | ``Int32``   | ``Int32``      |             |
|                   |                  |             |             |                |             |
|                   |                  |             |             | ``rune``       |             |
+-------------------+------------------+-------------+-------------+----------------+-------------+
| Unsigned          | ``uint32_t``     |             | ``uint``    | ``uint32``     | ``UInt32``  |
| 32-bit integer    |                  |             |             |                |             |
|                   |                  |             | ``UInt32``  | ``Uint32``     |             |
+-------------------+------------------+-------------+-------------+----------------+-------------+
| Signed            | ``int64_t``      | ``long``    | ``long``    | ``int64``      | ``Int64``   |
| 64-bit integer    |                  |             |             |                |             |
|                   |                  |             | ``Int64``   | ``Int64``      |             |
+-------------------+------------------+-------------+-------------+----------------+-------------+
| Unsigned          | ``uint64_t``     |             | ``ulong``   | ``uint64``     | ``UInt64``  |
| 64-bit integer    |                  |             |             |                |             |
|                   |                  |             | ``UInt64``  | ``Uint64``     |             |
+-------------------+------------------+-------------+-------------+----------------+-------------+
| Signed machine    | ``int``          |             |             | ``int``        | ``Int``     |
| dependant integer |                  |             |             |                |             |
|                   |                  |             |             | ``Int``        |             |
+-------------------+------------------+-------------+-------------+----------------+-------------+
| Unsigned machine  | ``unsigned int`` |             |             | ``uint``       | ``UInt``    |
| dependant integer |                  |             |             |                |             |
|                   |                  |             |             | ``Uint``       |             |
+-------------------+------------------+-------------+-------------+----------------+-------------+
| 32-bit            |                  | ``float``   | ``float``   | ``float32``    | ``Float``   |
| floating-point    |                  |             |             |                |             |
|                   |                  |             | ``Single``  | ``Float32``    | ``Float32`` |
+-------------------+------------------+-------------+-------------+----------------+-------------+
| 64-bit            |                  | ``double``  | ``double``  | ``float64``    | ``Double``  |
| floating-point    |                  |             |             |                |             |
|                   |                  |             | ``Double``  | ``Float64``    | ``Float64`` |
+-------------------+------------------+-------------+-------------+----------------+-------------+
| 80-bit            |                  |             |             |                | ``Float80`` |
| floating-point    |                  |             |             |                |             |
+-------------------+------------------+-------------+-------------+----------------+-------------+
| 128-bit           |                  |             | ``decimal`` |                |             |
| demical           |                  |             |             |                |             |
|                   |                  |             | ``Decimal`` |                |             |
+-------------------+------------------+-------------+-------------+----------------+-------------+
| 64-bit complex    |                  |             |             | ``complex64``  |             |
| (32 + 32)         |                  |             |             |                |             |
|                   |                  |             |             | ``Complex64``  |             |
+-------------------+------------------+-------------+-------------+----------------+-------------+
| 128-bit complex   |                  |             |             | ``complex128`` |             |
| (64 + 64)         |                  |             |             |                |             |
|                   |                  |             |             | ``Complex128`` |             |
+-------------------+------------------+-------------+-------------+----------------+-------------+

- C:

  - The ``intN_t`` and ``uintN_t`` types are ``typedef``\s from ``<stdint.h>``.

  - Reference: http://www.cplusplus.com/reference/cstdint/

  - Reference: http://en.cppreference.com/w/cpp/types/integer

- Java:

  - The ``boolean`` type is a boolean type. It is implemented using unsigned 8-bit integer.

  - Reference: http://docs.oracle.com/javase/8/docs/technotes/guides/jni/spec/types.html

- C#:

  - The built-in C# types are aliases of predefined types in the ``System``
    namespace in the .Net Framework. For example, ``Int`` is an alias of
    ``System.Int``.

  - The ``char`` type is dedicated for representing Unicode characters.
    It ranges from ``U+0000`` to ``U+FFFF``.

  - Reference: https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/keywords/value-types

  - Reference: https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/keywords/reference-tables-for-types

- Go:

  - ``byte`` is an alias for ``uint8``. ``rune`` is an alias for ``int32``.

  - The types starting with a capital letter are constants from the ``reflect`` package.

  - Reference: https://golang.org/ref/spec#Numeric_types

  - Reference: https://golang.org/pkg/reflect/#Kind
