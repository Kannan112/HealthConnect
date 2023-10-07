package req

type UserLogin struct {
	Email    string
	Password string
}

type UserRegister struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"not null,unique"`
	Password  string `json:"password"`
}
