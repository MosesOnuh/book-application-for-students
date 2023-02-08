package main

import (
	"log"
	"net/http"
	"time"
	"fmt"
	"github.com/gin-gonic/gin"
)

type book struct{
	Id string `json:"id"`
	Name string `json:"name"`
	Pages int 	`json:"pages"`
	Created_at int64 `json:"created_at"`
	Updated_at int64 `json:"updated_at"`
}

//the collections slice acts as a database in this application
var collections = []book{
	{
		Id: "1",
		Name: "Oliver twist",
		Pages: 45,
		Created_at: 1675596874,
		Updated_at: 0,
	},
}


func main(){
	router := gin.Default()
	log.Println("Application started on port 8080")

	//get all the books in collections(database)
	router.GET("/allBooks", getBooks)

	//create a book and store it in collections(database)
	router.POST("/createBook", createBook)

	//get a book that is stored in the collections(database)
	router.GET("/getBookById/:id", getBook)

	//update the details of a book in the collections(database)
	router.PUT("/updateBookById/:id", updateBookDetails)


	router.Run("localhost:8080")
}

func getBooks(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully retrieved all books",
		"books": collections,
	})
}

func createBook(c *gin.Context){
	var newbook book

	err := c.Bind(&newbook)
	if err != nil {
		c.JSON(400, gin.H{
			"error":"request format is not correct",
		})
		return
	}

	newbook.Created_at = time.Now().Unix()

	//storing the newbook to the collections
	collections = append(collections, newbook)

	c.JSON(201, gin.H{
		"message": "Successfully created a new Book",
	})

}

func getBook(c *gin.Context){
	//get Id of book from the endpoint
	id := c.Param("id")

	// create a variable to store the current book from the collections
	var currentBook book

	//loop through the collections
	//check if the id from the endpoint is equal to the id of a book in the collections
	//if the id of the endpoint matches the Id of the book, then assign the book to the variable currentBook
	for _, currentBookInLoop := range collections {
		if id == currentBookInLoop.Id {
			currentBook = currentBookInLoop
		}
	}

	//send the response message and book to the front end
	c.JSON(200, gin.H{
		"message": "Successfully got the requested book",
		"data": currentBook,
	})
}

//this handler updates the details of a book that is stored in the collections(database)
func updateBookDetails(c *gin.Context){
	id := c.Param("id")

	var requestBook book 

	err := c.Bind(&requestBook)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message":"request format is not correct",
		})
		return
	}

	for i, value := range collections{
		if id == value.Id {
			if requestBook.Name != "" {
				collections[i].Name = requestBook.Name
			}

			if requestBook.Pages != 0 {
				collections[i].Pages = requestBook.Pages
			}

			collections[i].Updated_at = time.Now().Unix()

			c.JSON(http.StatusOK, gin.H{
				"message": "Successfully updated book",
			})
			return
		}

		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("book with the ID %s is not stored in the database", id),
		})
	}

}

