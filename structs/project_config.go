package structs

import "time"

type ProjectConfig struct {
	Database     DatabaseConfig     `yaml:"database"`
	Environments EnvironmentsConfig `yaml:"environments"`
	Server       ServerConfig       `yaml:"server"`
	Token        TokenConfig        `yaml:"token"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type EnvironmentsConfig struct {
	Development EnvironmentConfig `yaml:"development"`
	Production  EnvironmentConfig `yaml:"production"`
}

type EnvironmentConfig struct {
	Debug bool `yaml:"debug"`
}

type ServerConfig struct {
	Ip   string `yaml:"ip"`
	Port int    `yaml:"port"`
}

type TokenConfig struct {
	ExpiredTime time.Duration `yaml:"expired_time"`
	SigningKey  string        `yaml:"signing_key"`
}
