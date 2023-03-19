package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yy/pkg/e"
	"yy/service"
)

/*
	接口地址：/controller/user.Register
	功能描述：用户注册接口
	参数：
		param username 用户注册使用的用户名
		param password 用户注册使用的密码
	请求方式：POST
	作者：贺凯恒
	创建时间：2023/3/19
	Copyright2023
*/
func Register(c *gin.Context) {
	var userRegisterService service.UserService
	username := c.Query("username")
	password := c.Query("password")
	//检验用户名和密码格式
	if len(password) < 6 || len(username) < 3 {
		code := e.InvalidParams
		c.JSON(code, e.GetMsg(code))
		return
	}
	userRegisterService.UserName = username
	userRegisterService.Password = password
	res := userRegisterService.Register()
	c.JSON(http.StatusOK, res)
}

/*
	接口地址：/controller/user.Login
	功能描述：用户登录接口
	参数：
		param username 用户登录使用的用户名
		param password 用户登录使用的密码
	请求方式：POST
	作者：贺凯恒
	创建时间：2022/5/25
	Copyright2022
*/
func Login(c *gin.Context) {
	var userLoginService service.UserService
	username := c.Query("username")
	password := c.Query("password")
	//检验用户名和密码格式
	if len(password) < 6 || len(username) < 3 {
		code := e.InvalidParams
		c.JSON(code, e.GetMsg(code))
		return
	}
	userLoginService.UserName = username
	userLoginService.Password = password
	res := userLoginService.Login()
	c.JSON(http.StatusOK, res)
}
