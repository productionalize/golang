package config

type Config struct {
	DBConnection *DBConnection
	Host         string
	Port         string
}

type DBConnection struct {
	Host     string
	Port     string
	Database string
}

func GetConfig() *Config {
	return &Config{
		DBConnection: &DBConnection{
			Host:     "localhost",
			Port:     "27017",
			Database: "todo-db",
		},
		Host: "localhost",
		Port: "8002",
	}
}
