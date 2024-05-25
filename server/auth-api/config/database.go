package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabase(env *Environment) (*gorm.DB, error) {
	connString := "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	dsn := fmt.Sprintf(connString, env.User, env.Password, env.Host, env.DbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}
