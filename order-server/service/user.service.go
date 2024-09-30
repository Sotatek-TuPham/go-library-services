package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type UserService struct {
	client  *http.Client
	baseURL string
}

func NewUserService(baseURL string) *UserService {
	return &UserService{
		client:  &http.Client{},
		baseURL: baseURL,
	}
}

func (s *UserService) PlaceReceipt(ctx context.Context, userID, bookID uint) error {
	url := fmt.Sprintf("%s/receipts", s.baseURL)

	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Body = io.NopCloser(bytes.NewBuffer([]byte(fmt.Sprintf(`{"user_id": %d, "book_id": %d}`, userID, bookID))))
	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func (s *UserService) CancelReceipt(ctx context.Context, receiptID uint) error {
	url := fmt.Sprintf("%s/receipts/%d/status", s.baseURL, receiptID)

	req, err := http.NewRequestWithContext(ctx, "PATCH", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Body = io.NopCloser(bytes.NewBuffer([]byte(`{"status": "canceled"}`)))
	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func (s *UserService) GetReceiptsByUserID(ctx context.Context, userID uint) ([]Receipt, error) {
	url := fmt.Sprintf("%s/receipts/user/%d", s.baseURL, userID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var receipts []Receipt
	if err := json.NewDecoder(resp.Body).Decode(&receipts); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return receipts, nil
}

type Receipt struct {
	ID     uint      `json:"id"`
	UserID uint      `json:"user_id"`
	BookID uint      `json:"book_id"`
	Date   time.Time `json:"date"`
}
