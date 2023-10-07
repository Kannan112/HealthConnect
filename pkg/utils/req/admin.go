package req

type AdminLogin struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
type Category struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}


