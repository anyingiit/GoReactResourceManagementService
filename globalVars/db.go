package globalVars

import (
	"gorm.io/gorm"
)

var Db *db

type db struct {
	*GlobalVars
}

// new db
func newDb() *db {
	return &db{
		GlobalVars: newGlobalVars("Db"),
	}
}

func (d *db) Set(dbObj *gorm.DB) error {
	return d.GlobalVars.Set(dbObj)
}

func (d *db) Get() (*gorm.DB, error) {
	val, err := d.GlobalVars.Get()
	if err != nil {
		return nil, err
	}
	return val.(*gorm.DB), nil
}
