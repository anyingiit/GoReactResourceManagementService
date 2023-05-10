package dao

import (
	"fmt"

	"gorm.io/gorm"
)

// ScopeFunc 作用域函数的类型
type ScopeFunc func(*gorm.DB) *gorm.DB

// ByField 通过字段名和字段值生成作用域函数，限定查询结果
func ByField(field string, value interface{}) ScopeFunc {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s = ?", field), value)
	}
}

// ByID 通过 ID 生成作用域函数，限定查询结果
func ByID(id uint) ScopeFunc {
	return ByField("id", id)
}

// ByRecordRange 通过起始位置和记录数量生成作用域函数，限定查询结果
func ByRecordRange(start, end int) ScopeFunc {
	return func(db *gorm.DB) *gorm.DB {
		offset := start - 1
		limit := end - start + 1
		return db.Offset(offset).Limit(limit)
	}
}

// Find 在数据库中查找数据
func Find(db *gorm.DB, result interface{}, scopes ...ScopeFunc) error {
	query := db.Model(result)
	for _, scope := range scopes {
		query = scope(query)
	}
	return query.Find(result).Error
}

// First 在数据库中查找第一条数据
func First(db *gorm.DB, result interface{}, scopes ...ScopeFunc) error {
	query := db.Model(result)
	for _, scope := range scopes {
		query = scope(query)
	}
	return query.First(result).Error
}

// Count 统计查询结果数量
func Count(db *gorm.DB, model interface{}, count *int64, scopes ...ScopeFunc) error {
	query := db.Model(model)
	for _, scope := range scopes {
		query = scope(query)
	}
	return query.Count(count).Error
}

// Create 在数据库中创建记录
func Create(db *gorm.DB, value interface{}) error {
	return db.Create(value).Error
}

// Delete 在数据库中删除记录
func Delete(db *gorm.DB, value interface{}, scopes ...ScopeFunc) error {
	query := db.Model(value)
	for _, scope := range scopes {
		query = scope(query)
	}
	return query.Delete(value).Error
}

// Update 在数据库中更新记录
func Update(db *gorm.DB, value interface{}, scopes ...ScopeFunc) error {
	query := db.Model(value)
	for _, scope := range scopes {
		query = scope(query)
	}
	return query.Updates(value).Error
}

// Exists 检查当前条件下是否存在匹配的记录
func Exists(db *gorm.DB, model interface{}, scopes ...ScopeFunc) (bool, error) {
	var count int64
	err := Count(db, model, &count, scopes...)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
