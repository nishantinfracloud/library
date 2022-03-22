package controller

import (
	"library/internal/app/models"
	"library/internal/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

var libraryService = services.NewLibraryService()

// FetchBooks: fetches data of all available books
func FetchBooks() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		booksData, err := libraryService.FetchBooks()
		if err != nil {
			glog.Error("FetchBooks Failed...")
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusOK, booksData)
	}
	return fn
}

// AddBook: adds a new book
func AddBook() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var requestData models.Books
		err := c.Bind(&requestData)
		if err != nil {
			glog.Error("Binding RequestData Failed...")
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err = libraryService.AddBook(requestData)
		if err != nil {
			glog.Error("Add Book Failed...")
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.String(http.StatusOK, "Added book successfully...")
	}
	return fn
}
