package main

/*
#cgo CFLAGS: -I./c
#cgo LDFLAGS: -L. -ldemo
#include "libdemo.h"
*/
import "C"

func main() {
	C.printOkay()
}
