# A collection of interfaces in Go's standard library

This is not an exhaustive list of all interfaces in Go's standard library.
I only list those I think are important.
Interfaces defined in frequently used packages (like `io`, `fmt`) are included.
Interfaces that have significant importance are also included.



### (builtin)

#### [error](https://golang.org/pkg/builtin/#error)
```go
type error interface {
	Error() string
}
```



### package `runtime`

#### [Error](https://golang.org/pkg/runtime/#Error)

```go
type Error interface {
	error
	RuntimeError()
}
```



### package `math/rand`

#### [Source](https://golang.org/pkg/math/rand/#Source)

```go
type Source interface {
	Int63() int64
	Seed(seed int64)
}
```



### package `sort`

#### [Interface](https://golang.org/pkg/sort/#Interface)

```go
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
```



### package `container/heap`

#### [Interface](https://golang.org/pkg/container/heap/#Interface)

```go
type Interface interface {
	sort.Interface
	Push(x interface{})
	Pop() interface{}
}
```



### package `io`

#### [Reader](https://golang.org/pkg/io/#Reader)

```go
type Reader interface {
	Read(p []byte) (n int, err error)
}
```

#### [Writer](https://golang.org/pkg/io/#Writer)

```go
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

#### [Closer](https://golang.org/pkg/io/#Closer)

```go
type Closer interface {
	Close() error
}
```

#### [Seeker](https://golang.org/pkg/io/#Seeker)

```go
type Seeker interface {
	Seek(offset int64, whence int) (int64, error)
}
```

#### [ReadWriter](https://golang.org/pkg/io/#ReadWriter)

```go
type ReadWriter interface {
	Reader
	Writer
}
```

#### [ReadCloser](https://golang.org/pkg/io/#ReadCloser)

```go
type ReadCloser interface {
	Reader
	Closer
}
```

#### [WriteCloser](https://golang.org/pkg/io/#WriteCloser)

```go
type WriteCloser interface {
	Writer
	Closer
}
```

#### [ReadWriteCloser](https://golang.org/pkg/io/#ReadWriteCloser)

```go
type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}
```

#### [ReadSeeker](https://golang.org/pkg/io/#ReadSeeker)

```go
type ReadSeeker interface {
	Reader
	Seeker
}
```

#### [WriteSeeker](https://golang.org/pkg/io/#WriteSeeker)

```go
type WriteSeeker interface {
	Writer
	Seeker
}
```

#### [ReadWriteSeeker](https://golang.org/pkg/io/#ReadWriteSeeker)

```go
type ReadWriteSeeker interface {
	Reader
	Writer
	Seeker
}
```

#### [ReaderFrom](https://golang.org/pkg/io/#ReaderFrom)

```go
type ReaderFrom interface {
	ReadFrom(r Reader) (n int64, err error)
}
```

#### [WriterTo](https://golang.org/pkg/io/#WriterTo)

```go
type WriterTo interface {
	WriteTo(w Writer) (n int64, err error)
}
```

#### [ReaderAt](https://golang.org/pkg/io/#ReaderAt)

```go
type ReaderAt interface {
	ReadAt(p []byte, off int64) (n int, err error)
}
```

#### [WriterAt](https://golang.org/pkg/io/#WriterAt)

```go
type WriterAt interface {
	WriteAt(p []byte, off int64) (n int, err error)
}
```

#### [ByteReader](https://golang.org/pkg/io/#ByteReader)

```go
type ByteReader interface {
	ReadByte() (c byte, err error)
}
```

#### [ByteScanner](https://golang.org/pkg/io/#ByteScanner)

```go
type ByteScanner interface {
	ByteReader
	UnreadByte() error
}
```

#### [ByteWriter](https://golang.org/pkg/io/#ByteWriter)

```go
type ByteWriter interface {
	WriteByte(c byte) error
}
```

#### [RuneReader](https://golang.org/pkg/io/#RuneReader)

```go
type RuneReader interface {
	ReadRune() (r rune, size int, err error)
}
```

#### [RuneScanner](https://golang.org/pkg/io/#RuneScanner)

```go
type RuneScanner interface {
	RuneReader
	UnreadRune() error
}
```



### package `fmt`

#### [State](https://golang.org/pkg/fmt/#State)

```go
type State interface {
	Write(b []byte) (ret int, err error)
	Width() (wid int, ok bool)
	Precision() (prec int, ok bool)
	Flag(c int) bool
}
```

#### [Formatter](https://golang.org/pkg/fmt/#Formatter)

```go
type Formatter interface {
	Format(f State, c rune)
}
```

#### [Stringer](https://golang.org/pkg/fmt/#Stringer)

```go
type Stringer interface {
	String() string
}
```

#### [GoStringer](https://golang.org/pkg/fmt/#GoStringer)

```go
type GoStringer interface {
	GoString() string
}
```

#### [ScanState](https://golang.org/pkg/fmt/#ScanState)

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

#### [Scanner](https://golang.org/pkg/fmt/#Scanner)

```go
type Scanner interface {
	Scan(state ScanState, verb rune) error
}
```



### package `encoding`

#### [BinaryMarshaler](https://golang.org/pkg/encoding/#BinaryMarshaler)

```go
type BinaryMarshaler interface {
	MarshalBinary() (data []byte, err error)
}
```

#### [BinaryUnmarshaler](https://golang.org/pkg/encoding/#BinaryUnmarshaler)

```go
type BinaryUnmarshaler interface {
	UnmarshalBinary(data []byte) error
}
```

#### [TextMarshaler](https://golang.org/pkg/encoding/#TextMarshaler)

```go
type TextMarshaler interface {
	MarshalText() (text []byte, err error)
}
```

#### [TextUnmarshaler](https://golang.org/pkg/encoding/#TextUnmarshaler)

```go
type TextUnmarshaler interface {
	UnmarshalText(text []byte) error
}
```



### package `image`

#### [Image](https://golang.org/pkg/image/#Image)

```go
type Image interface {
	ColorModel() color.Model
	Bounds() Rectangle
	At(x, y int) color.Color
}
```

#### [PalettedImage](https://golang.org/pkg/image/#PalettedImage)

```go
type PalettedImage interface {
	ColorIndexAt(x, y int) uint8
	Image
}
```



### package `image/color`

#### [Color](https://golang.org/pkg/image/color/#Color)

```go
type Color interface {
	RGBA() (r, g, b, a uint32)
}
```

#### [Model](https://golang.org/pkg/image/color/#Model)

```go
type Model interface {
	Convert(c Color) Color
}
```



### package `image/draw`

#### [Image](https://golang.org/pkg/image/draw/#Image)

```go
type Image interface {
	image.Image
	Set(x, y int, c color.Color)
}
```

#### [Quantizer](https://golang.org/pkg/image/draw/#Quantizer)

```go
type Quantizer interface {
	Quantize(p color.Palette, m image.Image) color.Palette
}
```

#### [Drawer](https://golang.org/pkg/image/draw/#Drawer)

```go
type Drawer interface {
	Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point)
}
```



### package `hash`

#### [Hash](https://golang.org/pkg/hash/#Hash)

```go
type Hash interface {
	io.Writer
	Sum(b []byte) []byte
	Reset()
	Size() int
	BlockSize() int
}
```

#### [Hash32](https://golang.org/pkg/hash/#Hash32)

```go
type Hash32 interface {
	Hash
	Sum32() uint32
}
```

#### [Hash64](https://golang.org/pkg/hash/#Hash64)

```go
type Hash64 interface {
	Hash
	Sum64() uint64
}
```



### package `crypto`

#### [Signer](https://golang.org/pkg/crypto/#Signer)

```go
type Signer interface {
	Public() PublicKey
	Sign(rand io.Reader, msg []byte, opts SignerOpts) (signature []byte, err error)
}
```

#### [SignerOpts](https://golang.org/pkg/crypto/#SignerOpts)

```go
type SignerOpts interface {
	HashFunc() Hash
}
```

#### [Decrypter](https://golang.org/pkg/crypto/#Decrypter)

```go
type Decrypter interface {
	Public() PublicKey
	Decrypt(rand io.Reader, msg []byte, opts DecrypterOpts) (plaintext []byte, err error)
}
```



### package `reflect`

#### [Type](https://golang.org/pkg/reflect/#Type)

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
}
```



### package `go/types`

#### [Type](https://golang.org/pkg/go/types/#Type)

```go
type Type interface {
	Underlying() Type
	String() string
}
```

#### [Object](https://golang.org/pkg/go/types/#Object)

```go
type Object interface {
	Parent() *Scope
	Pos() token.Pos
	Pkg() *Package
	Name() string
	Type() Type
	Exported() bool
	Id() string
	String() string
}
```

#### [Sizes](https://golang.org/pkg/go/types/#Sizes)

```go
type Sizes interface {
	Alignof(T Type) int64
	Offsetsof(fields []*Var) []int64
	Sizeof(T Type) int64
}
```

#### [Importer](https://golang.org/pkg/go/types/#Importer)

```go
type Importer interface {
	Import(path string) (*Package, error)
}
```



### package `go/constant`

#### [Value](https://golang.org/pkg/go/constant/#Value)

```go
type Value interface {
	Kind() Kind
	String() string
}
```



### package `os`

#### [FileInfo](https://golang.org/pkg/os/#FileInfo)

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

#### [Signal](https://golang.org/pkg/os/#Signal)

```go
type Signal interface {
	String() string
	Signal()
}
```
