package dao

/* 使用dao通用函数替换了如下专有函数
// FirstSysRecord
func FirstSysRecord(db *gorm.DB) (*models.Sys, error) {
	sys := models.Sys{}
	err := db.First(&sys).Error
	return &sys, err
}
*/
