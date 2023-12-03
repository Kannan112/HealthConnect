package interfaces

import (
	"context"

	"github.com/easy-health/pkg/utils/req"
	"github.com/easy-health/pkg/utils/res"
)

type UserUseCase interface {
	RegisterUser(ctx context.Context, reg req.UserRegister) error
	UserLogin(ctx context.Context, Login req.UserLogin) (map[string]string, error)
	ListCategoryUser(ctx context.Context) ([]res.CategoriesUser, error)
}
