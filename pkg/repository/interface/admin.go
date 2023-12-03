package interfaces

import (
	"context"

	"github.com/easy-health/pkg/domain"
	"github.com/easy-health/pkg/utils/req"
	"github.com/easy-health/pkg/utils/res"
)

type AdminRepository interface {
	AdminCheck(ctx context.Context, Email string) (bool, error)
	AdminSignup(ctx context.Context, AdminSignup req.AdminLogin) error
	AdminLogin(ctx context.Context, AdminLogin req.AdminLogin) (data domain.Admin, err error)
	OnlineStatusUpdate(ctx context.Context, adminID uint, value bool) error

	CategoryCheck(categoryName string) (bool, error)

	CreateCategory(ctx context.Context, category req.Category) error
	DeleteCategory(ctx context.Context, category_id int) error
	ListCategory(ctx context.Context) ([]domain.Categories, error)

	//List Doctors
	ListVerifiedDoctores(ctx context.Context) ([]res.Doctors, error)
	ListDoctores(ctx context.Context) ([]res.Doctors, error)
	WaitingList(ctx context.Context) ([]res.Doctors, error)
	AdminVerify(ctx context.Context, doctor_id int) error
}
