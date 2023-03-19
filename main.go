package main

import (
	"github.com/gin-gonic/gin"
	"yy/conf"
)

func main() {

	//从配置文件读入配置
	conf.Init()

	//转载路由 swag init -g common.go
	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
