package RegisterController

import (
	"github.com/gin-gonic/gin"
	"golang-blog/Model/Entity"
	"golang-blog/Model/ViewModel/ApiState"
	"golang-blog/Service/RegisterService"
	"golang-blog/Service/UtilService"
	"strconv"
)

func RegisterUser(c *gin.Context) {
	loginName := c.PostForm("loginName")
	password := c.PostForm("password")
	if loginName == "" || password == "" {
		ApiState.ArgErrApiResult(c, "loginName or password")
		return
	}
	nickName := c.PostForm("nickName")
	ageStr := c.PostForm("age")
	age, _ := strconv.Atoi(ageStr)
	user := Entity.UserEntity{
		Id:        UtilService.UUId(),
		Age:       age,
		Password:  password,
		LoginName: loginName,
		NickName:  nickName,
	}
	RegisterService.RegisterUser(user)
	ApiState.ResponseSuccess(c, "注册成功")
}
