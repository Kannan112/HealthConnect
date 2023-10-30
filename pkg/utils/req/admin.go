package req

type AdminLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Category struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
