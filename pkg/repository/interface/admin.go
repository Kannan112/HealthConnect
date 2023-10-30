package interfaces

import (
	"context"

	"github.com/easy-health/pkg/domain"
	"github.com/easy-health/pkg/utils/req"
	"github.com/easy-health/pkg/utils/res"
)

type AdminRepository interface {
	AdminSignup(ctx context.Context, AdminSignup req.AdminLogin) error
	AdminLogin(ctx context.Context, AdminLogin req.AdminLogin) (data domain.Admin, err error)

	CategoryCheck(categoryName string) (bool, error)

	CreateCategory(ctx context.Context, category req.Category) error
	DeleteCategory(ctx context.Context, category_id int) error
	ListCategory(ctx context.Context) ([]domain.Categories, error)

	//List Doctors
	ListDoctores(ctc context.Context) ([]res.Doctors, error)
	WaitingList(ctx context.Context) ([]res.Doctors, error)
	AdminVerify(ctx context.Context, doctor_id int) error
}
