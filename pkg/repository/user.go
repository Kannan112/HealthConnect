package repository

import (
	"context"
	"fmt"

	"github.com/easy-health/pkg/domain"
	interfaces "github.com/easy-health/pkg/repository/interface"
	"github.com/easy-health/pkg/utils/req"
	"github.com/easy-health/pkg/utils/res"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}

func (c *userDatabase) CreateUser(ctx context.Context, reg req.UserRegister) error {
	UserProfile := domain.User{
		UserName:  reg.UserName,
		FirstName: reg.FirstName,
		LastName:  reg.LastName,
		Email:     reg.Email,
		Age:       reg.Age,
		Password:  reg.Password,
	}
	if err := c.DB.Create(&UserProfile).Error; err != nil {
		return err
	}

	return nil
}

func (c *userDatabase) LoginUser(ctx context.Context, login req.UserLogin) (domain.User, error) {
	var Userdata domain.User
	fmt.Println(login.Email)
	if err := c.DB.Raw("select * from users where email=$1", login.Email).Scan(&Userdata).Error; err != nil {
		return Userdata, err
	}
	return Userdata, nil
}

func (c *userDatabase) CheckAccount(ctx context.Context, email string) (bool, error) {
	var check bool
	query := `select exists(select * from users where email=$1)`
	if err := c.DB.Raw(query, email).Scan(&check).Error; err != nil {
		return false, err
	}
	return check, nil
}

func (c *userDatabase) ListCategoryUser(ctx context.Context) ([]res.CategoriesUser, error) {
	var Categories []res.CategoriesUser
	query := `select * from categories`
	if err := c.DB.Raw(query).Scan(&Categories).Error; err != nil {
		return nil, err
	}
	return Categories, nil

}
