package models

import "gorm.io/gorm"

// Task 任务
type Task struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);not null;comment:任务名称"`
	Description string `gorm:"type:varchar(100);not null;comment:任务描述"`
}
