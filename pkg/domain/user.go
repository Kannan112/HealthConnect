package domain

type Users struct {
	ID        uint   `json:"id" gorm:"unique;not null"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Blocked   bool
}
