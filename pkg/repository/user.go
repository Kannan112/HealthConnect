package repository

import (
	"context"

	"github.com/easy-health/pkg/domain"
	interfaces "github.com/easy-health/pkg/repository/interface"
	"github.com/easy-health/pkg/utils/req"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}

func (c *userDatabase) CreateUser(ctx context.Context, reg req.UserRegister) error {
	UserProfile := domain.Users{
		FirstName: reg.FirstName,
		LastName:  reg.LastName,
		Email:     reg.Email,
		Password:  reg.Password,
		Blocked:   false,
	}
	if err := c.DB.Create(&UserProfile).Error; err != nil {
		return err
	}

	return nil
}

func (c *userDatabase) LoginUser(ctx context.Context, login req.UserLogin) (domain.Users, error) {
	var Userdata domain.Users
	if err := c.DB.Raw("select * from doctors where email=$1", login.Email).Scan(&Userdata).Error; err != nil {
		return Userdata, err
	}
	return Userdata, nil
}
