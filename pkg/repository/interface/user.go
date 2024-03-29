package interfaces

import (
	"context"

	"github.com/easy-health/pkg/domain"
	"github.com/easy-health/pkg/utils/req"
	"github.com/easy-health/pkg/utils/res"
)

type UserRepository interface {
	CheckAccount(ctx context.Context, email string) (bool, error)
	CreateUser(ctx context.Context, reg req.UserRegister) error
	LoginUser(ctx context.Context, login req.UserLogin) (domain.User, error)
	ListCategoryUser(ctx context.Context) ([]res.CategoriesUser, error)
}
