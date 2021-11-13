package main

// #cgo CFLAGS: -I${SRCDIR}/libc
// #cgo LDFLAGS: ${SRCDIR}/libc.a
// #include <stdio.h>
// #include <helloworld.h>
// #include <calc.h>
import "C"
import "fmt"

func main() {
	C.helloworld()
	addres := C.add(10, 20)
	fmt.Println("C.add result: ", addres)
	divres := C.div(10, 20)
	fmt.Println("C.add result: ", divres)
	modres := C.mod(10, 20)
	fmt.Println("C.add result: ", modres)
}
