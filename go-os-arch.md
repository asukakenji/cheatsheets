# Go (Golang) GOOS and GOARCH

All of the following results are taken from a `darwin/amd64` environment.

## A list of valid GOOS values

(Bold = supported by `go` out of the box, ie. without the help of a C compiler, etc.)

- `android`
- **`darwin`**
- **`dragonfly`**
- **`freebsd`**
- **`linux`**
- **`nacl`**
- **`netbsd`**
- **`openbsd`**
- **`plan9`**
- **`solaris`**
- **`windows`**
- `zos`

## A list of valid GOARCH values

(Bold = supported by `go` out of the box, ie. without the help of a C compiler, etc.)

- **`386`**
- **`amd64`**
- **`amd64p32`**
- **`arm`**
- `armbe`
- **`arm64`**
- `arm64be`
- **`ppc64`**
- **`ppc64le`**
- **`mips`**
- **`mipsle`**
- **`mips64`**
- **`mips64le`**
- `mips64p32`
- `mips64p32le`
- `ppc`
- `s390`
- **`s390x`**
- `sparc`
- `sparc64`

## A list of GOOS/GOARCH supported by `go` out of the box

- `darwin/386`
- `darwin/amd64`
- `dragonfly/amd64`
- `freebsd/386`
- `freebsd/amd64`
- `freebsd/arm`
- `linux/386`
- `linux/amd64`
- `linux/arm`
- `linux/arm64`
- `linux/ppc64`
- `linux/ppc64le`
- `linux/mips`
- `linux/mipsle`
- `linux/mips64`
- `linux/mips64le`
- `linux/s390x`
- `nacl/386`
- `nacl/amd64p32`
- `nacl/arm`
- `netbsd/386`
- `netbsd/amd64`
- `netbsd/arm`
- `openbsd/386`
- `openbsd/amd64`
- `openbsd/arm`
- `plan9/386`
- `plan9/amd64`
- `plan9/arm`
- `solaris/amd64`
- `windows/386`
- `windows/amd64`

## Support Grid

|                   | `a` | `m` | `d` | `f` | `l` | `c` | `n` | `o` | `p` | `s` | `w` | `z` |
| ----------------- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| **`386`**         | `A` | `O` |     | `O` | `O` | `O` | `O` | `O` | `O` |     | `O` |     |
| **`amd64`**       | `A` | `O` | `O` | `O` | `O` |     | `O` | `O` | `O` | `O` | `O` |     |
| **`amd64p32`**    |     |     |     |     |     | `O` |     |     |     |     |     |     |
| **`arm`**         | `A` | `B` |     | `O` | `O` | `O` | `O` | `O` | `O` |     |     |     |
| **`armbe`**       |     |     |     |     |     |     |     |     |     |     |     |     |
| **`arm64`**       | `A` | `C` |     |     | `O` |     |     |     |     |     |     |     |
| **`arm64be`**     |     |     |     |     |     |     |     |     |     |     |     |     |
| **`ppc64`**       |     |     |     |     | `O` |     |     |     |     |     |     |     |
| **`ppc64le`**     |     |     |     |     | `O` |     |     |     |     |     |     |     |
| **`mips`**        |     |     |     |     | `O` |     |     |     |     |     |     |     |
| **`mipsle`**      |     |     |     |     | `O` |     |     |     |     |     |     |     |
| **`mips64`**      |     |     |     |     | `O` |     |     |     |     |     |     |     |
| **`mips64le`**    |     |     |     |     | `O` |     |     |     |     |     |     |     |
| **`mips64p32`**   |     |     |     |     |     |     |     |     |     |     |     |     |
| **`mips64p32le`** |     |     |     |     |     |     |     |     |     |     |     |     |
| **`ppc`**         |     |     |     |     |     |     |     |     |     |     |     |     |
| **`s390`**        |     |     |     |     |     |     |     |     |     |     |     |     |
| **`s390x`**       |     |     |     |     | `O` |     |     |     |     |     |     |     |
| **`sparc`**       |     |     |     |     |     |     |     |     |     |     |     |     |
| **`sparc64`**     |     |     |     |     |     |     |     |     |     |     |     |     |

### Key

`a` = `android`, `m` = `darwin` (`macos`), `d` = `dragonfly`, `f` = `freebsd`,
`l` = `linux`, `c` = `nacl`, `n` = `netbsd`, `o` = `openbsd`, `p` = `plan9`,
`s` = `solaris`, `w` = `windows`, `z` = `zos`

(blank): Unsupported

`O`: Supported

`A`:
```
warning: unable to find runtime/cgo.a
/usr/local/go/pkg/tool/darwin_amd64/link: running clang failed: exit status 1
ld: unknown option: -z
clang: error: linker command failed with exit code 1 (use -v to see invocation)
```

`B`:
```
warning: unable to find runtime/cgo.a
/usr/local/go/pkg/tool/darwin_amd64/link: running clang failed: exit status 1
ld: warning: ignoring file /var/folders/dd/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx/T/go-link-xxxxxxxxx/go.o, file was built for armv7 which is not the architecture being linked (x86_64): /var/folders/dd/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx/T/go-link-xxxxxxxxx/go.o
Undefined symbols for architecture x86_64:
  "_main", referenced from:
     implicit entry/start for main executable
ld: symbol(s) not found for architecture x86_64
clang: error: linker command failed with exit code 1 (use -v to see invocation)
```

`C`:
```
warning: unable to find runtime/cgo.a
/usr/local/go/pkg/tool/darwin_amd64/link: running clang failed: exit status 1
ld: warning: ignoring file /var/folders/dd/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx/T/go-link-xxxxxxxxx/go.o, file was built for unsupported file format ( 0xCF 0xFA 0xED 0xFE 0x0C 0x00 0x00 0x01 0x00 0x00 0x00 0x00 0x01 0x00 0x00 0x00 ) which is not the architecture being linked (x86_64): /var/folders/dd/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx/T/go-link-xxxxxxxxx/go.o
Undefined symbols for architecture x86_64:
  "_main", referenced from:
     implicit entry/start for main executable
ld: symbol(s) not found for architecture x86_64
clang: error: linker command failed with exit code 1 (use -v to see invocation)
```

## Script

```sh
#!/bin/sh

# Reference:
# https://github.com/golang/go/blob/master/src/go/build/syslist.go
for goos in android darwin dragonfly freebsd linux nacl netbsd openbsd plan9 solaris windows zos
do
    for goarch in 386 amd64 amd64p32 arm armbe arm64 arm64be ppc64 ppc64le mips mipsle mips64 mips64le mips64p32 mips64p32le ppc s390 s390x sparc sparc64
    do
        GOOS=${goos} GOARCH=${goarch} go build -o /dev/null . >/dev/null 2>&1
        if [ $? -eq 0 ]
        then
            echo "${goos}/${goarch}"
        fi
    done
done
```
