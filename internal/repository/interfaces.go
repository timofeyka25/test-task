package repository

import (
	"context"
	"gorm.io/gorm"
	"test-task/internal/entities"
)

type AuthRepository interface {
	CreateUser(ctx context.Context, tx *gorm.DB, user *entities.User) (int, error)
	GetUserByUsername(ctx context.Context, tx *gorm.DB, username string) (*entities.User, error)
}

type RecordRepository interface {
	GetAllRecords(ctx context.Context, tx *gorm.DB) ([]entities.Record, error)
	GetRandomRecords(ctx context.Context, tx *gorm.DB) ([]entities.Record, error)
}
