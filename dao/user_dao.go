package dao

import (
	"gorm.io/gorm"
)

func UserByID(id uint) func(*gorm.DB) *gorm.DB {
	return ByID(id)
}

/* 使用dao通用函数替换了如下专有函数
// []*models.User 底层存储的是一个指针数组，指针指向的是models.User的实例
func FindUsers(db *gorm.DB, scopes ...func(*gorm.DB) *gorm.DB) ([]*models.User, error) {
	var users []*models.User
	query := db.Model(&models.User{}).Preload("Role")
	for _, scope := range scopes {
		query = scope(query)
	}
	err := query.Find(&users).Error // Find方法是通过切片的指针来接收结果的
	return users, err
}

// 返回的是一个指针，指向的是models.User的实例
func FirstUser(db *gorm.DB, scopes ...func(*gorm.DB) *gorm.DB) (*models.User, error) {
	user := models.User{}
	query := db.Model(&models.User{}).Preload("Role")
	for _, scope := range scopes {
		query = scope(query)
	}
	err := query.First(&user).Error // First需要一个模型实例的指针
	return &user, err
}
*/
