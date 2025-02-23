package pgsql

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"test-task/internal/entities"
	"test-task/internal/repository"
)

type authRepository struct {
	conn *gorm.DB
}

func NewAuthRepository(conn *gorm.DB) repository.AuthRepository {
	return &authRepository{conn: conn}
}

func (r *authRepository) getDB(tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	}
	return r.conn
}

func (r *authRepository) CreateUser(ctx context.Context, tx *gorm.DB, user *entities.User) (int, error) {
	err := r.getDB(tx).WithContext(ctx).Create(&user).Error
	if err != nil {
		if errors.As(err, &gorm.ErrDuplicatedKey) {
			return 0, nil
		}
		return 0, err
	}
	return user.ID, nil
}

func (r *authRepository) GetUserByUsername(ctx context.Context, tx *gorm.DB, username string) (*entities.User, error) {
	var user entities.User

	err := r.getDB(tx).WithContext(ctx).Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *authRepository) GetUserByID(ctx context.Context, tx *gorm.DB, id int) (*entities.User, error) {
	var user entities.User

	err := r.getDB(tx).WithContext(ctx).Where("id = ?", id).First(&user).Error
	return &user, err
}
