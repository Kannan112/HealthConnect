package res

type Doctors struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Specialise    string `json:"specialise"`
	LicenseNumber string `json:"license_number"`
	Approved      bool
}
