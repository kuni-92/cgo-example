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
	fmt.Println("C.div result: ", divres)
	modres := C.mod(10, 20)
	fmt.Println("C.mod result: ", modres)

	numbers := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
	var sum int
	for _, n := range numbers {
		sum = int(C.add(C.int(sum), C.int(n)))
	}
	fmt.Println("sum 1 to 10 result: ", sum)
}
