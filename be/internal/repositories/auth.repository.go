package repositories

import (
	"be/internal/config"
	"be/internal/models"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CreateUser tạo người dùng mới trong cơ sở dữ liệu
func CreateUser(username, password, email string) (*models.User, error) {
	// Mã hóa mật khẩu trước khi lưu
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Tạo đối tượng User
	user := &models.User{
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
	}

	// Lưu người dùng vào cơ sở dữ liệu
	if err := config.DB.Create(user).Error; err != nil {
		return nil, err
	}

	// Trả về người dùng vừa tạo
	return user, nil
}

// LoginUser kiểm tra thông tin đăng nhập và trả về JWT
func LoginUser(db *gorm.DB, username, password string) (string, error) {
	var user models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid password")
	}

	token, err := generateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

// generateJWT tạo JWT token
func generateJWT(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTSecretKey))
}
