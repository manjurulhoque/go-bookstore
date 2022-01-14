package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

var (
	db *gorm.DB
)

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {

	user := os.Getenv("APP_DB_USERNAME")
	password := os.Getenv("APP_DB_PASSWORD")
	db := os.Getenv("APP_DB_NAME")

	dbConfig := DBConfig{
		Host:     "0.0.0.0",
		Port:     3306,
		User:     user,
		DBName:   db,
		Password: password,
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func Connect() {
	d, err := gorm.Open("mysql", DbURL(BuildDBConfig()))

	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}