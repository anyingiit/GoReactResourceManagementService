package models

import "gorm.io/gorm"

type TaskQueue struct {
	gorm.Model
	Sequence         int            `gorm:"type:int;not null;comment:队列中的序列编号"`
	TaskID           uint           `gorm:"type:int;not null;comment:任务ID"`
	Task             Task           `gorm:"foreignKey:TaskID"`
	ClientID         uint           `gorm:"type:int;not null;comment:客户端ID"`
	Client           Client         `gorm:"foreignKey:ClientID"`
	TaskParamValueID uint           `gorm:"type:int;not null;comment:任务参数值ID"`
	TaskParamValue   TaskParamValue `gorm:"foreignKey:TaskParamValueID"`
}
