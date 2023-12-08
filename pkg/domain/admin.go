package domain

type Admin struct {
	ID           uint   `json:"id" gorm:"unique;not null"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	OnlineStatus bool   `json:"online_status"`
}
