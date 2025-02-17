package main

import (
	"be/internal/config"
	"be/internal/handlers"
	"be/internal/models"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	log.Println("Secret Key: ", config.JWTSecretKey)

	// Thực hiện migration để tạo bảng User nếu chưa có
	if err := models.Migrate(config.DB); err != nil {
		log.Fatal("Error migrating database:", err)
	}

	// Khởi tạo Gin router
	r := gin.Default()

	// Routes cho auth service
	r.POST("/signup", handlers.SignupHandler)
	r.POST("/login", handlers.LoginHandler)

	r.Run(":8081")
}
