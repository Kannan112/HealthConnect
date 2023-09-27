package repository

import (
	"context"

	interfaces "github.com/easy-health/pkg/repository/interface"
	"github.com/easy-health/pkg/utils/req"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AdminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &AdminDatabase{DB}
}

func (c *AdminDatabase) AdminSignup(ctx context.Context, AdminSignup req.AdminLogin) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(AdminSignup.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	query := `INSERT INTO admins (name, password) VALUES ($1, $2)`
	err = c.DB.Raw(query, AdminSignup.Name, hashedPassword).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *AdminDatabase) AdminLogin(ctx context.Context, AdminLogin req.AdminLogin) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(AdminLogin.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	query := `select * from admins where name=$1 AND password=$2 `
	err = c.DB.Raw(query, AdminLogin.Name, hashedPassword).Error
	if err != nil {
		return err
	}
	return nil
}

