package req

import "time"

type Categories struct {
	ID          int    `json:"id" gorm:"unique;not null"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateSlot struct {
	AppointmentTime time.Time `json:"appointment_time"`
	Discription     string    `json:"discription"`
}

type Category struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
