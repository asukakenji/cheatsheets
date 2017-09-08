# A collection of interfaces in Go's standard library

This is not an exhaustive list of all interfaces in Go's standard library.
I only list those I think are important.
Interfaces defined in frequently used packages (like `io`, `fmt`) are included.
Interfaces that have significant importance are also included.

All of the following information is based on `go version go1.8.3 darwin/amd64`.



### (builtin)

#### error [[doc](https://golang.org/pkg/builtin/#error)] [[src1](https://golang.org/src/builtin/builtin.go#L254)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/builtin/builtin.go#L254-L256)]

```go
type error interface {
	Error() string
}
```



### package `runtime`

#### Error [[doc](https://golang.org/pkg/runtime/#Error)] [[src1](https://golang.org/src/runtime/error.go#L8)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/runtime/error.go#L8-L16)]

```go
type Error interface {
	error
	RuntimeError()
}
```



### package `math/rand`

#### Source [[doc](https://golang.org/pkg/math/rand/#Source)] [[src1](https://golang.org/src/math/rand/rand.go#L21)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/math/rand/rand.go#L21-L24)]

```go
type Source interface {
	Int63() int64
	Seed(seed int64)
}
```

#### Source64 [[doc](https://golang.org/pkg/math/rand/#Source64)] [[src1](https://golang.org/src/math/rand/rand.go#L32)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/math/rand/rand.go#L32-L35)]

```go
type Source64 interface {
	Source
	Uint64() uint64
}
```



### package `sort`

#### Interface [[doc](https://golang.org/pkg/sort/#Interface)] [[src1](https://golang.org/src/sort/sort.go#L16)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/sort/sort.go#L16-L24)]

```go
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
```



### package `container/heap`

#### Interface [[doc](https://golang.org/pkg/container/heap/#Interface)] [[src1](https://golang.org/src/container/heap/heap.go#L30)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/container/heap/heap.go#L30-L34)]

```go
type Interface interface {
	sort.Interface
	Push(x interface{})
	Pop() interface{}
}
```



### package `io`

#### Reader [[doc](https://golang.org/pkg/io/#Reader)] [[src1](https://golang.org/src/io/io.go#L77)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/io/io.go#L77-L79)]

```go
type Reader interface {
	Read(p []byte) (n int, err error)
}
```

#### Writer [[doc](https://golang.org/pkg/io/#Writer)] [[src1](https://golang.org/src/io/io.go#L90)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/io/io.go#L90-L92)]

```go
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

#### Closer [[doc](https://golang.org/pkg/io/#Closer)] [[src1](https://golang.org/src/io/io.go#L98)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/io/io.go#L98-L100)]

```go
type Closer interface {
	Close() error
}
```

#### Seeker [[doc](https://golang.org/pkg/io/#Seeker)] [[src1](https://golang.org/src/io/io.go#L115)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/io/io.go#L115-L117)]

```go
type Seeker interface {
	Seek(offset int64, whence int) (int64, error)
}
```

#### ReadWriter [[doc](https://golang.org/pkg/io/#ReadWriter)] [[src1](https://golang.org/src/io/io.go#L120)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/io/io.go#L120-L123)]

```go
type ReadWriter interface {
	Reader
	Writer
}
```

#### ReadCloser [[doc](https://golang.org/pkg/io/#ReadCloser)] [[src1](https://golang.org/src/io/io.go#L126)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/io/io.go#L126-L129)]

```go
type ReadCloser interface {
	Reader
	Closer
}
```

#### WriteCloser [[doc](https://golang.org/pkg/io/#WriteCloser)] [[src1](https://golang.org/src/io/io.go#L132)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/io/io.go#L132-L135)]

```go
type WriteCloser interface {
	Writer
	Closer
}
```

#### ReadWriteCloser [[doc](https://golang.org/pkg/io/#ReadWriteCloser)] [[src1](https://golang.org/src/io/io.go#L138)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/io/io.go#L138-L142)]

```go
type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}
```

#### ReadSeeker [[doc](https://golang.org/pkg/io/#ReadSeeker)] [[src1](https://golang.org/src/io/io.go#L145)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/io/io.go#L145-L148)]

```go
type ReadSeeker interface {
	Reader
	Seeker
}
```

#### WriteSeeker [[doc](https://golang.org/pkg/io/#WriteSeeker)] [[src1](https://golang.org/src/io/io.go#L151)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/io/io.go#L151-L154)]

```go
type WriteSeeker interface {
	Writer
	Seeker
}
```

#### ReadWriteSeeker [[doc](https://golang.org/pkg/io/#ReadWriteSeeker)] [[src1](https://golang.org/src/io/io.go#L157)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/io/io.go#L157-L161)]

```go
type ReadWriteSeeker interface {
	Reader
	Writer
	Seeker
}
```

#### ReaderFrom [[doc](https://golang.org/pkg/io/#ReaderFrom)] [[src1](https://golang.org/src/io/io.go#L170)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/io/io.go#L170-L172)]

```go
type ReaderFrom interface {
	ReadFrom(r Reader) (n int64, err error)
}
```

#### WriterTo [[doc](https://golang.org/pkg/io/#WriterTo)] [[src1](https://golang.org/src/io/io.go#L181)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/io/io.go#L181-L183)]

```go
type WriterTo interface {
	WriteTo(w Writer) (n int64, err error)
}
```

#### ReaderAt [[doc](https://golang.org/pkg/io/#ReaderAt)] [[src1](https://golang.org/src/io/io.go#L211)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/io/io.go#L211-L213)]

```go
type ReaderAt interface {
	ReadAt(p []byte, off int64) (n int, err error)
}
```

#### WriterAt [[doc](https://golang.org/pkg/io/#WriterAt)] [[src1](https://golang.org/src/io/io.go#L230)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/io/io.go#L230-L232)]

```go
type WriterAt interface {
	WriteAt(p []byte, off int64) (n int, err error)
}
```

#### ByteReader [[doc](https://golang.org/pkg/io/#ByteReader)] [[src1](https://golang.org/src/io/io.go#L237)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/io/io.go#L237-L239)]

```go
type ByteReader interface {
	ReadByte() (byte, error)
}
```

#### ByteScanner [[doc](https://golang.org/pkg/io/#ByteScanner)] [[src1](https://golang.org/src/io/io.go#L248)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/io/io.go#L248-L251)]

```go
type ByteScanner interface {
	ByteReader
	UnreadByte() error
}
```

#### ByteWriter [[doc](https://golang.org/pkg/io/#ByteWriter)] [[src1](https://golang.org/src/io/io.go#L254)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/io/io.go#L254-L256)]

```go
type ByteWriter interface {
	WriteByte(c byte) error
}
```

#### RuneReader [[doc](https://golang.org/pkg/io/#RuneReader)] [[src1](https://golang.org/src/io/io.go#L263)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/io/io.go#L263-L265)]

```go
type RuneReader interface {
	ReadRune() (r rune, size int, err error)
}
```

#### RuneScanner [[doc](https://golang.org/pkg/io/#RuneScanner)] [[src1](https://golang.org/src/io/io.go#L274)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/io/io.go#L274-L277)]

```go
type RuneScanner interface {
	RuneReader
	UnreadRune() error
}
```



### package `fmt`

#### State [[doc](https://golang.org/pkg/fmt/#State)] [[src1](https://golang.org/src/fmt/print.go#L38)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/fmt/print.go#L38-L48)]

```go
type State interface {
	Write(b []byte) (n int, err error)
	Width() (wid int, ok bool)
	Precision() (prec int, ok bool)
	Flag(c int) bool
}
```

#### Formatter [[doc](https://golang.org/pkg/fmt/#Formatter)] [[src1](https://golang.org/src/fmt/print.go#L53)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/fmt/print.go#L53-L55)]

```go
type Formatter interface {
	Format(f State, c rune)
}
```

#### Stringer [[doc](https://golang.org/pkg/fmt/#Stringer)] [[src1](https://golang.org/src/fmt/print.go#L62)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/fmt/print.go#L62-L64)]

```go
type Stringer interface {
	String() string
}
```

#### GoStringer [[doc](https://golang.org/pkg/fmt/#GoStringer)] [[src1](https://golang.org/src/fmt/print.go#L70)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/fmt/print.go#L70-L72)]

```go
type GoStringer interface {
	GoString() string
}
```

#### ScanState [[doc](https://golang.org/pkg/fmt/#ScanState)] [[src1](https://golang.org/src/fmt/scan.go#L21)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/fmt/scan.go#L21-L49)]

```go
type ScanState interface {
	ReadRune() (r rune, size int, err error)
	UnreadRune() error
	SkipSpace()
	Token(skipSpace bool, f func(rune) bool) (token []byte, err error)
	Width() (wid int, ok bool)
	Read(buf []byte) (n int, err error)
}
```

#### Scanner [[doc](https://golang.org/pkg/fmt/#Scanner)] [[src1](https://golang.org/src/fmt/scan.go#L55)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/fmt/scan.go#L55-L57)]

```go
type Scanner interface {
	Scan(state ScanState, verb rune) error
}
```



### package `encoding`

#### BinaryMarshaler [[doc](https://golang.org/pkg/encoding/#BinaryMarshaler)] [[src1](https://golang.org/src/encoding/encoding.go#L18)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/encoding/encoding.go#L18-L20)]

```go
type BinaryMarshaler interface {
	MarshalBinary() (data []byte, err error)
}
```

#### BinaryUnmarshaler [[doc](https://golang.org/pkg/encoding/#BinaryUnmarshaler)] [[src1](https://golang.org/src/encoding/encoding.go#L28)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/encoding/encoding.go#L28-L30)]

```go
type BinaryUnmarshaler interface {
	UnmarshalBinary(data []byte) error
}
```

#### TextMarshaler [[doc](https://golang.org/pkg/encoding/#TextMarshaler)] [[src1](https://golang.org/src/encoding/encoding.go#L36)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/encoding/encoding.go#L36-L38)]

```go
type TextMarshaler interface {
	MarshalText() (text []byte, err error)
}
```

#### TextUnmarshaler [[doc](https://golang.org/pkg/encoding/#TextUnmarshaler)] [[src1](https://golang.org/src/encoding/encoding.go#L46)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/encoding/encoding.go#L46-L48)]

```go
type TextUnmarshaler interface {
	UnmarshalText(text []byte) error
}
```



### package `image`

#### Image [[doc](https://golang.org/pkg/image/#Image)] [[src1](https://golang.org/src/image/image.go#L36)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/image/image.go#L36-L46)]

```go
type Image interface {
	ColorModel() color.Model
	Bounds() Rectangle
	At(x, y int) color.Color
}
```

#### PalettedImage [[doc](https://golang.org/pkg/image/#PalettedImage)] [[src1](https://golang.org/src/image/image.go#L53)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/image/image.go#L53-L57)]

```go
type PalettedImage interface {
	ColorIndexAt(x, y int) uint8
	Image
}
```



### package `image/color`

#### Color [[doc](https://golang.org/pkg/image/color/#Color)] [[src1](https://golang.org/src/image/color/color.go#L10)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/image/color/color.go#L10-L19)]

```go
type Color interface {
	RGBA() (r, g, b, a uint32)
}
```

#### Model [[doc](https://golang.org/pkg/image/color/#Model)] [[src1](https://golang.org/src/image/color/color.go#L142)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/image/color/color.go#L142-L144)]

```go
type Model interface {
	Convert(c Color) Color
}
```



### package `image/draw`

#### Image [[doc](https://golang.org/pkg/image/draw/#Image)] [[src1](https://golang.org/src/image/draw/draw.go#L21)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/image/draw/draw.go#L21-L24)]

```go
type Image interface {
	image.Image
	Set(x, y int, c color.Color)
}
```

#### Quantizer [[doc](https://golang.org/pkg/image/draw/#Quantizer)] [[src1](https://golang.org/src/image/draw/draw.go#L27)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/image/draw/draw.go#L27-L31)]

```go
type Quantizer interface {
	Quantize(p color.Palette, m image.Image) color.Palette
}
```

#### Drawer [[doc](https://golang.org/pkg/image/draw/#Drawer)] [[src1](https://golang.org/src/image/draw/draw.go#L50)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/image/draw/draw.go#L50-L54)]

```go
type Drawer interface {
	Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point)
}
```



### package `hash`

#### Hash [[doc](https://golang.org/pkg/hash/#Hash)] [[src1](https://golang.org/src/hash/hash.go#L11)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/hash/hash.go#L11-L31)]

```go
type Hash interface {
	io.Writer
	Sum(b []byte) []byte
	Reset()
	Size() int
	BlockSize() int
}
```

#### Hash32 [[doc](https://golang.org/pkg/hash/#Hash32)] [[src1](https://golang.org/src/hash/hash.go#L34)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/hash/hash.go#L34-L37)]

```go
type Hash32 interface {
	Hash
	Sum32() uint32
}
```

#### Hash64 [[doc](https://golang.org/pkg/hash/#Hash64)] [[src1](https://golang.org/src/hash/hash.go#L40)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/hash/hash.go#L40-L43)]

```go
type Hash64 interface {
	Hash
	Sum64() uint64
}
```



### package `crypto`

#### Signer [[doc](https://golang.org/pkg/crypto/#Signer)] [[src1](https://golang.org/src/crypto/crypto.go#L107)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/crypto/crypto.go#L107-L126)]

```go
type Signer interface {
	Public() PublicKey
	Sign(rand io.Reader, digest []byte, opts SignerOpts) (signature []byte, err error)
}
```

#### SignerOpts [[doc](https://golang.org/pkg/crypto/#SignerOpts)] [[src1](https://golang.org/src/crypto/crypto.go#L129)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/crypto/crypto.go#L129-L134)]

```go
type SignerOpts interface {
	HashFunc() Hash
}
```

#### Decrypter [[doc](https://golang.org/pkg/crypto/#Decrypter)] [[src1](https://golang.org/src/crypto/crypto.go#L139)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/crypto/crypto.go#L139-L148)]

```go
type Decrypter interface {
	Public() PublicKey
	Decrypt(rand io.Reader, msg []byte, opts DecrypterOpts) (plaintext []byte, err error)
}
```



### package `reflect`

#### Type [[doc](https://golang.org/pkg/reflect/#Type)] [[src1](https://golang.org/src/reflect/type.go#L35)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/reflect/type.go#L35-L205)]

```go
type Type interface {
	Align() int
	FieldAlign() int
	Method(int) Method
	MethodByName(string) (Method, bool)
	NumMethod() int
	Name() string
	PkgPath() string
	Size() uintptr
	String() string
	Kind() Kind
	Implements(u Type) bool
	AssignableTo(u Type) bool
	ConvertibleTo(u Type) bool
	Comparable() bool
	Bits() int
	ChanDir() ChanDir
	IsVariadic() bool
	Elem() Type
	Field(i int) StructField
	FieldByIndex(index []int) StructField
	FieldByName(name string) (StructField, bool)
	FieldByNameFunc(match func(string) bool) (StructField, bool)
	In(i int) Type
	Key() Type
	Len() int
	NumField() int
	NumIn() int
	NumOut() int
	Out(i int) Type
	common() *rtype
	uncommon() *uncommonType
}
```



### package `os`

#### Signal [[doc](https://golang.org/pkg/os/#Signal)] [[src1](https://golang.org/src/os/exec.go#L64)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/os/exec.go#L64-L67)]

```go
type Signal interface {
	String() string
	Signal()
}
```

#### FileInfo [[doc](https://golang.org/pkg/os/#FileInfo)] [[src1](https://golang.org/src/os/types.go#L21)] [[src2](https://github.com/golang/go/blob/release-branch.go1.8/src/os/types.go#L21-L28)]

```go
type FileInfo interface {
	Name() string
	Size() int64
	Mode() FileMode
	ModTime() time.Time
	IsDir() bool
	Sys() interface{}
}
```





## Source

### iface.awk

```awk
BEGIN {
    if (package == "") {
        print "error: package is not defined"
        exit 1
    }

    if (branch == "") {
        print "error: branch is not defined"
        exit 1
    }

    state = 0
    indent0 = 0
    indent1 = 0
    from_line = 0
    to_line = 0
    filename = ""
    type = ""
    code = ""

    printf "\n"
    printf "\n"
    printf "\n"
    if (package == "builtin") {
        printf "### (builtin)\n"
    } else {
        printf "### package `%s`\n", package
    }
}

# Start of type
/type ([A-Z][^ ]*|error) interface {/ {
    if (state == 0) {
        state = 1

        indent0 = index($0, "type ")
        s = substr($0, indent0 + 5) # length("type ") == 5
        len = index(s, " ") - 1
        type = substr(s, 0, len)

        filename = FILENAME
        sub(/.*\//, "", filename)

        from_line = FNR
        code = ""
    }
}

# Inside type
{
    if (state == 1) {
        line = $0
        # Remove comments
        sub(/[\t ]*\/\/.*/, "", line)
        # Remove trailing whitespaces
        sub(/[\t ]*$/, "", line)
        # Only print non-blank lines
        if (line != "") {
            code = code line "\n"
        }
    }
}

# End of type
/}/ {
    if (state == 1) {
        indent1 = index($0, "}")
        if (indent0 == indent1) {
            state = 0
            to_line = FNR
            printf "\n"
            printf "#### %s " \
                "[[doc](https://golang.org/pkg/%s/#%s)] " \
                "[[src1](https://golang.org/src/%s/%s#L%d)] " \
                "[[src2](https://github.com/golang/go/blob/release-branch.%s/src/%s/%s#L%d-L%d)]\n",
                type,
                package, type,
                package, filename, from_line,
                branch, package, filename, from_line, to_line
            printf "\n"
            printf "```go\n"
            printf "%s", code
            printf "```\n"
        }
    }
}
```

### make.sh

```sh
#!/bin/sh

packages=(
    'builtin'
    'runtime'
    'math/rand'
    'sort'
    'container/heap'
    'io'
    'fmt'
    'encoding'
    'image'
    'image/color'
    'image/draw'
    'hash'
    'crypto'
    'reflect'
    'os'
)

if [ -z "${GOROOT}" ]
then
    eval $(go env | grep -e '^GOROOT=')
fi

if [ -z "${GOROOT}" ]
then
    echo 'Cannot find GOROOT'
    exit 1
fi

go_version=$(go version)
go_branch=${go_version#go version }
go_branch=${go_branch% *}

case ${go_branch} in
go[0-9].[0-9])
    ;;
go[0-9].[0-9].[0-9])
    go_branch=${go_branch%.[0-9]}
    ;;
*)
    printf 'Unexpected go version: %s\n' ${go_version}
    ;;
esac

echo "# A collection of interfaces in Go's standard library"
echo
echo "This is not an exhaustive list of all interfaces in Go's standard library."
echo 'I only list those I think are important.'
echo 'Interfaces defined in frequently used packages (like `io`, `fmt`) are included.'
echo 'Interfaces that have significant importance are also included.'
echo
printf 'All of the following information is based on `%s`.\n' "$(go version)"

for package in ${packages[@]}
do
    find ${GOROOT}/src/${package} -maxdepth 1 \
        -type f '(' \
            -name '*_test.go' -prune -o \
            -name '*.go' -exec \
                awk -f iface.awk -v package="${package}" -v branch="${go_branch}" '{}' '+' \
        ')'
done

printf '\n'
printf '\n'
printf '\n'
printf '\n'
printf '\n'
printf '## Source\n'
printf '\n'

printf '### iface.awk\n'
printf '\n'
printf '```awk\n'
cat iface.awk
printf '```\n'
printf '\n'

printf '### make.sh\n'
printf '\n'
printf '```sh\n'
cat make.sh
printf '```\n'
```
