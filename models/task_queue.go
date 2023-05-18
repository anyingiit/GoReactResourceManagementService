package models

import "gorm.io/gorm"

type TaskQueue struct {
	gorm.Model
	Sequence  int     `gorm:"type:int;not null;comment:队列中的序列编号"`
	TaskID    uint    `gorm:"type:int;not null;comment:任务ID"`
	Task      Task    `gorm:"foreignKey:TaskID"`
	ClientID  uint    `gorm:"type:int;not null;comment:客户端ID"`
	Client    Client  `gorm:"foreignKey:ClientID"`
	ServiceID uint    `gorm:"type:int;not null;comment:服务ID"`
	Service   Service `gorm:"foreignKey:ServiceID"`
}
