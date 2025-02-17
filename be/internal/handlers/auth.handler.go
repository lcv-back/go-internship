package handlers

import (
	"be/internal/config"
	"be/internal/models"
	"be/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SignupHandler xử lý yêu cầu đăng ký người dùng
func SignupHandler(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	// Bind dữ liệu JSON từ request body vào struct input
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Tạo người dùng mới trong cơ sở dữ liệu
	user, err := repositories.CreateUser(input.Username, input.Password, input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Trả về thông tin người dùng đã đăng ký
	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	})
}

// LoginHandler xử lý đăng nhập người dùng
func LoginHandler(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := repositories.LoginUser(config.DB, input.Username, input.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"token": token})
}
