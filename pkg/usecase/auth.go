package usecase

import (
	interfaces "github.com/easy-health/pkg/repository/interface"
	services "github.com/easy-health/pkg/usecase/interface"
)

type AuthUseCase struct {
	adminRepo interfaces.AdminRepository
	docRepo   interfaces.DoctorRepository
	userRepo  interfaces.UserRepository
}

func AuthUser(adminRepo interfaces.AdminRepository, doc interfaces.DoctorRepository /* user interfaces.UserRepository*/) services.AdminUseCase {
	return &AdminUseCase{
		adminRepo: adminRepo,
		docRepo:   doc,
	}
}

// func (c *AuthUseCase) GoogleLogin(ctx context.Context, user domain.User) (userID uint, err error) {
// 	existUser, err := c.userRepo.CheckAccount(ctx, user.Email)
// 	if err != nil {
// 		return 0, err
// 	}
// 	if !existUser {

// 	}
// 	return 0, err
// }
