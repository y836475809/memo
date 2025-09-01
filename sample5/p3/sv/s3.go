package sv

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter3() *gin.Engine {
	router := gin.Default()
	router.GET("/ping3", func(c *gin.Context) {
		c.String(200, "pong3")
	})

	// ping3-2 -> p2 /ping2
	router.GET("/ping3-2", func(c *gin.Context) {
		r, err := http.Get("http://localhost:8088/ping2")
		if err != nil {
			c.String(500, "error pong3")
		} else {
			buf := make([]byte, 2048)
			n, _ := r.Body.Read(buf)
			b := string(buf[0:n])
			fmt.Println(b)
			c.String(200, b+"_pong3")
		}
	})
	return router
}
