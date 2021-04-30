package UserController

import (
	"github.com/gin-gonic/gin"
	"golang-blog/Model/ViewModel/ApiState"
	"golang-blog/Service/UserService"
)

func GetUserInfo(c *gin.Context) {
	loginName := c.Query("loginName")
	password := c.Query("password")
	if loginName == "" || password == "" {
		ApiState.ArgErrApiResult(c, "loginName or password")
		return
	}
	userInfo := UserService.GetUserInfo(loginName, password)
	ApiState.ResponseSuccess(c, userInfo)
}
