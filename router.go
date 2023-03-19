package main

import (
	"github.com/gin-gonic/gin"
	"yy/controller"
)

func initRouter(r *gin.Engine) {
	r.Static("/static", "./public")
	apiRouter := r.Group("/yunyin")

	// basic apis
	apiRouter.GET("/publish/", controller.Publish)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)

}
