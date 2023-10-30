package interfaces

import (
	"context"

	"github.com/easy-health/pkg/utils/req"
)

type AuthUseCase interface {
	// Admin
	AdminSignup(ctx context.Context, adminSignup req.AdminLogin) error
	AdminLogin(ctx context.Context, adminLogin req.AdminLogin) (string, error)

	// Doctor
	Registration(ctx context.Context, data req.DoctorRegistration, categoryId uint) error
	DoctorLogin(ctx context.Context, login req.DoctorLogin) (string, error)

	// Patients
	RegisterUser(ctx context.Context, reg req.UserRegister) error
	UserLogin(ctx context.Context, Login req.UserLogin) (string, error)

	// Google
	GoogleLogin(ctx context.Context)
}
