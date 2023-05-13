package models

import "gorm.io/gorm"

// TaskParamValue 任务参数值
type TaskParamValue struct {
	gorm.Model
	TaskParamTypeID uint          `gorm:"type:int;not null;comment:任务参数类型ID"`
	TaskParamType   TaskParamType `gorm:"foreignKey:TaskParamTypeID"`
	Value           string        `gorm:"type:varchar(100);not null;comment:任务参数值"`
}
