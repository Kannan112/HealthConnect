package interfaces

import (
	"context"

	"github.com/easy-health/pkg/domain"
	"github.com/easy-health/pkg/utils/req"
)

type DoctorRepository interface {
	EmailChecking(email string) (bool, error)
	Register(ctx context.Context, doctor req.DoctorRegistration, hashpassword string, categoryId uint) error
	Login(ctx context.Context, data req.DoctorLogin) (domain.Doctors, error)
	Profile(ctx context.Context, id int) (req.DoctorProfile, error)
	CheckDoctorId(ctx context.Context, id int) (bool, error)
	//CategoryIdCheck(ctx context.Context, categoryId uint) (bool, error)
	//CreateAppointment(ctx context.Context, slotCreate req.CreateSlot) (req.CreateSlot, error)
}
