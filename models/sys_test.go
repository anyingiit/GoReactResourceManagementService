package models_test

import (
	"testing"

	"github.com/anyingiit/GoReactResourceManagement/models"
	"github.com/stretchr/testify/assert"
)

func TestSysTableName(t *testing.T) {
	sys := &models.Sys{}
	tableName := sys.TableName()

	assert.Equal(t, "sys", tableName, "Table name should be sys")
}
