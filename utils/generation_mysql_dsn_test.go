package utils_test

import (
	"testing"

	"github.com/anyingiit/GoReactResourceManagement/utils"
)

func TestGenerationMysqlDsn(t *testing.T) {
	username := "testuser"
	password := "testpassword"
	host := "localhost"
	port := "3306"
	dbName := "testdb"

	expected := "testuser:testpassword@tcp(localhost:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	actual := utils.GenerationMysqlDsn(username, password, host, port, dbName)

	if actual != expected {
		t.Errorf("GenerationMysqlDsn() failed, expected %s but got %s", expected, actual)
	}
}
