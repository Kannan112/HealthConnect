package usecase

import (
	"context"
	"errors"

	"github.com/easy-health/pkg/api/middleware/token"
	interfaces "github.com/easy-health/pkg/repository/interface"
	services "github.com/easy-health/pkg/usecase/interface"
	"github.com/easy-health/pkg/utils/req"
	"github.com/easy-health/pkg/utils/res"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	user     interfaces.UserRepository
	admin    interfaces.AdminRepository
	category interfaces.CategoryRepository
}

func NewUserUseCase(UserRepo interfaces.UserRepository, AdminRepo interfaces.AdminRepository, categoryRepo interfaces.CategoryRepository) services.UserUseCase {
	return &userUseCase{
		user:     UserRepo,
		admin:    AdminRepo,
		category: categoryRepo,
	}
}
func (c *userUseCase) RegisterUser(ctx context.Context, reg req.UserRegister) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reg.Password), bcrypt.DefaultCost)
	if err != nil {
		return err

	}
	reg.Password = string(hashedPassword)
	if err := c.user.CreateUser(ctx, reg); err != nil {
		return err
	}
	return nil
}
func (c *userUseCase) UserLogin(ctx context.Context, Login req.UserLogin) (map[string]string, error) {
	exist, err := c.user.CheckAccount(ctx, Login.Email)
	if err != nil {
		return nil, err

	} else if !exist {
		return nil, errors.New("account not registered")

	}
	UserProfile, err := c.user.LoginUser(ctx, Login)
	if err != nil {
		return nil, err

	}
	err = bcrypt.CompareHashAndPassword([]byte(UserProfile.Password), []byte(Login.Password))
	if err != nil {
		return nil, errors.New("wrong password")
	}

	accessToken, err := token.GenerateAccessToken(int(UserProfile.ID), "user")
	if err != nil {
		return nil, err
	}
	refreshToken, err := token.GenerateRefreshToken(int(UserProfile.ID), "user")
	if err != nil {
		return nil, err
	}
	tokenMap := map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	return tokenMap, nil
}

func (c *userUseCase) ListCategoryUser(ctx context.Context) ([]res.CategoriesUser, error) {
	var data []res.CategoriesUser
	data, err := c.user.ListCategoryUser(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}
