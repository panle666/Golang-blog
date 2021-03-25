package main

import (
	"github.com/gin-gonic/gin"
	"golang-blog/Routers"
)

func main() {
	router := gin.Default()
	Routers.Init(router)
}
