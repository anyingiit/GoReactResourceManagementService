package globalVars

import (
	"testing"

	"gorm.io/gorm"
)

func TestDb(t *testing.T) {
	dbObj := &gorm.DB{}
	db := newDb()

	// Test Set and Get
	err := db.Set(dbObj)
	if err != nil {
		t.Errorf("Expected error to be nil, but got %v", err)
	}

	val, err := db.Get()
	if err != nil {
		t.Errorf("Expected error to be nil, but got %v", err)
	}

	if val != dbObj {
		t.Errorf("Expected dbObj to be %v, but got %v", dbObj, val)
	}
}
