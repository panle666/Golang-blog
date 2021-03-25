package HomeControoler

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
	c.JSON(200, "hello world")
}

func Hi(c *gin.Context) {
	c.JSON(200, "hi")
}
