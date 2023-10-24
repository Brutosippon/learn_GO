package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"errors"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", Quantity: 2},
	{ID: "2", Title: "The Hobbit", Author: "J.R.R. Tolkien", Quantity: 2},
	{ID: "3", Title: "The Lord of the Rings", Author: "J.R.R. Tolkien", Quantity: 2},
}

// "c *gin" is a pointer to the context, or the current state of the application (request, response, etc.)
// c *gin.Context is a gin context, which holds information about the request and response.
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook book
	// Call BindJSON to bind the received JSON to newBook
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	// Add the new book to the slice.
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

/*
Test POST with PowerShell:
Invoke-WebRequest -Uri 'http://localhost:8081/books' -Method POST -Headers @{"Content-Type"="application/json"} -InFile 'body.json'
*/

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	/*router.GET("/books/:id", getBookByID)
	router.POST("/books", addBook)
	router.PUT("/books/:id", updateBook)
	router.DELETE("/books/:id", deleteBook)*/
	router.Run("localhost:8081")
}
