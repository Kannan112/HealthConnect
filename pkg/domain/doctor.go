package domain

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
