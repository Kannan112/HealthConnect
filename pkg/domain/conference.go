package domain

type Meeting struct {
	ID uint `json:"id" gorm:"primaryKey;not null"`
	
}
