package sv

import "github.com/gin-gonic/gin"

func SetupRouter2() *gin.Engine {
	router := gin.Default()
	router.GET("/ping2", func(c *gin.Context) {
		c.String(200, "pong2")
	})
	return router
}
