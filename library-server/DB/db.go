package db

import (
	"library-server/model"
	"os"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

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
