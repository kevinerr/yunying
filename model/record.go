package model

import "time"

type Reocrd struct {
	Id        int64     `gorm:"column:id"`         //记录ID
	UserId    int64     `gorm:"column:user_id"`    //用户ID
	StartTime string    `gorm:"column:start_time"` //开始时间
	EndTime   string    `gorm:"column:end_time"`   //结束时间
	StartUrl  time.Time `gorm:"column:start_url"`  //剪辑前url
	EndUrl    time.Time `gorm:"column:end_url"`    //剪辑后url
}
