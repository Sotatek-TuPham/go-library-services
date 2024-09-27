package handler

import (
	"net/http"
	"strconv"

	"library-server/model"
	"library-server/service"

	"github.com/gin-gonic/gin"
)

// CreateReceipt godoc
// @Summary Create a new receipt
// @Description Create a new receipt with the input payload
// @Tags receipts
// @Accept json
// @Produce json
// @Param receipt body model.receipt true "Create receipt"
// @Success 201 {object} model.Receipt
// @Failure 400 {object} map[string]string
// @Router /receipts [post]
func CreateReceipt(c *gin.Context) {
	var receipt model.Receipt
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.CreateReceipt(&receipt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create receipt"})
		return
	}
	c.JSON(http.StatusCreated, receipt)
}

// GetReceiptByID godoc
// @Summary Get a receipt by ID
// @Description Get a single receipt by its ID
// @Tags receipts
// @Produce json
// @Param id path int true "Receipt ID"
// @Success 200 {object} model.Receipt
// @Failure 404 {object} map[string]string
// @Router /receipts/{id} [get]
func GetReceiptByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	receipt, err := service.GetReceiptByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found"})
		return
	}
	c.JSON(http.StatusOK, receipt)
}

// GetReceiptsByUserID godoc
// @Summary Get receipts by user ID
// @Description Get all receipts for a specific user
// @Tags receipts
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {array} model.Receipt
// @Failure 404 {object} map[string]string
// @Router /receipts/user/{user_id} [get]
func GetReceiptsByUserID(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Param("user_id"), 10, 32)
	receipts, err := service.GetReceiptByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Receipts not found"})
		return
	}
	c.JSON(http.StatusOK, receipts)
}

// UpdateReceiptStatus godoc
// @Summary Update a receipt's status
// @Description Update the status of an existing receipt
// @Tags receipts
// @Accept json
// @Produce json
// @Param id path int true "Receipt ID"
// @Param status body model.ReceiptStatus true "New receipt status"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /receipts/{id}/status [patch]
func UpdateReceiptStatus(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var newStatus model.ReceiptStatus
	if err := c.ShouldBindJSON(&newStatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.UpdateReceiptStatus(uint(id), newStatus); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to update receipt status"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Receipt status updated successfully"})
}

// DeleteReceipt godoc
// @Summary Delete a receipt
// @Description Delete a receipt by its ID
// @Tags receipts
// @Produce json
// @Param id path int true "Receipt ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /receipts/{id} [delete]
func DeleteReceipt(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := service.DeleteReceipt(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to delete receipt"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Receipt deleted successfully"})
}

// GetAllReceipts godoc
// @Summary Get all receipts
// @Description Get all receipts with pagination
// @Tags receipts
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /receipts [get]
func GetAllReceipts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	receipts, totalCount, err := service.GetAllReceipts(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch receipts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"receipts":    receipts,
		"total_count": totalCount,
		"page":        page,
		"page_size":   pageSize,
	})
}
