package utils_test

import (
	"os"
	"testing"
	"time"

	"github.com/anyingiit/GoReactResourceManagement/structs"
	"github.com/anyingiit/GoReactResourceManagement/utils"
)

func TestReadConfigFile(t *testing.T) {
	// 创建一个临时配置文件
	configFile, err := os.CreateTemp("", "config-*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(configFile.Name())
	defer configFile.Close()

	// 在临时文件中写入一些 YAML 格式的配置数据
	_, err = configFile.WriteString(`
database:
  host: localhost
  port: 3306
  username: root
  password: secret
  database: mydb
environments:
  development:
    debug: true
  production:
    debug: false
server:
  ip: 127.0.0.1
  port: 8080
token:
  expired_time: 10s
  signing_key: my-secret-key
`)
	if err != nil {
		t.Fatal(err)
	}

	// 调用 ReadConfigFile 函数读取配置数据
	config, err := utils.ReadConfigFile(configFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	// 验证返回的配置数据是否正确
	expectedConfig := &structs.ProjectConfig{
		Database: structs.DatabaseConfig{
			Host:     "localhost",
			Port:     3306,
			Username: "root",
			Password: "secret",
			Database: "mydb",
		},
		Environments: structs.EnvironmentsConfig{
			Development: structs.EnvironmentConfig{
				Debug: true,
			},
			Production: structs.EnvironmentConfig{
				Debug: false,
			},
		},
		Server: structs.ServerConfig{
			Ip:   "127.0.0.1",
			Port: 8080,
		},
		Token: structs.TokenConfig{
			ExpiredTime: 10 * time.Second,
			SigningKey:  "my-secret-key",
		},
	}
	if *config != *expectedConfig {
		t.Errorf("ReadConfigFile() failed, expected %+v but got %+v", expectedConfig, config)
	}
}
