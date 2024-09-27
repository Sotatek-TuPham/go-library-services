package model

import (
	"time"

	"gorm.io/gorm"
)

type ReceiptStatus string

const (
	ReceiptStatusPending  ReceiptStatus = "pending"
	ReceiptStatusOwned    ReceiptStatus = "owned"
	ReceiptStatusReturned ReceiptStatus = "returned"
)

type Receipt struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint           `json:"user_id"`
	BookID    uint           `json:"book_id"`
	Book      Book           `gorm:"foreignKey:BookID" json:"book"`
	Status    ReceiptStatus  `gorm:"not null;type:varchar(10);check:status IN ('pending', 'owned', 'returned')" json:"status"`
	DueDate   time.Time      `json:"due_date"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
