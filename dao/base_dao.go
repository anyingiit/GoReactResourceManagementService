package dao

import (
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ScopeFunc 作用域函数的类型
type ScopeFunc func(*gorm.DB) *gorm.DB

// ByField 通过字段名和字段值生成作用域函数，限定查询结果
func ByField(field string, value interface{}) ScopeFunc {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(clause.Eq{Column: clause.Column{Name: field}, Value: value})
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

// ByRecordRangeWithPage 通过页码和每页记录数量生成作用域函数，限定查询结果
func ByRecordRangeWithPage(page, perPage int) ScopeFunc {
	return ByRecordRange((page-1)*perPage+1, page*perPage)
}

// Preload 函数实现了预载入关联数据的功能，并返回一个通用的作用域（scope）函数，可以在查询时使用。
func Preload(scope string) ScopeFunc {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload(scope)
	}
}

// FieldIn 通过字段名和字段值列表生成作用域函数，限定查询结果
// 例如：FieldIn("id", []uint{1, 2, 3}) 生成的作用域函数为：db.Where("id IN ?", []uint{1, 2, 3})
func FieldIn(field string, values []interface{}) ScopeFunc {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(clause.IN{Column: clause.Column{Name: field}, Values: values})
	}
}

// IDsIn 通过 ID 列表生成作用域函数，限定查询结果
func IDsIn(ids []uint) ScopeFunc {

	interfaceValues := make([]interface{}, len(ids))
	for i, val := range ids {
		interfaceValues[i] = val
	}

	return FieldIn("id", interfaceValues)
}

// Joins
func Joins(query string, args ...interface{}) ScopeFunc {
	return func(db *gorm.DB) *gorm.DB {
		return db.Joins(query, args...)
	}
}

// Order
func Order(sortField, sortOrder string) ScopeFunc {
	return func(db *gorm.DB) *gorm.DB {
		// return db.Order(gorm.Expr("? ?", sortField, sortOrder))
		return db.Order(clause.OrderByColumn{Column: clause.Column{Name: sortField}, Desc: strings.ToUpper(sortOrder) == "DESC"})
	}
}

// ByFields
func ByFields(fields map[string]interface{}) ScopeFunc {
	return func(db *gorm.DB) *gorm.DB {
		for k, v := range fields {
			db = ByField(k, v)(db)
		}
		return db
	}
}

// CountScope
func CountScope(count *int64) ScopeFunc {
	return func(db *gorm.DB) *gorm.DB {
		return db.Count(count)
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
