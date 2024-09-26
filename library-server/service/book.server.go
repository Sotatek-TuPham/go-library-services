package service

import (
	"errors"
	db "library-server/DB"
	"library-server/model"
)

// CreateBook creates a new book in the database
func CreateBook(book *model.Book) error {
	result := db.DB.Create(book)
	return result.Error
}

// GetBookByID retrieves a book by its ID
func GetBookByID(id uint) (*model.Book, error) {
	var book model.Book
	result := db.DB.Preload("Category").First(&book, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

// GetAllBooks retrieves books from the database with pagination and optional filters
func GetAllBooks(page, pageSize int, author, category, title string) ([]model.Book, int64, error) {
	var books []model.Book
	var totalCount int64

	query := db.DB.Model(&model.Book{}).Preload("Category")

	if author != "" {
		query = query.Where("author LIKE ?", "%"+author+"%")
	}
	if category != "" {
		query = query.Joins("JOIN categories ON books.category_id = categories.id").
			Where("categories.name LIKE ?", "%"+category+"%")
	}
	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}

	// Count total matching records
	if err := query.Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	offset := (page - 1) * pageSize
	result := query.Offset(offset).Limit(pageSize).Find(&books)

	return books, totalCount, result.Error
}

// UpdateBook updates an existing book in the database
func UpdateBook(book *model.Book) error {
	result := db.DB.Save(book)
	return result.Error
}

// DeleteBook deletes a book from the database
func DeleteBook(id uint) error {
	result := db.DB.Delete(&model.Book{}, id)
	if result.RowsAffected == 0 {
		return errors.New("book not found")
	}
	return result.Error
}

// GetBooksByCategory retrieves all books in a specific category
func GetBooksByCategory(categoryID uint) ([]model.Book, error) {
	var books []model.Book
	result := db.DB.Preload("Category").Where("category_id = ?", categoryID).Find(&books)
	return books, result.Error
}

// UpdateBookStatus updates the status of a book
func UpdateBookStatus(id uint, status model.BookStatus) error {
	result := db.DB.Model(&model.Book{}).Where("id = ?", id).Update("status", status)
	if result.RowsAffected == 0 {
		return errors.New("book not found")
	}
	return result.Error
}
