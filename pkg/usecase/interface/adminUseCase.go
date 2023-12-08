package interfaces

import (
	"context"

	domain "github.com/easy-health/pkg/domain"
	"github.com/easy-health/pkg/utils/req"
	"github.com/easy-health/pkg/utils/res"
)

type AdminUseCase interface {
	AdminSignup(ctx context.Context, adminSignup req.AdminLogin) error
	AdminLogin(ctx context.Context, adminLogin req.AdminLogin) (map[string]string, error)
	AdminLogout(ctx context.Context, adminID int) error

	CreateCateogry(ctx context.Context, category req.Category) error
	DeleteCategory(ctx context.Context, categoryId int) error
	ListCategory(ctx context.Context, page int, count int) ([]domain.Categories, error)
	UpdateCategory(ctx context.Context, category req.Category, id uint) error

	//doctorProfile
	ListVerifiedDoctores(ctx context.Context, page int, pageSize int) ([]res.Doctors, error)
	ListDoctores(ctx context.Context) ([]res.Doctors, error)
	ApprovePending(ctx context.Context) ([]res.Doctors, error)
	AdminVerify(ctx context.Context, doctor_id int) error
	// AdminApprove(ctx context.Context)
}
