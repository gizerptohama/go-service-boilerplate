package usecase

import (
	"boilerplate/internal/dto"
	"boilerplate/internal/entity"
	apperrors "boilerplate/internal/errors"
	"boilerplate/internal/repository"
	"boilerplate/internal/utils"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Register(context.Context, dto.RegisterRequest) (*entity.User, error)
	Login(context.Context, dto.LoginRequest) (string, error)
}

type authUsecase struct {
	userRepository repository.UserRepository
	tokenGenerator utils.TokenGenerator
}

func NewAuthUsecase(r repository.UserRepository, tg utils.TokenGenerator) AuthUsecase {
	return &authUsecase{
		userRepository: r,
		tokenGenerator: tg,
	}
}

func (u *authUsecase) Register(ctx context.Context, req dto.RegisterRequest) (*entity.User, error) {
	var err error
	var result *entity.User
	req.Password, err = utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	result, err = u.userRepository.CreateUser(ctx, req.ToUser())
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *authUsecase) Login(ctx context.Context, req dto.LoginRequest) (string, error) {
	userDB, err := s.userRepository.FindUserByEmail(ctx, req.Email)
	if errors.Is(err, apperrors.ErrUserNotFound) {
		return "", apperrors.ErrInvalidEmailPass
	}
	if err != nil {
		return "", err
	}

	err = verifyPassword(req.Password, userDB.Password)
	if err != nil {
		return "", apperrors.ErrInvalidEmailPass
	}

	token, err := s.tokenGenerator.GenerateToken(userDB.ID, userDB.IsAdmin)
	if err != nil {
		return "", err
	}
	return token, nil
}

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
