package domain

type User struct {
	ID        uint   `json:"id" gorm:"unique;not null"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"not null,unique"`
	Password  string `json:"password"`
	Blocked   bool
}
