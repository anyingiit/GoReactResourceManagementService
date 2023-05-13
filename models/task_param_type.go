package models

import "gorm.io/gorm"

// TaskParamType 任务参数类型
type TaskParamType struct {
	gorm.Model
	Type        string `gorm:"type:varchar(100);not null;comment:参数类型"`
	DefultValue string `gorm:"type:varchar(100);not null;comment:默认值"`
}
