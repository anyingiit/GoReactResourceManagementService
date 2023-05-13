package db

import (
	"fmt"

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

	return &SetupDataResult{
		User: *newSuperAdminUser,
	}, nil
}

// // check is setup data
// func IsSetupData(db *gorm.DB) (bool, error) {
// 	var count int64
// 	if err := db.Model(&models.Sys{}).Count(&count).Error; err != nil {
// 		return false, fmt.Errorf("failed to get sys count, %v", err)
// 	}
// 	return count > 0, nil
// }
