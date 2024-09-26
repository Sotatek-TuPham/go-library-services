package model

type BookStatus string

const (
	BookStatusAvailable BookStatus = "available"
	BookStatusPlaced    BookStatus = "placed"
	BookStatusTaken     BookStatus = "taken"
)

type Book struct {
	ID         uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Title      string     `gorm:"not null" json:"title"`
	Author     string     `gorm:"not null" json:"author"`
	CategoryID uint       `gorm:"not null" json:"category_id"`
	Category   Category   `gorm:"foreignKey:CategoryID" json:"category"`
	Location   string     `gorm:"not null" json:"location"`
	Status     BookStatus `gorm:"not null;type:varchar(10);check:status IN ('available', 'placed', 'taken')" json:"status"`
}
