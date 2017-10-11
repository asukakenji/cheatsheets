# Go (Golang) Testing, Benchmark, and Example Facilities

## Overall

References:
- [Package testing](https://golang.org/pkg/testing/)

## White Box Testing

```go
package xxx

import "testing"

func TestXxx(t *testing.T) {
	/* ... */
}
```

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
- [Command go > Description of testing functions](https://golang.org/cmd/go/#hdr-Description_of_testing_functions)

## Black Box Testing

```go
package xxx_test

import "testing"

func TestXxx(t *testing.T) {
	/* ... */
}
```

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
- [Command go > Test packages](https://golang.org/cmd/go/#hdr-Test_packages)

## Gray Box Testing (Backdoors)

```go
package xxx

var (
	Vvv1 = vvv1
	Vvv2 = vvv2
)
```

Filename (Not enforced by the toolchain):
- `export_test.go`

Filename Examples:

- [`$GOROOT/src/bufio/export_test.go`](https://golang.org/src/bufio/export_test.go)
- [`$GOROOT/src/bytes/export_test.go`](https://golang.org/src/bytes/export_test.go)
- [`$GOROOT/src/math/export_test.go`](https://golang.org/src/math/export_test.go)
- [`$GOROOT/src/net/http/export_test.go`](https://golang.org/src/net/http/export_test.go)
- [`$GOROOT/src/os/export_test.go`](https://golang.org/src/os/export_test.go)
- [`$GOROOT/src/reflect/export_test.go`](https://golang.org/src/reflect/export_test.go)
- [`$GOROOT/src/runtime/export_arm_test.go`](https://golang.org/src/runtime/export_arm_test.go)
- [`$GOROOT/src/runtime/export_linux_test.go`](https://golang.org/src/runtime/export_linux_test.go)
- [`$GOROOT/src/runtime/export_mmap_test.go`](https://golang.org/src/runtime/export_mmap_test.go)
- [`$GOROOT/src/runtime/export_test.go`](https://golang.org/src/runtime/export_test.go)
- [`$GOROOT/src/runtime/export_windows_test.go`](https://golang.org/src/runtime/export_windows_test.go)
- [`$GOROOT/src/sync/export_test.go`](https://golang.org/src/sync/export_test.go)

Note:
- The backdoor file exposes internals that are needed by the test package;
- The package declaration in the backdoor file is the same as that in the source files being tested;
- The backdoor file does not contain any tests.

## Test Functions

```go
func TestXxx(t *testing.T) {
	/* ... */
}
```

Filename:
- `*_test.go`

Filename Examples:
- [`$GOROOT/src/math/big/bits_test.go`](https://golang.org/src/math/big/bits_test.go)
- [`$GOROOT/src/math/big/calibrate_test.go`](https://golang.org/src/math/big/calibrate_test.go)
- [`$GOROOT/src/math/big/intconv_test.go`](https://golang.org/src/math/big/intconv_test.go)
- [`$GOROOT/src/math/big/intmarsh_test.go`](https://golang.org/src/math/big/intmarsh_test.go)
- [`$GOROOT/src/math/rand/race_test.go`](https://golang.org/src/math/rand/race_test.go)
- [`$GOROOT/src/time/format_test.go`](https://golang.org/src/time/format_test.go)
- [`$GOROOT/src/time/zoneinfo_test.go`](https://golang.org/src/time/zoneinfo_test.go)
- [`$GOROOT/src/time/zoneinfo_windows_test.go`](https://golang.org/src/time/zoneinfo_windows_test.go)
- [`$GOROOT/src/unicode/digit_test.go`](https://golang.org/src/unicode/digit_test.go)
- [`$GOROOT/src/unicode/graphic_test.go`](https://golang.org/src/unicode/graphic_test.go)
- [`$GOROOT/src/unicode/letter_test.go`](https://golang.org/src/unicode/letter_test.go)
- [`$GOROOT/src/unicode/script_test.go`](https://golang.org/src/unicode/script_test.go)

## Benchmark Functions

```go
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

## Example Functions

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
