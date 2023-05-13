package models

import "gorm.io/gorm"

// ServiceStatus 服务状态
type ServiceStatus struct {
	gorm.Model
	ServiceID    uint `gorm:"type:int;not null;comment:服务ID"`
	Service      Service
	TaskResultID uint `gorm:"type:int;not null;comment:任务结果ID"`
	TaskResult   TaskQueueResult
}
