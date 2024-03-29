package repository

import (
	"context"
	"fmt"

	"github.com/easy-health/pkg/domain"
	interfaces "github.com/easy-health/pkg/repository/interface"
	"github.com/easy-health/pkg/utils/req"
	"github.com/easy-health/pkg/utils/res"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AdminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &AdminDatabase{DB}
}

func (c *AdminDatabase) AdminCheck(ctx context.Context, Email string) (bool, error) {
	var check bool
	query := `SELECT EXISTS(select 1 FROM admins where email=$1)`
	if err := c.DB.Raw(query, Email).Scan(&check).Error; err != nil {
		return false, err
	}
	return check, nil
}

func (c *AdminDatabase) AdminSignup(ctx context.Context, AdminSignup req.AdminLogin) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(AdminSignup.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	query := `INSERT INTO admins (email, password) VALUES ($1, $2)`
	err = c.DB.Exec(query, AdminSignup.Email, hashedPassword).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *AdminDatabase) AdminLogin(ctx context.Context, AdminLogin req.AdminLogin) (data domain.Admin, err error) {
	query := `select * from admins where email=$1`
	err = c.DB.Raw(query, AdminLogin.Email).Scan(&data).Error
	if err != nil {
		return data, err
	}
	return data, nil

}
func (c *AdminDatabase) OnlineStatusUpdate(ctx context.Context, adminID uint, value bool) error {
	query := `UPDATE admins SET online_status = $1 WHERE id = $2`
	err := c.DB.Exec(query, value, adminID).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *AdminDatabase) CategoryCheck(categoryName string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM categories WHERE name = $1)`
	var check bool
	err := c.DB.Raw(query, categoryName).Scan(&check).Error
	if err != nil {
		return false, err
	}
	fmt.Println(check)
	return check, nil
}
func (c *AdminDatabase) CreateCategory(ctx context.Context, category req.Category) error {
	query := `INSERT INTO categories (name, created_at, description) VALUES ($1,NOW(), $2)`
	err := c.DB.Exec(query, category.Name, category.Description).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *AdminDatabase) ListCategory(ctx context.Context) (data []domain.Categories, err error) {
	query := `select * from categories`
	err = c.DB.Raw(query).Scan(&data).Error
	if err != nil {
		return data, err
	}
	return data, nil
}

func (c *AdminDatabase) DeleteCategory(ctx context.Context, category_id int) error {
	query := `DELETE FROM categories WHERE id = $1`
	if err := c.DB.Exec(query, category_id).Error; err != nil {
		return err
	}
	return nil
}

func (c *AdminDatabase) ListDoctores(ctx context.Context) ([]res.Doctors, error) {
	var profile []res.Doctors
	query := `select * from doctors where approved=false`
	if err := c.DB.Raw(query).Scan(&profile).Error; err != nil {
		return profile, err
	}
	return profile, nil
}

func (c *AdminDatabase) WaitingList(ctx context.Context) ([]res.Doctors, error) {
	var profile []res.Doctors
	query := `select * from doctors where Verified=false`
	if err := c.DB.Raw(query).Scan(&profile).Error; err != nil {
		return profile, err
	}
	return profile, nil
}

func (c *AdminDatabase) AdminVerify(ctx context.Context, doctor_id int) error {
	if err := c.DB.Exec(`UPDATE doctors SET approved=$1 where id=$2`, true, doctor_id).Error; err != nil {
		return err
	}
	return nil
}

func (c *AdminDatabase) ListVerifiedDoctores(ctx context.Context, offset int, pageSize int) ([]res.Doctors, error) {
	var profile []res.Doctors
	query := `SELECT * FROM doctors WHERE verified = true OFFSET $1 LIMIT $2;`

	if err := c.DB.Raw(query, offset, pageSize).Scan(&profile).Error; err != nil {
		return profile, err
	}
	return profile, nil
}
