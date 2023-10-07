package interfaces

import (
	"context"

	"github.com/easy-health/pkg/domain"
	"github.com/easy-health/pkg/utils/req"
)

type DoctorUseCase interface {
	Registration(ctx context.Context, data req.DoctorRegistration, categoryId uint) error
	DoctorLogin(ctx context.Context, login req.DoctorLogin) (string, error)
	DoctorProfile(ctx context.Context, doctorId int) (req.DoctorProfile, error)
	ListCategory(ctx context.Context) ([]domain.Categories, error)
}
