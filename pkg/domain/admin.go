package domain

import (
	"time"
)

type Admin struct {
	ID       uint   `json:"id" gorm:"unique;not null"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Categories struct {
	ID          int       `json:"id" gorm:"unique;not null"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt   time.Time `json:"updated_at"`
}
