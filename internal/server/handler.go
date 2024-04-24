package server

import (
	"boilerplate/internal/constants"
	"boilerplate/internal/handler"
	"boilerplate/internal/repository"
	"boilerplate/internal/usecase"
	"boilerplate/internal/utils"

	"gorm.io/gorm"
)

var userRepository repository.UserRepository

func RegisterHandlers(db *gorm.DB) *RouterOpts {
	return &RouterOpts{
		registerAuthHandler(db),
		registerUserHandler(db),
	}
}

func registerAuthHandler(db *gorm.DB) *handler.AuthHandler {
	userRepository = repository.NewUserRepository(db)
	tokenGenerator := utils.NewJWTTokenGenerator(constants.JWT_Secret)
	authUsecase := usecase.NewAuthUsecase(userRepository, tokenGenerator)
	return handler.NewAuthHandler(authUsecase)
}

func registerUserHandler(db *gorm.DB) *handler.UserHandler {
	userRepository = repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	return handler.NewUserHandler(userUsecase)
}
