package dao

import (
	"gorm.io/gorm"
)

// InvateClientByClientID
func InvateClientByClientID(id uint) func(*gorm.DB) *gorm.DB {
	return ByField("client_id", id)
}

// InvateClientByInvateCode
func InvateClientByInvateCode(invateCode string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("invate_code = ?", invateCode)
	}
}

/* 使用dao通用函数替换了如下专有函数
func CreateInvateClient(db *gorm.DB, invateClient *models.InvateClient) error {
	return db.Create(invateClient).Error
}

func FindInvateClients(db *gorm.DB, scopes ...ScopeFunc) ([]*models.InvateClient, error) {
	var invateClients []*models.InvateClient
	err := Find(db, &invateClients, scopes...)
	return invateClients, err
}

func FirstInvateClient(db *gorm.DB, scopes ...ScopeFunc) (*models.InvateClient, error) {
	invateClient := new(models.InvateClient)
	err := First(db, invateClient, scopes...)
	return invateClient, err
}

// DeleteInvateClient
func DeleteInvateClient(db *gorm.DB, scopes ...func(*gorm.DB) *gorm.DB) error {
	query := db.Model(&models.InvateClient{})
	for _, scope := range scopes {
		query = scope(query)
	}
	return query.Delete(&models.InvateClient{}).Error
}
*/
