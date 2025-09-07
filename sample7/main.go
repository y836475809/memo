package main

// #cgo LDFLAGS: -L. -lcalc
// #include <calc.h>
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	//$ gcc -c calc.c
	//$ ar rusv libcalc.a calc.o
	//$ go run main.go

	length := 1
	people := make([]C.person_t, length)
	people[0].id = C.int(10)

	dataList := []C.data_t{
		{a: 10},
		{a: 20},
		{a: 30},
	}
	people[0].size = C.int(len(dataList))
	people[0].data = (*C.data_t)(unsafe.Pointer(&dataList[0]))

	cStrctArray := C.allocStruct(C.int(length))
	defer C.free(unsafe.Pointer(cStrctArray))
	for i, strct := range people {
		C.convertStructData(cStrctArray, C.int(i), strct)
	}

	fmt.Println("len(people)=", len(people))
	// sum := C.data_sum(unsafe.SliceData(people), C.size_t(len(people)))
	// sum := C.data_sum((*C.person_t)(unsafe.Pointer(&people[0])), C.size_t(len(people)))
	sum := C.data_sum(cStrctArray, C.size_t(len(people)))
	fmt.Printf("sum: %f\n", sum)
}
