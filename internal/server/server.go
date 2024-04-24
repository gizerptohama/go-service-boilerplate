package server

import (
	"boilerplate/internal/handler"
	"boilerplate/internal/middleware"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type RouterOpts struct {
	AuthHandler *handler.AuthHandler
	UserHandler *handler.UserHandler
}

func NewRouter(opts RouterOpts) http.Handler {
	r := gin.New()
	gin.SetMode(os.Getenv("GIN_MODE"))
	r.ContextWithFallback = true

	r.Use(
		middleware.LoggerMiddleware(),
		middleware.ErrorMiddleware(),
		gin.Recovery(),
	)
	registerRoutes(r, opts,
		NewAuthRoute,
	)

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "resource not found",
		})
	})

	return r
}

func registerRoutes(r *gin.Engine, opts RouterOpts, routes ...func(*gin.Engine, RouterOpts)) {
	for _, route := range routes {
		route(r, opts)
	}
}
