package HomeControler

import (
	"github.com/gin-gonic/gin"
	"golang-blog/Model/ViewModel/ApiState"
)

func Index(c *gin.Context) {
	ApiState.ResponseSuccess(c, "hello world")
}

func Hi(c *gin.Context) {

	// 获取参数name
	name := c.Query("name")

	// 参数name为空响应参数错误提示
	if name == "" {
		ApiState.ArgErrApiResult(c, "name")
		return
	}
	ApiState.ResponseSuccess(c, "Hi")
}
