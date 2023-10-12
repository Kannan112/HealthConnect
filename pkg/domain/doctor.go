package domain

import "time"

type Doctors struct {
	ID            uint   `json:"id" gorm:"primaryKey;not null"`
	Name          string `json:"name" gorm:"not null" binding:"required,min=3,max=15"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Specialise    string `json:"specialise"`
	LicenseNumber string `json:"license_number"`
	Approved      bool
	CategoriesId  uint
}

type Reviews struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	DoctorsID uint   `json:"doctors_id"`
	UsersID   uint   `json:"users_id"`
	Rating    int    `json:"rating" gorm:"not null"` // Fixed the typo here
	Comment   string `json:"comment"`
}

type Appointment struct {
	ID              uint      `json:"id" gorm:"primaryKey;not null"`
	DoctorsID       uint      `json:"doctors_id" gorm:"not null"`
	Doctors         Doctors   `json:"doctors" gorm:"foreignKey:DoctorsID"`
	AppointmentTime time.Time `json:"appointment_time" gorm:"not null"`
	Discription     string    `json:"description"`
	UserID          uint      `json:"user_id"`
	User            User      `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt       time.Time `json:"created_at"`
}
