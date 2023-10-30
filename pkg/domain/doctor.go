package domain

import "time"

type Doctors struct {
	ID            uint      `json:"id" gorm:"primaryKey;not null"`
	Name          string    `json:"name" gorm:"not null" binding:"required,min=3,max=15"`
	Email         string    `json:"email" gorm:"unique;not null" binding:"required,email"`
	Password      string    `json:"password" binding:"required"`
	About         string    `json:"about" gorm:"required,min=10,max30"`
	LicenseNumber string    `json:"license_number"`
	Verified      bool      `json:"verified" gorm:"default:false"`
	BlockStatus   bool      `json:"block_status" gorm:"not null;default:false"`
	CreatedAt     time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt     time.Time `json:"updated_at"`
	CategoriesId  uint
}

// type DoctorsHistory struct {
// 	ID             uint
// 	DoctorID       uint   `json:"doctor_id" gorm:"not null"` // Foreign key referencing Doctors.ID
// 	WorkExperience string `json:"work_experience"`
// 	ValidationInfo string `json:"validation_info"`
// }

type Reviews struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	DoctorsID uint   `json:"doctors_id"`
	UsersID   uint   `json:"users_id"`
	Rating    int    `json:"rating" gorm:"not null" binding:"required,gte=1,lte=5"`
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
