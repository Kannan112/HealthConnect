package interfaces

import (
	"context"

	"github.com/easy-health/pkg/utils/req"
)

type AdminUseCase interface {
	AdminSignup(ctx context.Context, adminSignup req.AdminLogin) error
	AdminLogin(ctx context.Context, adminLogin req.AdminLogin) error
}
