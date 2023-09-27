package domain

type Admin struct {
	ID       int    `json:"id" gorm:"unique;not null"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
