# Go (Golang) Naming Conventions Cheatsheet

## Variables

#### Example 1: [`$GOROOT/src/os/error.go`](https://golang.org/src/os/error.go)

```go
// Portable analogs of some common system call errors.
var (
	ErrInvalid    = errors.New("invalid argument") // methods on File will return this error when the receiver is nil
	ErrPermission = errors.New("permission denied")
	ErrExist      = errors.New("file already exists")
	ErrNotExist   = errors.New("file does not exist")
)
```

#### Example 2: [`$GOROOT/src/os/file.go`](https://golang.org/src/os/file.go)

```go
// Stdin, Stdout, and Stderr are open Files pointing to the standard input,
// standard output, and standard error file descriptors.
var (
	Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
	Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
)
```

## Constants

#### Example 1: [`$GOROOT/src/os/file.go`](https://golang.org/src/os/file.go)

```go
// Flags to OpenFile wrapping those of the underlying system. Not all
// flags may be implemented on a given system.
const (
	O_RDONLY int = syscall.O_RDONLY // open the file read-only.
	O_WRONLY int = syscall.O_WRONLY // open the file write-only.
	O_RDWR   int = syscall.O_RDWR   // open the file read-write.
	O_APPEND int = syscall.O_APPEND // append data to the file when writing.
	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist
	O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
	O_TRUNC  int = syscall.O_TRUNC  // if possible, truncate file when opened.
)

// Seek whence values.
const (
	SEEK_SET int = 0 // seek relative to the origin of the file
	SEEK_CUR int = 1 // seek relative to the current offset
	SEEK_END int = 2 // seek relative to the end
)
```

#### Example 2: [`$GOROOT/src/os/types.go`](https://golang.org/src/os/types.go)

```go
// The defined file mode bits are the most significant bits of the FileMode.
// The nine least-significant bits are the standard Unix rwxrwxrwx permissions.
// The values of these bits should be considered part of the public API and
// may be used in wire protocols or disk representations: they must not be
// changed, although new bits might be added.
const (
	// The single letters are the abbreviations
	// used by the String method's formatting.
	ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory
	ModeAppend                                     // a: append-only
	ModeExclusive                                  // l: exclusive use
	ModeTemporary                                  // T: temporary file (not backed up)
	ModeSymlink                                    // L: symbolic link
	ModeDevice                                     // D: device file
	ModeNamedPipe                                  // p: named pipe (FIFO)
	ModeSocket                                     // S: Unix domain socket
	ModeSetuid                                     // u: setuid
	ModeSetgid                                     // g: setgid
	ModeCharDevice                                 // c: Unix character device, when ModeDevice is set
	ModeSticky                                     // t: sticky

	// Mask for the type bits. For regular files, none will be set.
	ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice

	ModePerm FileMode = 0777 // Unix permission bits
)
```

#### Example 3: [`$GOROOT/src/math/const.go`](https://golang.org/src/math/const.go)

```go
// Mathematical constants.
const (
	E   = 2.71828182845904523536028747135266249775724709369995957496696763 // http://oeis.org/A001113
	Pi  = 3.14159265358979323846264338327950288419716939937510582097494459 // http://oeis.org/A000796
	Phi = 1.61803398874989484820458683436563811772030917980576286213544862 // http://oeis.org/A001622

	Sqrt2   = 1.41421356237309504880168872420969807856967187537694807317667974 // http://oeis.org/A002193
	SqrtE   = 1.64872127070012814684865078781416357165377610071014801157507931 // http://oeis.org/A019774
	SqrtPi  = 1.77245385090551602729816748334114518279754945612238712821380779 // http://oeis.org/A002161
	SqrtPhi = 1.27201964951406896425242246173749149171560804184009624861664038 // http://oeis.org/A139339

	Ln2    = 0.693147180559945309417232121458176568075500134360255254120680009 // http://oeis.org/A002162
	Log2E  = 1 / Ln2
	Ln10   = 2.30258509299404568401799145468436420760110148862877297603332790 // http://oeis.org/A002392
	Log10E = 1 / Ln10
)

// Floating-point limit values.
// Max is the largest finite value representable by the type.
// SmallestNonzero is the smallest positive, non-zero value representable by the type.
const (
	MaxFloat32             = 3.40282346638528859811704183484516925440e+38  // 2**127 * (2**24 - 1) / 2**23
	SmallestNonzeroFloat32 = 1.401298464324817070923729583289916131280e-45 // 1 / 2**(127 - 1 + 23)

	MaxFloat64             = 1.797693134862315708145274237317043567981e+308 // 2**1023 * (2**53 - 1) / 2**52
	SmallestNonzeroFloat64 = 4.940656458412465441765687928682213723651e-324 // 1 / 2**(1023 - 1 + 52)
)

// Integer limit values.
const (
	MaxInt8   = 1<<7 - 1
	MinInt8   = -1 << 7
	MaxInt16  = 1<<15 - 1
	MinInt16  = -1 << 15
	MaxInt32  = 1<<31 - 1
	MinInt32  = -1 << 31
	MaxInt64  = 1<<63 - 1
	MinInt64  = -1 << 63
	MaxUint8  = 1<<8 - 1
	MaxUint16 = 1<<16 - 1
	MaxUint32 = 1<<32 - 1
	MaxUint64 = 1<<64 - 1
)
```

## Struct Types

TODO: Write this!

## Interface Types

TODO: Write this!

## Main Package

Filename:
- `$src/cmd/$executable/*.go`

Filename Examples:
- [`$GOROOT/src/cmd/go/*.go`](https://golang.org/src/cmd/go/)

Code:
```go
package main
```

### Main Function

Filename:
- `main.go`; OR
- `$executable.go`

Filename Examples:
- [`$GOROOT/src/cmd/go/main.go`](https://golang.org/src/cmd/go/main.go)
- [`$GOROOT/src/cmd/gofmt/gofmt.go`](https://golang.org/src/cmd/gofmt/gofmt.go)
- [`$GOROOT/src/cmd/nm/nm.go`](https://golang.org/src/cmd/nm/nm.go)
- [`$GOROOT/src/cmd/objdump/main.go`](https://golang.org/src/cmd/objdump/main.go)

Code:
```go
func main() {
	/* ... */
}
```

## Testing

### White Box Testing

Filename:
- `*_test.go`

Filename Examples:
- [`$GOROOT/src/math/big/arith_test.go`](https://golang.org/src/math/big/arith_test.go)
- [`$GOROOT/src/math/big/decimal_test.go`](https://golang.org/src/math/big/decimal_test.go)
- [`$GOROOT/src/math/big/floatconv_test.go`](https://golang.org/src/math/big/floatconv_test.go)
- [`$GOROOT/src/math/big/hilbert_test.go`](https://golang.org/src/math/big/hilbert_test.go)
- [`$GOROOT/src/math/big/int_test.go`](https://golang.org/src/math/big/int_test.go)
- [`$GOROOT/src/math/big/nat_test.go`](https://golang.org/src/math/big/nat_test.go)
- [`$GOROOT/src/math/big/natconv_test.go`](https://golang.org/src/math/big/natconv_test.go)
- [`$GOROOT/src/math/cmplx/cmath_test.go`](https://golang.org/src/math/cmplx/cmath_test.go)
- [`$GOROOT/src/math/rand/rand_test.go`](https://golang.org/src/math/rand/rand_test.go)

Note:
- The test files are located in the same directory as the source files being tested;
- The package declaration in the test files is the same as that in the source files being tested;
- The test code is in the *SAME* package as the code under test.

References:
- [Command go > Description of testing functions][cmdgo-testing]

[cmdgo-testing]: https://golang.org/cmd/go/#hdr-Description_of_testing_functions

### Black Box Testing

Filename:
- `*_test.go`

Filename Examples:
- [`$GOROOT/src/math/all_test.go`](https://golang.org/src/math/all_test.go)
- [`$GOROOT/src/math/rand/regress_test.go`](https://golang.org/src/math/rand/regress_test.go)

Note:
- The test files are located in the same directory as the source files being tested;
- If the package declaration of the source files being tested is `package xxx`, then that of the test files is `package xxx_test`
- The test code is in a *DIFFERENT* package from the code under test.

References:
- [Command go > Test packages][cmdgo-test-packages]

[cmdgo-test-packages]: https://golang.org/cmd/go/#hdr-Test_packages

Code:
```go
package xxx_test
```

### Gray Box Testing (Backdoors)

Filename (Not enforced by the toolchain):
- `export_test.go`

Filename Examples:
- [`$GOROOT/src/flag/export_test.go`](https://golang.org/src/flag/export_test.go)
- [`$GOROOT/src/math/export_test.go`](https://golang.org/src/math/export_test.go)
- [`$GOROOT/src/net/http/export_test.go`](https://golang.org/src/net/http/export_test.go)

Note:
- The backdoor file exposes internals that are needed by the test package;
- The package declaration in the backdoor file is the same as that in the source files being tested;
- The backdoor file does not contain any tests.

### Test Functions

```go
import "testing"

func TestXxx(t *testing.T) {
	/* ... */
}
```

Filename:
- `*_test.go`

Filename Examples:
- [`$GOROOT/src/math/big/bits_test.go`](https://golang.org/src/math/big/bits_test.go)
- [`$GOROOT/src/math/big/calibrate_test.go`](https://golang.org/src/math/big/calibrate_test.go)
- [`$GOROOT/src/math/big/float_test.go`](https://golang.org/src/math/big/float_test.go)
- [`$GOROOT/src/math/big/floatmarsh_test.go`](https://golang.org/src/math/big/floatmarsh_test.go)
- [`$GOROOT/src/math/big/intconv_test.go`](https://golang.org/src/math/big/intconv_test.go)
- [`$GOROOT/src/math/big/intmarsh_test.go`](https://golang.org/src/math/big/intmarsh_test.go)
- [`$GOROOT/src/math/big/rat_test.go`](https://golang.org/src/math/big/rat_test.go)
- [`$GOROOT/src/math/big/ratconv_test.go`](https://golang.org/src/math/big/ratconv_test.go)
- [`$GOROOT/src/math/big/ratmarsh_test.go`](https://golang.org/src/math/big/ratmarsh_test.go)
- [`$GOROOT/src/math/rand/race_test.go`](https://golang.org/src/math/rand/race_test.go)
- [`$GOROOT/src/time/format_test.go`](https://golang.org/src/time/format_test.go)
- [`$GOROOT/src/time/zoneinfo_test.go`](https://golang.org/src/time/zoneinfo_test.go)
- [`$GOROOT/src/time/zoneinfo_windows_test.go`](https://golang.org/src/time/zoneinfo_windows_test.go)
- [`$GOROOT/src/unicode/digit_test.go`](https://golang.org/src/unicode/digit_test.go)
- [`$GOROOT/src/unicode/graphic_test.go`](https://golang.org/src/unicode/graphic_test.go)
- [`$GOROOT/src/unicode/letter_test.go`](https://golang.org/src/unicode/letter_test.go)
- [`$GOROOT/src/unicode/script_test.go`](https://golang.org/src/unicode/script_test.go)

### Benchmark Functions

```go
import "testing"

func BenchmarkXxx(b *testing.B) {
	/* ... */
}
```

Filename (Not enforced by the toolchain):
- `bench_test.go`; OR
- `performance_test.go`; OR
- `benchmark_test.go`

Preferred Package:
- `xxx` (`xxx` is the package of the source files)

Filename Examples:
- [`$GOROOT/src/crypto/cipher/benchmark_test.go`](https://golang.org/src/crypto/cipher/benchmark_test.go)
- [`$GOROOT/src/encoding/json/bench_test.go`](https://golang.org/src/encoding/json/bench_test.go)
- [`$GOROOT/src/go/parser/performance_test.go`](https://golang.org/src/go/parser/performance_test.go)
- [`$GOROOT/src/go/printer/performance_test.go`](https://golang.org/src/go/printer/performance_test.go)
- [`$GOROOT/src/image/draw/bench_test.go`](https://golang.org/src/image/draw/bench_test.go)
- [`$GOROOT/src/math/big/gcd_test.go`](https://golang.org/src/math/big/gcd_test.go)

### Example Functions

```go
func ExamplePrintln() {
	Println("The output of\nthis example.")
	// Output: The output of
	// this example.
}
```

Filename (Not enforced by the toolchain):
- `example_test.go`

Preferred Package:
- `xxx_test` (`xxx` is the package of the source files)

Filename Examples:
- [`$GOROOT/src/math/big/example_rat_test.go`](https://golang.org/src/math/big/example_rat_test.go)
- [`$GOROOT/src/math/big/example_test.go`](https://golang.org/src/math/big/example_test.go)
- [`$GOROOT/src/math/big/floatexample_test.go`](https://golang.org/src/math/big/floatexample_test.go)
- [`$GOROOT/src/math/rand/example_test.go`](https://golang.org/src/math/rand/example_test.go)
- [`$GOROOT/src/time/example_test.go`](https://golang.org/src/time/example_test.go)
- [`$GOROOT/src/unicode/example_test.go`](https://golang.org/src/unicode/example_test.go)
- [`$GOROOT/src/unicode/utf8/example_test.go`](https://golang.org/src/unicode/utf8/example_test.go)

## Internal Packages

TODO: Write this!

References:
- [Go 1.4 “Internal” Packages (Design Document)][designdoc-internal]
- [Go 1.4 Release Notes > Internal packages][go14relnotes-internal]
- [Go 1.5 Release Notes > Go command][go15relnotes-internal]
- [Command go > Internal Directories][cmdgo-internal]

[designdoc-internal]: https://golang.org/s/go14internal
[go14relnotes-internal]: https://golang.org/doc/go1.4#internalpackages
[go15relnotes-internal]: https://golang.org/doc/go1.5#go_command
[cmdgo-internal]: https://golang.org/cmd/go/#hdr-Internal_Directories
