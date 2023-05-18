package db

import (
	"fmt"

	"github.com/anyingiit/GoReactResourceManagement/dao"
	"github.com/anyingiit/GoReactResourceManagement/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SetupData 返回的信息结构体
type SetupDataResult struct {
	User models.User
}

// SetupData 用于初始化项目
// 使用时注意使用tx，以便于回滚
func SetupData(db *gorm.DB, superAdminPassword string) (*SetupDataResult, error) {
	// 1. 创建Sys表中的唯一一条记录
	// 2. 创建Role表中的默认数据
	// 3. 创建一个默认的SuperAdmin用户
	// 4. 创建默认支持的任务

	// 创建Sys表中的唯一一条记录
	sys := models.Sys{}
	if db.Model(&models.Sys{}).Create(&sys).Error != nil {
		return nil, fmt.Errorf("failed to create sys")
	}

	// 创建Role表中的默认数据
	roles := []models.Role{}
	roles = append(roles, models.Role{Name: "SuperAdmin", Description: "SuperAdmin can do anything"})
	roles = append(roles, models.Role{Name: "Admin", Description: "Admin can do anything except manage admin and SuperAdmin"})
	roles = append(roles, models.Role{Name: "User"})
	if db.Model(&models.Role{}).Create(&roles).Error != nil {
		return nil, fmt.Errorf("failed to create roles")
	}

	// 创建一个默认的SuperAdmin用户
	// 1. 获取SuperAdmin角色
	// 2. 创建SuperAdmin模型实体
	// 3. 创建SuperAdmin用户
	superAdmin := &models.Role{}
	if db.Model(&models.Role{}).Where("name = ?", "SuperAdmin").First(&superAdmin).Error != nil {
		return nil, fmt.Errorf("failed to get SuperAdmin role")
	}

	username := uuid.NewString()
	password := superAdminPassword
	newSuperAdminUser, err := models.NewUser(username,
		password,
		"DefaultSuperAdmin",
		18,
		*superAdmin)
	if err != nil {
		return nil, fmt.Errorf("failed to create SuperAdmin user, %v", err)
	}

	// // 需要强制修改密码
	// newSuperAdminUser.MustChangePassword = true

	if db.Model(&models.User{}).Create(&newSuperAdminUser).Error != nil {
		return nil, fmt.Errorf("failed to create SuperAdmin user, %v", err)
	}

	// 打印默认SuperAdmin用户的用户名和密码
	fmt.Println("SuperAdmin username: ", username)
	// fmt.Println("SuperAdmin password: ", password)

	// 创建默认支持的任务
	tasks := []models.Task{}
	tasks = append(tasks, models.Task{Name: "context_task_queue_pull", Description: "context_task_queue_pull"})
	tasks = append(tasks, models.Task{Name: "context_task_queue_clear", Description: "context_task_queue_clear"})
	tasks = append(tasks, models.Task{Name: "context_task_queue_reload", Description: "context_task_queue_reload"})
	tasks = append(tasks, models.Task{Name: "context_init", Description: "context_init"})
	tasks = append(tasks, models.Task{Name: "context_client_config_is_validate", Description: "context_client_config_is_validate"})
	tasks = append(tasks, models.Task{Name: "context_invate_info_load_with_command_line_param", Description: "context_invate_info_load_with_command_line_param"})
	tasks = append(tasks, models.Task{Name: "client_regist_with_invate_info", Description: "client_regist_with_invate_info"})
	tasks = append(tasks, models.Task{Name: "update_config_file_with_context", Description: "update_config_file_with_context"})
	tasks = append(tasks, models.Task{Name: "ping", Description: "Ping a host"})
	tasks = append(tasks, models.Task{Name: "wait_five_minute", Description: "wait_five_minute"})
	tasks = append(tasks, models.Task{Name: "contextl_task_queue_result_push", Description: "contextl_task_queue_result_push"})
	if err := dao.Create(db, &tasks); err != nil {
		return nil, fmt.Errorf("failed to create tasks")
	}

	return &SetupDataResult{
		User: *newSuperAdminUser,
	}, nil
}
