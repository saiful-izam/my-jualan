package env

import "os"

type Environment struct {
	User       string
	Password   string
	Host       string
	DbName     string
	ServerPort string
}

func GetEnvironment() *Environment {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("PORT")

	return &Environment{
		User:       user,
		Password:   pass,
		Host:       host,
		DbName:     dbName,
		ServerPort: port,
	}
}
