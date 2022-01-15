package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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

	//user := os.Getenv("APP_DB_USERNAME")
	//password := os.Getenv("APP_DB_PASSWORD")
	//db := os.Getenv("APP_DB_NAME")

	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "rumi",
		DBName:   "go_book_store",
		Password: "Abcdefgh.1",
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	//"root:root@/simplerest?charset=utf8&parseTime=True&loc=Local"
	return fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
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
