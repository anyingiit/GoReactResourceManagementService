package dao

import (
	"github.com/anyingiit/GoReactResourceManagement/structs"
	"gorm.io/gorm"
)

type StanderedQueryMethod string

var (
	StanderedQueryMethodFind  StanderedQueryMethod = "find"
	StanderedQueryMethodFirst StanderedQueryMethod = "first"
)

func StanderedQueryBase(db *gorm.DB, result interface{}, standeredQueryData *structs.StanderedQueryParsed, Finishers func(db *gorm.DB) error, scopes ...ScopeFunc) error {
	query := db.Model(result)

	scopeFuncs := []ScopeFunc{
		ByFields(standeredQueryData.FilterJson),
		ByRecordRangeWithPage(standeredQueryData.PaginationPage, standeredQueryData.PaginationPerPage),
		Order(standeredQueryData.SortField, standeredQueryData.SortOrder),
	}
	scopeFuncs = append(scopeFuncs, scopes...)

	for _, scope := range scopeFuncs {
		query = scope(query)
	}

	if err := Finishers(query); err != nil {
		return err
	}

	return nil
}

// func StanderedQueryFind(db *gorm.DB, result interface{}, count *int64, standeredQueryData *structs.StanderedQueryParsed, scopes ...ScopeFunc) error {
// 	query := db.Model(result)

// 	scopeFuncs := []ScopeFunc{
// 		ByFields(standeredQueryData.FilterJson),
// 		ByRecordRangeWithPage(standeredQueryData.PaginationPage, standeredQueryData.PaginationPerPage),
// 		CountScope(count),
// 		Order(standeredQueryData.SortField, standeredQueryData.SortOrder),
// 	}
// 	scopeFuncs = append(scopeFuncs, scopes...)

// 	for _, scope := range scopeFuncs {
// 		query = scope(query)
// 	}

// 	if err := query.Find(result).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }

func StanderedQueryFind(db *gorm.DB, result interface{}, count *int64, standeredQueryData *structs.StanderedQueryParsed, scopes ...ScopeFunc) error {
	scopes = append(scopes, CountScope(count))

	return StanderedQueryBase(db, result, standeredQueryData, func(db *gorm.DB) error {
		return db.Find(result).Error
	}, scopes...)
}

func StanderedQueryFirst(db *gorm.DB, result interface{}, standeredQueryData *structs.StanderedQueryParsed, scopes ...ScopeFunc) error {
	return StanderedQueryBase(db, result, standeredQueryData, func(db *gorm.DB) error {
		return db.Find(result).Error
	}, scopes...)
}
