package req

type DoctorRegistration struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	About         string `json:"about"`
	LicenseNumber string `json:"license_number"`
}

type DoctorLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DoctorProfile struct {
	Name          string
	Email         string
	Specialise    string
	LicenseNumber string
	AverageRating int
	Reviews       []string //incomplete
}
