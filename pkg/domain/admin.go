package domain

import (
	"time"
)

type Admin struct {
	ID       uint   `json:"id" gorm:"unique;not null"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Categories struct {
	ID          int    `json:"id" gorm:"unique;not null"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   time.Time
}
