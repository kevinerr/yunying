package service

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"os/exec"
	"time"
	"yy/model"
	"yy/pkg/e"
	"yy/pkg/util"
	"yy/repository"
	"yy/serializer"
)

type PublishService struct {
}

func (service *PublishService) Publish(token string, videoURL string, startTime string, endTime string, c *gin.Context) serializer.PublishResponse {
	var recordRepository repository.RecordRepository

	code := e.SUCCESS
	//身份判断
	claims, err := util.ParseToken(token)
	if err != nil {
		code = e.ErrorAuthCheckTokenFail
		return serializer.PublishResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	} else if time.Now().Unix() > claims.ExpiresAt {
		code = e.ErrorAuthCheckTokenTimeout
		return serializer.PublishResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	}
	userId := claims.Id

	//组合记录对象
	st := time.Now()
	snow := util.Snowflake{}
	recordId := snow.Generate()

	// 生成 ffmpeg 命令行参数
	cmdArgs := []string{"-i", videoURL, "-ss", startTime, "-to", endTime, "-c", "copy", "-y", "clip.mp4"}
	// 执行 ffmpeg 命令
	cmd := exec.Command("ffmpeg", cmdArgs...)
	err = cmd.Run()
	endUrl := "http://example.com/clip.mp4"
	et := time.Now()
	record := &model.Reocrd{
		Id:        recordId,
		UserId:    userId,
		StartTime: startTime,
		EndTime:   endTime,
		StartUrl:  st,
		EndUrl:    et,
	}
	//插库操作
	if err := recordRepository.CreatComment(record); err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.PublishResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	}
	if err != nil {
		code = e.SuccessUpLoadFile
		return serializer.PublishResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	}
	return serializer.PublishResponse{
		Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		VideoUrl: endUrl,
	}
}
