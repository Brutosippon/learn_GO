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

// GET BOOK
// "c *gin" is a pointer to the context, or the current state of the application (request, response, etc.)
// c *gin.Context is a gin context, which holds information about the request and response.
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// POST CREATE BOOK
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

// POST SEARCH BOOK
func getBookByID(c *gin.Context) {
	id := c.Param("id") // Param returns the value of the URL parameter
	for _, a := range books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"}) // 404 not found
}

// PUT  UPDATE THE BOOK
func updateBook(c *gin.Context) {
	var updatedBook book
	// Call BindJSON to bind the received JSON to updatedBook
	if err := c.BindJSON(&updatedBook); err != nil {
		return
	}
	id := c.Param("id") // Param returns the value of the URL parameter
	for i, a := range books {
		if a.ID == id {
			books[i] = updatedBook
			c.IndentedJSON(http.StatusOK, updatedBook)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"}) // 404 not found
}

// DELETE BOOK
func deleteBook(c *gin.Context) {
	id := c.Param("id") // Param returns the value of the URL parameter
	for i, a := range books {
		if a.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "book deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"}) // 404 not found
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("/books/:id", getBookByID)
	router.PUT("/books/:id", updateBook)
	router.DELETE("/books/:id", deleteBook)
	router.Run("localhost:8081")
}
