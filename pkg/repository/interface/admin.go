package interfaces

import (
	"context"

	"github.com/easy-health/pkg/utils/req"
)

type AdminRepository interface {
	AdminSignup(ctx context.Context, AdminSignup req.AdminLogin) error
	AdminLogin(ctx context.Context, AdminLogin req.AdminLogin) error
	AdminCateogry(ctx context.Context,)
}
