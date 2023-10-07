package interfaces

import (
	"context"

	domain "github.com/easy-health/pkg/domain"
	"github.com/easy-health/pkg/utils/req"
	"github.com/easy-health/pkg/utils/res"
)

type AdminUseCase interface {
	AdminSignup(ctx context.Context, adminSignup req.AdminLogin) error
	AdminLogin(ctx context.Context, adminLogin req.AdminLogin) (string, error)

	CreateCateogry(ctx context.Context, category req.Category) error
	DeleteCategory(ctx context.Context, categoryId int) error
	ListCategory(ctx context.Context) ([]domain.Categories, error)

	//doctorProfile
	ListDoctors(ctx context.Context) ([]res.Doctors, error)
	AdminVerify(ctx context.Context, doctor_id int) error
	// AdminApprove(ctx context.Context)
}
