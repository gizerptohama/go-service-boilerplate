package server

import (
	"boilerplate/internal/middleware"

	"github.com/gin-gonic/gin"
)

func NewAuthRoute(r *gin.Engine, opts RouterOpts) {
	auth := r.Group("/auth")
	auth.POST("/register", opts.AuthHandler.Register)
	auth.POST("/login", opts.AuthHandler.Login)
}

func NewUserRoute(r *gin.Engine, opts RouterOpts) {
	user := r.Group("/users")
	user.GET("/profile", middleware.AuthMiddleware(false), opts.UserHandler.GetUserData)
}
