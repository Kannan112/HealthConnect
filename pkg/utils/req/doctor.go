package req

type DoctorRegistration struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Specialise    string `json:"specialise"`
	LicenseNumber string `json:"license_number"`
}

type DoctorLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
