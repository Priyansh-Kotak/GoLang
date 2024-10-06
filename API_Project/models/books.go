package models

import (
	"gorm.io/gorm"
)

type Book struct {
	ID       uint    `gorm:"primarykey"; autoIncrement" json: "id"`
	Author   *string `json: "author"`
	Title    *string `json: "title"`
	Quantity int     `json: "quantity"`
}

func MigrateBook(db *gorm.DB) error {
	err := db.AutoMigrate(&Book{})

	return err
}
