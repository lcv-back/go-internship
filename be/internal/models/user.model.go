package models

import (
	"time"

	"gorm.io/gorm"
)

// User đại diện cho bảng người dùng trong cơ sở dữ liệu
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Password  string    `json:"password" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BeforeSave được gọi trước khi lưu dữ liệu
func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	// Cập nhật trường UpdatedAt mỗi khi có thay đổi
	user.UpdatedAt = time.Now()
	return
}

// Migrate để tự động tạo bảng User trong cơ sở dữ liệu
func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	return nil
}
