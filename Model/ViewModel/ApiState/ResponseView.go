package ApiState

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type ResponseResult struct {
	Result bool        // 是否成功
	Msg    string      // 错误描述
	Code   int         // 错误码
	Data   interface{} // 返回数据
}

// 参数错误响应
func ArgErrApiResult(c *gin.Context, argName string) {
	c.JSON(SUCCESS, &ResponseResult{
		Result: false,
		Msg:    fmt.Sprintf("参数【%s】有误", argName),
		Code:   SUCCESS,
		Data:   nil,
	})
}

// 成功响应，带msg形参，msg为空设为success
func ResponseSuccessMsg(c *gin.Context, msg string, data interface{}) {
	if msg == "" {
		msg = "success"
	}
	c.JSON(SUCCESS, &ResponseResult{
		Result: true,
		Msg:    msg,
		Code:   SUCCESS,
		Data:   data,
	})
}

// 成功响应，不带msg形参，默认success
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(SUCCESS, &ResponseResult{
		Result: true,
		Msg:    "success",
		Code:   SUCCESS,
		Data:   data,
	})
}

// 模板响应
func ResponseHtml(c *gin.Context, path string, data interface{}) {
	c.HTML(SUCCESS, path, data)
}

// 普通业务异常
func BusinessErrApiResult(c *gin.Context, msg string) {
	c.JSON(SUCCESS, &ResponseResult{
		Result: false,
		Msg:    msg,
		Code:   SUCCESS,
		Data:   nil,
	})
}
