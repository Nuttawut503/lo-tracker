package config

import "os"

type DatabaseConfig struct {
	Path     string
	Password string
}

func GetDBConfig() DatabaseConfig {
	return DatabaseConfig{
		os.Getenv("dbpath"),
		os.Getenv("dbpassword"),
	}
}
