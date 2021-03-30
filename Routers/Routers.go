package Routers

import (
	"github.com/gin-gonic/gin"
	"golang-blog/Controller/HomeControler"
	"golang-blog/Service/ConfigService"
)

func Init(router *gin.Engine) {
	home := router.Group("Home")

	// 1.首位多余元素会被删除(../ or //);
	//2.然后路由会对新的路径进行不区分大小写的查找;
	//3.如果能正常找到对应的handler，路由就会重定向到正确的handler上并返回301或者307.(比如: 用户访问/FOO 和 /..//Foo可能会被重定向到/foo这个路由上)
	router.RedirectFixedPath = true

	{
		home.GET("/", HomeControler.Index)
		home.GET("/Hi", HomeControler.Hi)
	}

	serverConfig := ConfigService.GetServerConfig()
	router.Run(serverConfig.HTTP_PORT) // 监听并在 127.0.0.1:8888 上启动服务
}
