package service

import (
	"context"
	"errors"
	db "order-server/DB"
	"order-server/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Register(ctx context.Context, user *model.User) error {
	// Check if user already exists
	var existingUser model.User
	if result := db.DB.Where("username = ? OR email = ?", user.Username, user.Email).First(&existingUser); result.Error == nil {
		return errors.New("user already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Create the user
	if result := db.DB.Create(user); result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *AuthService) Login(ctx context.Context, username, password string) (*model.User, string, error) {
	var user model.User
	if result := db.DB.Where("username = ?", username).First(&user); result.Error != nil {
		return nil, "", errors.New("user not found")
	}

	// Compare the provided password with the stored hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", errors.New("invalid password")
	}

	// Generate JWT token
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}).SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, "", errors.New("failed to generate token")
	}

	// Return user and token
	return &user, token, nil
}
