package main

// #cgo LDFLAGS: -L. -lcalc
// #include <calc.h>
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"log"
	"unsafe"

	"github.com/gin-gonic/gin"
)

func main() {
	//$ gcc -shared -fPIC -o libcalc.so calc.c
	//$ go run main.go
	router := gin.Default()

	plist := []C.person_t{
		{id: 1, size: 10},
		{id: 2, size: 20},
		{id: 3, size: 30},
		{id: 4, size: 40},
	}

	// ルートエンドポイント
	router.GET("/", func(c *gin.Context) {
		n := 100
		s := "sss"
		log.Printf("%v, %v\n", n, s)
		sum := C.data_sum((*C.person_t)(unsafe.Pointer(&plist[0])), C.size_t(len(plist)))
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Hello, World! sum=%v", sum),
		})
	})

	// サーバーを起動
	router.Run(":8080") // ポート8080でリッスン
}
