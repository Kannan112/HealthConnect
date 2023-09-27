package usecase

import (
	"context"

	interfaces "github.com/easy-health/pkg/repository/interface"
	services "github.com/easy-health/pkg/usecase/interface"
	"github.com/easy-health/pkg/utils/req"
)

type AdminUseCase struct {
	adminRepo *interfaces.AdminRepository
}

func NewAdminUseCase(adminRepo *interfaces.AdminRepository) services.AdminUseCase {
	return &AdminUseCase{
		adminRepo: adminRepo,
	}
}

func (c *AdminUseCase) AdminSignup(ctx context.Context, adminSignup req.AdminLogin) error {
	err := c.AdminSignup(ctx, adminSignup)
	if err != nil {
		return err
	}
	return nil
}
func (c *AdminUseCase) AdminLogin(ctx context.Context, adminLogin req.AdminLogin) error {
	err := c.AdminSignup(ctx, adminLogin)
	if err != nil {
		return err
	}
	return nil
}
