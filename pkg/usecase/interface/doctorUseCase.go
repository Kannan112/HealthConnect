package interfaces

import (
	"context"

	"github.com/easy-health/pkg/utils/req"
)

type DoctorUseCase interface {
	Registration(ctx context.Context, data req.DoctorRegistration) error
}
