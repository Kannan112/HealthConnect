package interfaces

import (
	"context"

	"github.com/easy-health/pkg/domain"
)

type AuthRepository interface {
	SaveRefreshSession(ctx context.Context, refreshSession domain.RefreshSession) error
	FindRefreshSessionByTokenID(ctx context.Context, tokenID string) (domain.RefreshSession, error)
}
