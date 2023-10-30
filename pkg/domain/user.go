package domain

import "time"

type User struct {
	ID          uint      `json:"id" gorm:"unique;not null"`
	GoogleImage string    `json:"google_profile_image"`
	UserName    string    `json:"user_name" gorm:"unique"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email" gorm:"not null,unique"`
	Age         uint      `json:"age" binding:"required,numeric"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt   time.Time `json:"updated_at"`
	BlockStatus bool      `json:"block_status" gorm:"not null;default:false"`
}
