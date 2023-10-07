package interfaces

import (
	"context"

	"github.com/easy-health/pkg/utils/req"
)

type UserUseCase interface {
	RegisterUser(ctx context.Context, reg req.UserRegister) error
	UserLogin(ctx context.Context, Login req.UserLogin) (string, error)
}
