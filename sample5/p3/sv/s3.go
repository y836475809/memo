package sv

import "github.com/gin-gonic/gin"

func SetupRouter3() *gin.Engine {
	router := gin.Default()
	router.GET("/ping3", func(c *gin.Context) {
		c.String(200, "pong2")
	})
	return router
}
