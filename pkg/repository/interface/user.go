package interfaces

import (
	"context"

	"github.com/easy-health/pkg/domain"
	"github.com/easy-health/pkg/utils/req"
)

type UserRepository interface {
	CreateUser(ctx context.Context, reg req.UserRegister) error
	LoginUser(ctx context.Context, login req.UserLogin) (domain.Users, error)
}
