package services

import (
	"errors"
	"library/internal/app/database"
	"library/internal/app/models"

	"github.com/golang/glog"
)

func FetchBooks() ([]models.Book, error) {
	var booksData []models.Book

	dbConnection := database.GetDB()
	if dbConnection == nil {
		glog.Error("DB Connection Is Null While Fetching Books")
		return booksData, errors.New("No DB Connection Found...")
	}

	err := dbConnection.Raw(
		`SELECT *
			FROM books b`).Scan(&booksData).Error
	if len(booksData) > 0 {
		return booksData, nil
	}
	return booksData, err
}

func AddBook(bookData models.Book) error {
	dbConnection := database.GetDB()
	if dbConnection == nil {
		glog.Error("DB Connection Is Null While Fetching Books")
		return errors.New("No DB Connection Found...")
	}

	err := dbConnection.Create(&bookData).Error
	if err != nil {
		glog.Error("Create new record in book table failed...")
		return errors.New("Add Book failed...")
	}
	return nil
}
