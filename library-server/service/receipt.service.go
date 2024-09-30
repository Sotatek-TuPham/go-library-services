package service

import (
	"errors"
	db "library-server/DB"
	"library-server/model"
)

func CreateReceipt(receipt *model.Receipt) error {
	result := db.DB.Create(receipt)
	if result.Error == nil {
		// Update the book status to "placed" only if it's currently "available"
		err := db.DB.Model(&model.Book{}).
			Where("id = ? AND status = ?", receipt.BookID, model.BookStatusAvailable).
			Update("status", model.BookStatusPlaced).Error
		if err != nil {
			// If updating book status fails, delete the created receipt
			db.DB.Delete(receipt)
			return err
		}
		// Check if the book status was actually updated
		var updatedBook model.Book
		if err := db.DB.First(&updatedBook, receipt.BookID).Error; err != nil {
			db.DB.Delete(receipt)
			return err
		}
		if updatedBook.Status != model.BookStatusPlaced {
			// If the book wasn't available, delete the created receipt
			db.DB.Delete(receipt)
			return errors.New("book is not available for placement")
		}
	}
	return result.Error
}
func GetReceiptByUserID(userID uint) ([]model.Receipt, error) {
	var receipts []model.Receipt
	result := db.DB.Where("user_id = ?", userID).Find(&receipts)
	return receipts, result.Error
}

func GetReceiptByID(id uint) (*model.Receipt, error) {
	var receipt model.Receipt
	result := db.DB.Preload("Book").First(&receipt, id)
	return &receipt, result.Error
}

func UpdateReceiptStatus(id uint, newStatus model.ReceiptStatus) error {
	result := db.DB.Model(&model.Receipt{}).Where("id = ?", id).Update("status", newStatus)
	// Get the receipt to check its current status
	var receipt model.Receipt
	if err := db.DB.First(&receipt, id).Error; err != nil {
		return err
	}

	// Update book status based on the new receipt status
	var bookStatus model.BookStatus
	switch newStatus {
	case model.ReceiptStatusPending:
		bookStatus = model.BookStatusPlaced
	case model.ReceiptStatusReturned:
		bookStatus = model.BookStatusAvailable
	default:
		// For other statuses, we don't change the book status
		return nil
	}

	// Update the book status
	if err := db.DB.Model(&model.Book{}).Where("id = ?", receipt.BookID).Update("status", bookStatus).Error; err != nil {
		return err
	}
	return result.Error
}

func DeleteReceipt(id uint) error {
	result := db.DB.Delete(&model.Receipt{}, id)
	return result.Error
}

func GetAllReceipts(page, pageSize int) ([]model.Receipt, int64, error) {
	var receipts []model.Receipt
	var totalCount int64

	query := db.DB.Model(&model.Receipt{}).Preload("Book")

	if page > 0 && pageSize > 0 {
		offset := (page - 1) * pageSize
		query = query.Offset(offset).Limit(pageSize)
	}

	err := query.Count(&totalCount).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Find(&receipts).Error
	if err != nil {
		return nil, 0, err
	}

	return receipts, totalCount, nil
}
