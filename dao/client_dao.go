package dao

import (
	"github.com/anyingiit/GoReactResourceManagement/models"
	"gorm.io/gorm"
)

// ClientByID
func ClientByID(id uint) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	}
}

// find clients
func FindClients(db *gorm.DB, scopes ...func(*gorm.DB) *gorm.DB) ([]*models.Client, error) {
	var clients []*models.Client
	query := db.Model(&models.Client{})
	for _, scope := range scopes {
		query = scope(query)
	}
	err := query.Find(&clients).Error
	return clients, err
}

// find client
func FirstClient(db *gorm.DB, scopes ...func(*gorm.DB) *gorm.DB) (*models.Client, error) {
	client := models.Client{}
	query := db.Model(&models.Client{})
	for _, scope := range scopes {
		query = scope(query)
	}
	err := query.First(&client).Error
	return &client, err
}

// Update Client by id
func UpdateClientById(db *gorm.DB, id uint, updateFields *models.Client) error {
	return db.Model(&models.Client{}).Where("id = ?", id).Updates(updateFields).Error
}

// Save Client
func SaveClient(db *gorm.DB, client *models.Client) error {
	return db.Save(client).Error
}

// Delete Client by id
func DeleteClientById(db *gorm.DB, id uint) error {
	return db.Where("id = ?", id).Delete(&models.Client{}).Error
}

// Create Client
func CreateClient(db *gorm.DB, client *models.Client) error {
	return db.Create(client).Error
}
