package interfaces

import (
	"context"

	"github.com/easy-health/pkg/utils/req"
)

type DoctorRepository interface {
	EmailChecking(email string) (bool, error)
	Register(ctx context.Context, doctor req.DoctorRegistration, hashpassword string) error
}
