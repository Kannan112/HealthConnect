package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/easy-health/pkg/domain"
	interfaces "github.com/easy-health/pkg/repository/interface"
	services "github.com/easy-health/pkg/usecase/interface"
	"github.com/easy-health/pkg/utils/req"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type DoctorUseCase struct {
	doctorRepo   interfaces.DoctorRepository
	adminRepo    interfaces.AdminRepository
	categoryRepo interfaces.CategoryRepository
}

func NewDoctorUseCase(doctorRepo interfaces.DoctorRepository, adminRepo interfaces.AdminRepository, categoryRepo interfaces.CategoryRepository) services.DoctorUseCase {
	return &DoctorUseCase{
		doctorRepo:   doctorRepo,
		adminRepo:    adminRepo,
		categoryRepo: categoryRepo,
	}
}

func (connect *DoctorUseCase) Registration(ctx context.Context, data req.DoctorRegistration, categoryId uint) error {
	category, err := connect.categoryRepo.CategoryIdCheck(ctx, categoryId)
	if err != nil {
		return err
	}
	if !category {
		return errors.New("category dont exist")
	}

	check, err := connect.doctorRepo.EmailChecking(data.Email)
	if err != nil {
		return err
	}
	if check {
		return errors.New("email is already registered")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	if err = connect.doctorRepo.Register(ctx, data, string(hashedPassword), categoryId); err != nil {
		return err
	}
	return nil
}

func (c *DoctorUseCase) DoctorLogin(ctx context.Context, login req.DoctorLogin) (string, error) {
	data, err := c.doctorRepo.Login(ctx, login)
	if err != nil {
		return "", err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(login.Password)); err != nil {
		return "", errors.New("wrong password")
	}
	if data.Verified != true {
		return "", errors.New("waiting for admin to approve")
	}
	claims := jwt.MapClaims{
		"id":  data.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("strre"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (c *DoctorUseCase) DoctorProfile(ctx context.Context, doctorId int) (req.DoctorProfile, error) {
	doctorProfile, err := c.doctorRepo.Profile(ctx, doctorId)
	if err != nil {
		return doctorProfile, err
	}
	return doctorProfile, nil
}

func (c *DoctorUseCase) ListCategory(ctx context.Context) ([]domain.Categories, error) {
	data, err := c.adminRepo.ListCategory(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil

}
