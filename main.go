package main

// #cgo CFLAGS: -I${SRCDIR}/libc
// #cgo LDFLAGS: ${SRCDIR}/libc.a
// #include <stdio.h>
// #include <helloworld.h>
import "C"

func main() {
	C.helloworld()
}
