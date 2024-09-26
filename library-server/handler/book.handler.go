package handler

import (
	"math"
	"net/http"
	"strconv"

	"library-server/model"
	"library-server/service"

	"github.com/gin-gonic/gin"
)


// CreateBook godoc
// @Summary Create a new book
// @Description Create a new book with the input payload
// @Tags books
// @Accept json
// @Produce json
// @Param book body model.Book true "Create book"
// @Param Authorization header string true "Bearer token" default(Bearer <Add access token here>)
// @Success 201 {object} model.Book
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Security BearerAuth
// @Router /books [post]
func CreateBook(c *gin.Context) {
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.CreateBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}
	c.JSON(http.StatusCreated, book)
}

// GetBookByID godoc
// @Summary Get a book by ID
// @Description Get a single book by its ID
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Param Authorization header string true "Bearer token" default(Bearer <Add access token here>)
// @Success 200 {object} model.Book
// @Failure 404 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Security BearerAuth
// @Router /books/{id} [get]
func GetBookByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	book, err := service.GetBookByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// GetAllBooks godoc
// @Summary Get all books
// @Description Get a list of all books with pagination and optional filters
// @Tags books
// @Produce json
// @Param Authorization header string true "Bearer token" default(Bearer <Add access token here>)
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Number of items per page" default(10)
// @Param author query string false "Filter by author"
// @Param category query string false "Filter by category name"
// @Param title query string false "Filter by book title"
// @Success 200 {object} map[string]interface{} "Contains 'books' array, 'total' count, and 'pages' count"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Security BearerAuth
// @Router /books [get]
func GetAllBooks(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	author := c.Query("author")
	category := c.Query("category")
	title := c.Query("title")

	if page < 1 || pageSize < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page or pageSize"})
		return
	}

	books, totalCount, err := service.GetAllBooks(page, pageSize, author, category, title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(pageSize)))

	c.JSON(http.StatusOK, gin.H{
		"books": books,
		"total": totalCount,
		"pages": totalPages,
	})
}

// UpdateBook godoc
// @Summary Update a book
// @Description Update a book with the input payload
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body model.Book true "Update book"
// @Param Authorization header string true "Bearer token" default(Bearer <Add access token here>)
// @Success 200 {object} model.Book
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Security BearerAuth
// @Router /books/{id} [put]
func UpdateBook(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book.ID = uint(id)
	if err := service.UpdateBook(&book); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// DeleteBook godoc
// @Summary Delete a book
// @Description Delete a book by its ID
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Param Authorization header string true "Bearer token" default(Bearer <Add access token here>)
// @Success 204 "No Content"
// @Failure 404 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Security BearerAuth
// @Router /books/{id} [delete]
func DeleteBook(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := service.DeleteBook(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.Status(http.StatusNoContent)
}

// GetBooksByCategory godoc
// @Summary Get books by category
// @Description Get a list of books in a specific category
// @Tags books
// @Produce json
// @Param categoryID path int true "Category ID"
// @Param Authorization header string true "Bearer token" default(Bearer <Add access token here>)
// @Success 200 {array} model.Book
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Security BearerAuth
// @Router /books/category/{categoryID} [get]
func GetBooksByCategory(c *gin.Context) {
	categoryID, err := strconv.ParseUint(c.Param("categoryID"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}
	books, err := service.GetBooksByCategory(uint(categoryID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}
	c.JSON(http.StatusOK, books)
}

// UpdateBookStatus godoc
// @Summary Update book status
// @Description Update the status of a book
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param status body model.BookStatus true "New book status"
// @Param Authorization header string true "Bearer token" default(Bearer <Add access token here>)
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Security BearerAuth
// @Router /books/{id}/status [patch]
func UpdateBookStatus(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var status model.BookStatus
	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.UpdateBookStatus(uint(id), status); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book status updated successfully"})
}
