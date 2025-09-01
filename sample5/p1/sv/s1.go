package sv

import "github.com/gin-gonic/gin"

type User struct {
	Username string `json:"username"`
	Gender   string `json:"gender"`
}

func SetupRouter1() *gin.Engine {
	router := gin.Default()
	router.GET("/ping1", func(c *gin.Context) {
		c.String(200, "pong1")
	})
	return router
}

func PostUser(r *gin.Engine) *gin.Engine {
	r.POST("/user/add", func(c *gin.Context) {
		var user User
		c.BindJSON(&user)
		c.JSON(200, user)
	})
	return r
}
