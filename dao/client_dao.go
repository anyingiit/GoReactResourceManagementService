package dao

/*以下函数已被弃用，使用新的通用函数替代

// ClientByID
func ClientByID(id uint) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	}
}

// ClientsByRecordRange
func ClientsByRecordRange(start, end int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := start - 1
		limit := end - start + 1
		return db.Offset(offset).Limit(limit)
	}
}

// find clients
func FindClients(db *gorm.DB, scopes ...ScopeFunc) ([]*models.Client, error) {
	var clients []*models.Client
	err := Find(db, &clients, scopes...)
	return clients, err
}

// first client
func FirstClient(db *gorm.DB, scopes ...ScopeFunc) (*models.Client, error) {
	client := models.Client{}
	err := First(db, &client, scopes...)
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

*/
