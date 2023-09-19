package interfaces

import (
	"context"

	"github.com/easy-health/pkg/utils/req"
)

type DoctorRepository interface {
	Register(ctx context.Context, doctor req.DoctorRegistration) error
}
