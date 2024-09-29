package handler

import (
	"net/http"
	"order-server/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserHandler handles HTTP requests related to user operations
type UserHandler struct {
	userService *service.UserService
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// PlaceReceipt godoc
// @Summary Place a new receipt
// @Description Place a new receipt for a user and book
// @Tags receipts
// @Accept json
// @Produce json
// @Param request body PlaceReceiptRequest true "Receipt details"
// @Param Authorization header string true "Bearer token" default(Bearer <Add access token here>)
// @Success 201 "Created"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /receipts [post]
func (h *UserHandler) PlaceReceipt(c *gin.Context) {
	var request PlaceReceiptRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid request body"})
		return
	}

	err := h.userService.PlaceReceipt(c.Request.Context(), request.UserID, request.BookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

// CancelReceipt godoc
// @Summary Cancel a receipt
// @Description Cancel an existing receipt
// @Tags receipts
// @Accept json
// @Produce json
// @Param id path int true "Receipt ID"
// @Param Authorization header string true "Bearer token" default(Bearer <Add access token here>)
// @Success 200 "OK"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /receipts/{id}/cancel [post]
func (h *UserHandler) CancelReceipt(c *gin.Context) {
	receiptID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid receipt ID"})
		return
	}

	err = h.userService.CancelReceipt(c.Request.Context(), uint(receiptID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// GetReceiptsByUserID godoc
// @Summary Get receipts for authenticated user
// @Description Get receipts for the authenticated user
// @Tags receipts
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token" default(Bearer <Add access token here>)
// @Security BearerAuth
// @Success 200 {array} service.Receipt "Receipts"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /users/receipts [get]
func (h *UserHandler) GetReceiptsByUserID(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "User not authenticated"})
		return
	}

	receipts, err := h.userService.GetReceiptsByUserID(c.Request.Context(), userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, receipts)
}

// PlaceReceiptRequest represents the request body for placing a receipt
type PlaceReceiptRequest struct {
	UserID uint `json:"user_id"`
	BookID uint `json:"book_id"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}
