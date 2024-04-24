package usecase

import (
	"boilerplate/internal/entity"
	apperrors "boilerplate/internal/errors"
	"boilerplate/internal/repository"
	"context"
	"errors"
)

type UserUsecase interface {
	GetUserDataById(context.Context, uint) (*entity.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: r,
	}
}

func (s *userUsecase) GetUserDataById(ctx context.Context, id uint) (*entity.User, error) {
	result, err := s.userRepository.FindUserById(ctx, id)
	if errors.Is(err, apperrors.ErrUserNotFound) {
		return nil, apperrors.ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	return result, nil

}
