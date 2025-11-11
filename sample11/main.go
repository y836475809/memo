package main

// #cgo LDFLAGS: -L. -lcalc
// #include <calc.h>
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	// mingw
	//$ gcc -g -shared -fPIC -o libcalc.dll calc.c
	plist := []C.person_t{
		{id: 1, size: 10},
		{id: 2, size: 20},
		{id: 3, size: 30},
		{id: 4, size: 40},
	}
	sum := C.data_sum((*C.person_t)(unsafe.Pointer(&plist[0])), C.size_t(len(plist)))
	fmt.Printf("%v", sum)
}

func calcSum() float64 {
	plist := []C.person_t{
		{id: 1, size: 10},
		{id: 2, size: 20},
		{id: 3, size: 30},
		{id: 4, size: 40},
	}
	sum := C.data_sum((*C.person_t)(unsafe.Pointer(&plist[0])), C.size_t(len(plist)))
	return float64(sum)
}
