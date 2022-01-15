package models

import (
	"github.com/jinzhu/gorm"
	"github.com/manjurulhoque/go-bookstore/pkg/config"
	"time"
)

var db *gorm.DB

type Book struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `gorm:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"deleted_at" sql:"index" json:"deleted_at"`
	Name        string     `gorm:"name" json:"name" validate:"required,min=4"`
	Author      string     `gorm:"author" json:"author" validate:"required"`
	Publication string     `gorm:"publication" json:"publication" validate:"required"`
}

//type GetBook struct {
//	Id          int64  `gorm:"" json:"id"`
//	Name        string `gorm:"" json:"name"`
//	Author      string `json:"author"`
//	Publication string `json:"publication"`
//	CreatedAt   string `json:"created_at"`
//	UpdatedAt   string `json:"updated_at"`
//	DeletedAt   string `json:"deleted_at"`
//}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Unscoped().Where("ID=?", Id).Delete(book)
	return book
}
