package main

// #cgo LDFLAGS: -L. -lhello
// #include <hello.h>
import "C"
import "fmt"

func main() {
	fmt.Printf("Hello World\n")
	C.hello()
}
