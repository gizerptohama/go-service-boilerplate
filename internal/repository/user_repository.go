package repository

import (
	"boilerplate/internal/dto"
	"boilerplate/internal/entity"
	apperrors "boilerplate/internal/errors"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	FindUserById(context.Context, uint) (*entity.User, error)
	FindUserByEmail(context.Context, string) (*entity.User, error)
	CreateUser(context.Context, entity.User) (*entity.User, error)
	ResetPassword(context.Context, dto.ResetPassword) (*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindUserById(ctx context.Context, id uint) (*entity.User, error) {
	var result entity.User
	err := r.db.WithContext(ctx).Where("id = ?", id).Find(&result).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperrors.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *userRepository) FindUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var result entity.User
	err := r.db.WithContext(ctx).Table("users").Where("email = ?", email).First(&result).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperrors.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *userRepository) CreateUser(ctx context.Context, user entity.User) (*entity.User, error) {
	var result *entity.User
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) (err error) {
		userDB := entity.User{}
		err = tx.Table("users").Where("email = ?", user.Email).First(&userDB).Error
		if !errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
			return err
		}

		if userDB.Email != "" {
			return apperrors.ErrEmailAlreadyRegistered
		}

		err = tx.Create(&user).Error
		if err != nil {
			return err
		}

		result = &user
		return nil
	})
	return result, err
}

func (r *userRepository) ResetPassword(ctx context.Context, resetPass dto.ResetPassword) (*entity.User, error) {
	var result *entity.User
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) (err error) {
		var entry entity.ForgotPassword
		err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("user_email = ?", resetPass.Email).
			Where("code = ?", resetPass.Code).
			First(&entry).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.ErrForgotPasswordNotFound
		}
		if err != nil {
			return err
		}

		if entry.ExpiredAt.Before(time.Now().UTC()) {
			return apperrors.ErrCodeInvalid
		}

		var user entity.User
		err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("email = ?", resetPass.Email).
			First(&user).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.ErrForgotPasswordNotFound
		}
		if err != nil {
			return err
		}

		user.Password = resetPass.Password
		tx.Save(&user)

		err = tx.Delete(&entry).Error
		if err != nil {
			return err
		}

		result = &user
		return nil
	})
	return result, err
}
