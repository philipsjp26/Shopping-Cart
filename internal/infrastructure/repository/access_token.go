package repository

import (
	"context"
	"go_playground/internal/core/entity"
	"go_playground/internal/core/port/repository"

	"gorm.io/gorm"
)

type accessTokenRepo struct {
	db *gorm.DB
}

func NewAccessTokenRepo(db *gorm.DB) repository.AccessTokenRepository {
	return &accessTokenRepo{db: db}
}
func (ac *accessTokenRepo) Store(ctx context.Context, param entity.AccessToken) error {
	err := ac.db.Create(&param).Error
	if err != nil {
		return err
	}
	return nil
}

func (ac *accessTokenRepo) FindOne(param entity.AccessToken) (*entity.AccessToken, error) {
	var (
		dest entity.AccessToken
	)
	err := ac.db.Where(&param).First(&dest).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &dest, nil
}
