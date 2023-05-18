package models

import "gorm.io/gorm"

type TaskQueueResult struct {
	gorm.Model
	Succeed     bool      `gorm:"type:bool;not null;comment:任务是否成功"`
	Detail      string    `gorm:"type:varchar(100);not null;comment:任务执行详情"`
	TaskQueueID uint      `gorm:"type:int;not null;comment:任务队列ID"`
	TaskQueue   TaskQueue `gorm:"foreignKey:TaskQueueID"`
}
