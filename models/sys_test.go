package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSysTableName(t *testing.T) {
	sys := &Sys{}
	tableName := sys.TableName()

	assert.Equal(t, "sys", tableName, "Table name should be sys")
}
