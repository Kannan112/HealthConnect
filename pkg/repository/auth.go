package repository

import (
	"context"

	"github.com/easy-health/pkg/domain"
	interfaces "github.com/easy-health/pkg/repository/interface"
	"gorm.io/gorm"
)

type authDatabase struct {
	DB *gorm.DB
}

func NewAuthRepository(DB *gorm.DB) interfaces.AuthRepository {
	return &authDatabase{
		DB,
	}
}

func (c *authDatabase) SaveRefreshSession(ctx context.Context, refreshSession domain.RefreshSession) error {
	query := `INSERT INTO refresh_sessions (token_id, users_id, refresh_token, expire_at) 
VALUES ($1, $2, $3, $4)`
	err := c.DB.Exec(query, refreshSession.TokenID, refreshSession.UsersID, refreshSession.RefreshToken, refreshSession.ExpireAt).Error

	return err
}
func (c *authDatabase) FindRefreshSessionByTokenID(ctx context.Context, tokenID string) (refreshSession domain.RefreshSession, err error) {
	query := `SELECT * FROM refresh_sessions WHERE token_id = $1`

	err = c.DB.Raw(query, tokenID).Scan(&refreshSession).Error

	return
}
