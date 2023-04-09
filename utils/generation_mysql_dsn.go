package utils

func GenerationMysqlDsn(username, password, host, port, dbName string) string {
	return username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
}
