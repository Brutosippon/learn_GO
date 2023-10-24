package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"errors"
)

type book struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Quality string `json:"quality"`
}

var books = []book{
	{ID: "1", Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", Quality: "Good"},
	{ID: "2", Title: "The Hobbit", Author: "J.R.R. Tolkien", Quality: "Good"},
	{ID: "3", Title: "The Lord of the Rings", Author: "J.R.R. Tolkien", Quality: "Good"},
	{ID: "4", Title: "The Silmarillion", Author: "J.R.R. Tolkien", Quality: "Good"},
	{ID: "5", Title: "The C Programming Language", Author: "Brian W. Kernighan, Dennis M. Ritchie", Quality: "Good"},
}

// "c *gin" is a pointer to the context, or the current state of the application (request, response, etc.)
// c *gin.Context is a gin context, which holds information about the request and response.
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	/*router.GET("/books/:id", getBookByID)
	router.POST("/books", addBook)
	router.PUT("/books/:id", updateBook)
	router.DELETE("/books/:id", deleteBook)*/
	router.Run("localhost:8081")
}
