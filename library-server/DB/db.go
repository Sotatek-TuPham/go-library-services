package db

import (
	"library-server/model"
	"os"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GenerateInitialData() {
	var bookCount, categoryCount int64

	DB.Model(&model.Book{}).Count(&bookCount)
	DB.Model(&model.Category{}).Count(&categoryCount)

	if bookCount == 0 && categoryCount == 0 {
		categories := []string{"Fiction", "Non-Fiction", "Science", "History", "Technology"}
		for _, cat := range categories {
			DB.Create(&model.Category{Name: cat})
		}

		books := []model.Book{
			{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", CategoryID: 1, Location: "A1", Status: model.BookStatusAvailable},
			{Title: "To Kill a Mockingbird", Author: "Harper Lee", CategoryID: 1, Location: "A2", Status: model.BookStatusAvailable},
			{Title: "1984", Author: "George Orwell", CategoryID: 1, Location: "A3", Status: model.BookStatusAvailable},
			{Title: "The Catcher in the Rye", Author: "J.D. Salinger", CategoryID: 1, Location: "A4", Status: model.BookStatusAvailable},
			{Title: "Sapiens", Author: "Yuval Noah Harari", CategoryID: 2, Location: "B1", Status: model.BookStatusAvailable},
			{Title: "A Brief History of Time", Author: "Stephen Hawking", CategoryID: 3, Location: "C1", Status: model.BookStatusAvailable},
			{Title: "The Guns of August", Author: "Barbara Tuchman", CategoryID: 4, Location: "D1", Status: model.BookStatusAvailable},
			{Title: "Clean Code", Author: "Robert C. Martin", CategoryID: 5, Location: "E1", Status: model.BookStatusAvailable},
			{Title: "The Pragmatic Programmer", Author: "Andrew Hunt", CategoryID: 5, Location: "E2", Status: model.BookStatusAvailable},
			{Title: "Design Patterns", Author: "Erich Gamma", CategoryID: 5, Location: "E3", Status: model.BookStatusAvailable},
		}

		for _, book := range books {
			DB.Create(&book)
		}
	}
}

// Call this function after establishing the database connection
func InitializeDatabase() {
	Connect()
	GenerateInitialData()
}

func Connect() {
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var adminCount int64
	db.Raw("SELECT COUNT(*) FROM admins").Scan(&adminCount)
	if adminCount == 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		db.Exec("INSERT INTO admins (username, password) VALUES (?, ?)", "admin", string(hashedPassword))
	}

	db.AutoMigrate(&model.Admin{}, &model.Book{}, &model.Category{})
	DB = db
}
