package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/easy-health/pkg/api/middleware/token"
	"github.com/easy-health/pkg/config"
	domain "github.com/easy-health/pkg/domain"
	interfaces "github.com/easy-health/pkg/repository/interface"
	services "github.com/easy-health/pkg/usecase/interface"
	"github.com/easy-health/pkg/utils/req"
	"github.com/easy-health/pkg/utils/res"
	"golang.org/x/crypto/bcrypt"
)

type AdminUseCase struct {
	adminRepo    interfaces.AdminRepository
	categoryRepo interfaces.CategoryRepository
	docRepo      interfaces.DoctorRepository
	authRepo     interfaces.AuthRepository
	config       config.Config
}

func NewAdminUseCase(adminRepo interfaces.AdminRepository, doc interfaces.DoctorRepository, auth interfaces.AuthRepository, con config.Config, categoryRepo interfaces.CategoryRepository) services.AdminUseCase {
	return &AdminUseCase{
		adminRepo:    adminRepo,
		categoryRepo: categoryRepo,
		docRepo:      doc,
		config:       con,
		authRepo:     auth,
	}
}

func (c *AdminUseCase) AdminSignup(ctx context.Context, adminSignup req.AdminLogin) error {
	check, err := c.adminRepo.AdminCheck(ctx, adminSignup.Email)
	if err != nil {
		return err
	}
	fmt.Println("test1")
	if check {
		return fmt.Errorf("email is already registered")
	}
	err = c.adminRepo.AdminSignup(ctx, adminSignup)
	if err != nil {
		return err
	}
	return nil
}
func (c *AdminUseCase) AdminLogin(ctx context.Context, adminLogin req.AdminLogin) (map[string]string, error) {
	data, err := c.adminRepo.AdminLogin(ctx, adminLogin)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(adminLogin.Password))
	if err != nil {
		return nil, errors.New("wrong password")
	}

	if err := c.adminRepo.OnlineStatusUpdate(ctx, data.ID, true); err != nil {
		return nil, fmt.Errorf("failed to change%v", err)
	}
	// Generate an access token
	AccessToken, err := token.GenerateAccessToken(int(data.ID), "admin")
	fmt.Println(AccessToken)

	if err != nil {
		return nil, fmt.Errorf("failed to create access token")
	}

	// Generate a refresh token
	RefreshToken, err := token.GenerateRefreshToken(int(data.ID), "admin")
	if err != nil {
		return nil, err
	}

	tokenMap := map[string]string{
		"access_token":  AccessToken,
		"refresh_token": RefreshToken,
	}

	return tokenMap, nil
}
func (c *AdminUseCase) AdminLogout(ctx context.Context, adminID int) error {
	adminUUID := uint(adminID)
	if err := c.adminRepo.OnlineStatusUpdate(ctx, adminUUID, false); err != nil {
		return err
	}
	return nil
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

	if err := c.categoryRepo.DeleteCategory(ctx, categoryId); err != nil {
		return err
	}
	return nil
}

func (c *AdminUseCase) ListCategory(ctx context.Context, page int, count int) ([]domain.Categories, error) {
	data, err := c.categoryRepo.ListCategory(ctx)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (c *AdminUseCase) UpdateCategory(ctx context.Context, category req.Category, id uint) error {
	// data, err := c.categoryRepo.FeatchCategoryDetails(ctx, category.Name)
	// if err != nil {
	// 	return err
	// }
	if category.Name != "" {
		err := c.categoryRepo.UpdateCategoryName(ctx, category.Name, uint(id))
		if err != nil {
			return err
		}
	}
	if category.Description != "" {
		err := c.categoryRepo.UpdateCategoryDescription(ctx, category.Description, uint(id))
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *AdminUseCase) ApprovePending(ctx context.Context) ([]res.Doctors, error) {
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

func (c *AdminUseCase) ListDoctores(ctx context.Context) ([]res.Doctors, error) {
	data, err := c.adminRepo.ListDoctores(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get data from doctors%v", err.Error())
	}
	return data, nil
}

func (c *AdminUseCase) ListVerifiedDoctores(ctx context.Context, page int, pageSize int) ([]res.Doctors, error) {
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * pageSize
	data, err := c.adminRepo.ListVerifiedDoctores(ctx, offset, pageSize)
	if err != nil {
		return nil, fmt.Errorf("failed to get data from doctors%v", err.Error())
	}
	return data, nil
}
