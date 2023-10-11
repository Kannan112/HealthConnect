package domain

import "time"

type Chat struct {
	ID        uint      `json:"id" gorm:"primaryKey;not null"`
	UserID    uint      `json:"user1_id" gorm:"not null"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	DoctorID  uint      `json:"doctor_id" gorm:"not null"`
	Doctor    Doctors   `json:"doctor" gorm:"foreignKey:DoctorID"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
}

type Message struct {
	ID        uint      `json:"id" gorm:"primaryKey;not null"`
	ChatID    uint      `json:"chat_id" gorm:"not null"`
	Chat      Chat      `json:"chat" gorm:"foreignKey:ChatID"`
	SenderID  uint      `json:"sender_id" gorm:"not null"`
	Sender    User      `json:"sender" gorm:"foreignKey:SenderID"`
	Content   string    `json:"content" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at"`
}