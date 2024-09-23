package model

type Book struct {
	ID         uint     `gorm:"primaryKey;autoIncrement" json:"id"`
	Title      string   `gorm:"not null" json:"title"`
	Author     string   `gorm:"not null" json:"author"`
	CategoryID uint     `gorm:"not null" json:"category_id"`
	Category   Category `gorm:"foreignKey:CategoryID" json:"category"`
}
