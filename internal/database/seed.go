package database

import (
	"boilerplate/internal/entity"
	"boilerplate/internal/repository"
	"boilerplate/internal/utils"
	"context"
	"log"

	"gorm.io/gorm"
)

var (
	user entity.User = entity.User{
		UserName: "Bob",
		Password: "password",
		Email:    "bob@mail.com",
	}
	admin entity.User = entity.User{
		UserName: "admin",
		Password: "password",
		Email:    "admin@mail.com",
		IsAdmin:  true,
	}

	users []entity.User = []entity.User{
		user, admin,
	}
)

func Seed(db *gorm.DB) {
	for i, user := range users {
		var err error

		users[i].Password, err = utils.HashPassword(user.Password)
		if err != nil {
			log.Println("error seeding when hashing", err)
		}

		userRepository := repository.NewUserRepository(db)
		_, err = userRepository.CreateUser(context.Background(), users[i])
		if err != nil {
			log.Println("error seeding data users :", err)
		}
	}

}
