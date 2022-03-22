package models

type Book struct {
	Id     int    `gorm:"primary_key";"AUTO_INCREMENT";"column:id"`
	Name   string `gorm:"column:name"`
	Author string `gorm:"column:author"`
}
