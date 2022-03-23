package services

import (
	"errors"
	"library/internal/app/database"
	"library/internal/app/models"

	"github.com/golang/glog"
)

type LibraryService struct{}

func NewLibraryService() *LibraryService {
	return &LibraryService{}
}

func (l *LibraryService) FetchBooks() ([]models.Books, error) {
	var booksData []models.Books

	// fetch active database connection
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

func (l *LibraryService) AddBook(bookData models.Books) error {
	// fetch active database connection
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
