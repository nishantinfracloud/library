package models

// this struct denotes a book with required attributes
type Books struct {
	Id     int    `gorm:"primary_key";"AUTO_INCREMENT";"column:id"`
	Name   string `gorm:"column:name"`
	Author string `gorm:"column:author"`
}
