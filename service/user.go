package service

import (
	logging "github.com/sirupsen/logrus"
	"yy/pkg/e"
	"yy/pkg/util"
	"yy/repository"
	"yy/serializer"
)

//-------------------------------------------------------
//项目结构路径：/service/user.go
//创建者：贺凯恒
//描述：用户登录、注册功能相关的service层
//创建时间：2023/3/19
//Copyright2023
//--------------------------------------------------------

//UserRegisterService 用户服务
type UserService struct {
	UserName string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func (service *UserService) Register() serializer.UserLoginResponse {
	code := e.SUCCESS
	var userRegisterRepository repository.UserRepository               //Dao层
	user, flag := userRegisterRepository.IsExistUser(service.UserName) //判断用户名是否存在，存在flag返回true，否则返回false
	//表单验证
	if flag {
		code = e.ErrorExistUser
		return serializer.UserLoginResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	}
	user.Username = service.UserName
	//加密密码
	if err := user.SetPassword(service.Password); err != nil {
		logging.Info(err)
		code = e.ErrorFailEncryption
		return serializer.UserLoginResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	}
	//雪花算法生成ID
	snow := util.Snowflake{}
	user.Id = snow.Generate()
	//生成token
	token, err := util.GenerateToken(user.Id, service.UserName, 0)
	if err != nil {
		logging.Info(err)
		code = e.ErrorAuthToken
		return serializer.UserLoginResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	}
	//创建用户
	if err := userRegisterRepository.CreateUser(user); err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.UserLoginResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	}
	return serializer.UserLoginResponse{
		Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		UserId:   user.Id,
		Token:    token,
	}
}

//Login 用户登陆函数
func (service *UserService) Login() serializer.UserLoginResponse {
	code := e.SUCCESS
	var userLoginRepository repository.UserRepository
	user, flag := userLoginRepository.IsExistUser(service.UserName) //判断用户名是否存在，存在flag返回true，否则返回false
	if !flag {
		//如果查询不到，返回相应的错误
		code = e.ErrorNotExistUser
		return serializer.UserLoginResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	}
	if user.CheckPassword(service.Password) == false { //检查密码
		code = e.ErrorNotCompare
		return serializer.UserLoginResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	}
	token, err := util.GenerateToken(user.Id, service.UserName, 0) //生成token
	if err != nil {
		logging.Info(err)
		code = e.ErrorAuthToken
		return serializer.UserLoginResponse{
			Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		}
	}
	return serializer.UserLoginResponse{
		Response: serializer.Response{StatusCode: code, StatusMsg: e.GetMsg(code)},
		UserId:   user.Id,
		Token:    token,
	}
}
