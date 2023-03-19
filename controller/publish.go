package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yy/service"
)

/*
	接口地址：/controller/publish.Publish
	功能描述：登录用户选择视频上传
	详细描述：检查 token 然后上传视频到 public 目录下
	参数：
		param videoURL 登陆的用户
		param startTime 视频数据信息
		param endTime 视频标题
	请求方式：POST
	作者：贺凯恒
	创建时间：2023/3/19
	Copyright2023
*/
func Publish(c *gin.Context) {
	var publishService service.PublishService
	token := c.PostForm("token")
	videoURL := c.PostForm("videoURL")
	startTime := c.PostForm("startTime")
	endTime := c.PostForm("endTime")
	res := publishService.Publish(token, videoURL, startTime, endTime, c)
	c.JSON(http.StatusOK, res)
}
