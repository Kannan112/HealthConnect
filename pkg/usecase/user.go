package usecase

import (
	"context"
	"errors"
	"time"

	interfaces "github.com/easy-health/pkg/repository/interface"
	services "github.com/easy-health/pkg/usecase/interface"
	"github.com/easy-health/pkg/utils/req"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}
func (c *userUseCase) RegisterUser(ctx context.Context, reg req.UserRegister) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reg.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	reg.Password = string(hashedPassword)
	if err := c.userRepo.CreateUser(ctx, reg); err != nil {
		return err
	}
	return nil
}
func (c *userUseCase) UserLogin(ctx context.Context, Login req.UserLogin) (string, error) {
	UserProfile, err := c.userRepo.LoginUser(ctx, Login)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(UserProfile.Password), []byte(Login.Password))
	if err != nil {
		return "", errors.New("wrong password")
	}
	claims := jwt.MapClaims{
		"id":  UserProfile.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte("strre"))
	if err != nil {
		return "", err
	}
	return ss, nil
}
