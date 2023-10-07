package usecase

import (
	"context"
	"errors"
	"time"

	domain "github.com/easy-health/pkg/domain"
	interfaces "github.com/easy-health/pkg/repository/interface"
	services "github.com/easy-health/pkg/usecase/interface"
	"github.com/easy-health/pkg/utils/req"
	"github.com/easy-health/pkg/utils/res"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AdminUseCase struct {
	adminRepo interfaces.AdminRepository
	docRepo   interfaces.DoctorRepository
}

func NewAdminUseCase(adminRepo interfaces.AdminRepository, doc interfaces.DoctorRepository) services.AdminUseCase {
	return &AdminUseCase{
		adminRepo: adminRepo,
		docRepo:   doc,
	}
}

func (c *AdminUseCase) AdminSignup(ctx context.Context, adminSignup req.AdminLogin) error {
	err := c.adminRepo.AdminSignup(ctx, adminSignup)
	if err != nil {
		return err
	}
	return nil
}
func (c *AdminUseCase) AdminLogin(ctx context.Context, adminLogin req.AdminLogin) (string, error) {
	data, err := c.adminRepo.AdminLogin(ctx, adminLogin)
	if err != nil {
		return "", errors.New("failed to find")
	}
	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(adminLogin.Password))
	if err != nil {
		return "", errors.New("wrong password")
	}

	claims := jwt.MapClaims{
		"id":  data.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte("strre"))
	if err != nil {
		return "", err
	}
	return ss, nil

}

func (c *AdminUseCase) CreateCateogry(ctx context.Context, category req.Category) error {
	check, err := c.adminRepo.CategoryCheck(category.Name)
	if err != nil {
		return err
	} else if check != false {
		return errors.New("category name already exists")
	} else {
		err := c.adminRepo.CreateCategory(ctx, category)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *AdminUseCase) DeleteCategory(ctx context.Context, categoryId int) error {

	if err := c.adminRepo.DeleteCategory(ctx, categoryId); err != nil {
		return err
	}
	return nil
}

func (c *AdminUseCase) ListCategory(ctx context.Context) ([]domain.Categories, error) {
	data, err := c.adminRepo.ListCategory(ctx)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (c *AdminUseCase) ListDoctors(ctx context.Context) ([]res.Doctors, error) {
	var data []res.Doctors
	data, err := c.adminRepo.WaitingList(ctx)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (c *AdminUseCase) AdminVerify(ctx context.Context, doctor_id int) error {
	check, _ := c.docRepo.CheckDoctorId(ctx, doctor_id)
	if !check {
		return errors.New("wrong id")
	}
	if err := c.adminRepo.AdminVerify(ctx, doctor_id); err != nil {
		return err
	}
	return nil

}
