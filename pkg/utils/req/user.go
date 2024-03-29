package req

type UserLogin struct {
	Email    string `json:"email" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
}

type UserRegister struct {
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"not null,unique"`
	Age       uint   `json:"age"`
	Password  string `json:"password"`
}
