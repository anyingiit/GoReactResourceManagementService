package client

import "gorm.io/gorm"

type TaskController struct {
	*BaseController
}

// NewTaskController
func NewTaskController(db *gorm.DB) *TaskController {
	return &TaskController{
		BaseController: NewBaseController(db),
	}
}

// Get: 获取任务
