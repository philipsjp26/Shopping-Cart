package repository

import (
	"context"
	"go_playground/internal/core/entity"
)

type AccessTokenRepository interface {
	Store(ctx context.Context, param entity.AccessToken) error
}
