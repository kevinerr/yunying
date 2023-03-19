package repository

import (
	"yy/model"
)

type UserRepository struct {
}

//创建一个用户
func (c UserRepository) CreateUser(user *model.User) error {
	err := model.DB.Create(user).Error
	return err
}

//判断用户名是否存在，存在flag返回true，否则返回false
func (c UserRepository) IsExistUser(username string) (*model.User, bool) {
	var user model.User
	var count int64
	model.DB.Where("username=?", username).First(&user).Count(&count)
	if count == 1 {
		return &user, true
	}
	return &user, false
}
