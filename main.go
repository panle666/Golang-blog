package main

import (
	"github.com/gin-gonic/gin"
	"golang-blog/Routers"
	"golang-blog/Service/DbService"
)

func main() {
	router := gin.Default()
	DbService.ConnectDb()
	Routers.Init(router)
}
