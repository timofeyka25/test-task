package services

import (
	"context"
	"errors"
	"test-task/internal/dto"
	"test-task/internal/entities"
	"test-task/internal/repository"
	"test-task/pkg/jwt"
	"test-task/pkg/utils"
)

type AuthService struct {
	authRepo     repository.AuthRepository
	recordRepo   repository.RecordRepository
	accessSecret string
}

type AccessConfig struct {
	AccessSecret string
}

func NewAuthService(
	authRepo repository.AuthRepository,
	recordRepo repository.RecordRepository,
	config *AccessConfig) *AuthService {
	return &AuthService{
		authRepo:     authRepo,
		recordRepo:   recordRepo,
		accessSecret: config.AccessSecret,
	}
}

func (s *AuthService) SignUp(ctx context.Context, req dto.SignUpRequest) (*dto.AuthResponse, error) {
	passwordHash, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := &entities.User{
		Username:     req.Username,
		PasswordHash: passwordHash,
	}
	id, err := s.authRepo.CreateUser(ctx, nil, user)
	if err != nil {
		return nil, errors.New("failed to create user")
	}
	if id == 0 {
		return nil, errors.New("user already exists")
	}

	accessToken, accessExpireTime, err := jwt.GenerateAccessToken(user.ID, s.accessSecret)
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	records, err := s.recordRepo.GetRandomRecords(ctx, nil)
	if err != nil {
		return nil, errors.New("failed to get random records")
	}

	return &dto.AuthResponse{
		Token:     accessToken,
		ExpiresAt: accessExpireTime,
		Records:   dto.RecordsFromEntities(records),
	}, nil
}

func (s *AuthService) SignIn(ctx context.Context, req dto.SignInRequest) (*dto.AuthResponse, error) {
	user, err := s.authRepo.GetUserByUsername(ctx, nil, req.Username)
	if err != nil {
		return nil, errors.New("failed to get user")
	}

	if !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
		return nil, errors.New("invalid password")
	}

	accessToken, accessExpireTime, err := jwt.GenerateAccessToken(user.ID, s.accessSecret)
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	records, err := s.recordRepo.GetRandomRecords(ctx, nil)
	if err != nil {
		return nil, errors.New("failed to get random records")
	}

	return &dto.AuthResponse{
		Token:     accessToken,
		ExpiresAt: accessExpireTime,
		Records:   dto.RecordsFromEntities(records),
	}, nil
}

func (s *AuthService) ValidateToken(tokenString string) (map[string]interface{}, error) {
	claims, err := jwt.ValidateAccessToken(tokenString, s.accessSecret)
	if err != nil {
		return nil, err
	}

	return claims, nil
}
