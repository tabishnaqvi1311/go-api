package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type book struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Qty int	`json:"qty"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Qty: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Qty: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Qty: 6},
}

func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK, books)
}

func bookById(c *gin.Context){
	id := c.Param("id")
	book, err := getBook(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func getBook(id string)(*book, error){
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("404 not found");
}

func addBook(c *gin.Context){
	var newBook book

	if err:= c.BindJSON(&newBook); err != nil{
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func checkoutBook(c* gin.Context){
	id, ok := c.GetQuery("id")
	if !ok{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing query "})
		return
	}

	book, err := getBook(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	if book.Qty <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "book not available"})
		return
	}

	book.Qty -= 1
	c.IndentedJSON(http.StatusOK, book)
}
func returnBook(c* gin.Context){
	id, ok := c.GetQuery("id")
	if !ok{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing query "})
		return
	}

	book, err := getBook(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	book.Qty += 1
	c.IndentedJSON(http.StatusOK, book)
}

func main(){
	router:= gin.Default()

	router.GET("/books", getBooks)
	router.GET("/book/:id", bookById)
	router.POST("/books", addBook)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	router.Run("localhost:8181")
}