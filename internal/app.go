package app

import (
	"boilerplate/internal/constants"
	"boilerplate/internal/database"
	"boilerplate/internal/server"
	"boilerplate/internal/utils/logger"
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

type App struct{}

func NewApp() (app *App) {
	app = &App{}
	return
}

func (app *App) Run() {
	// setup logger
	log, logFile := logger.NewLogger(true)
	if logFile != nil {
		defer logFile.Close()
	}
	logger.SetLogger(log)

	// retrive environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// connecting to DB
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("unable to connect to the database: %v", err)
		return
	}
	defer database.CloseDB(db)

	// Comment the block of code right below this comment
	// if you want to keep the data
	database.Migrate(db)
	database.Seed(db)

	// Setup JWT
	constants.SetupJWTEnv()

	// Setup Logic Env
	constants.SetupLogicEnv()

	// setup router
	opts := server.RegisterHandlers(db)
	router := server.NewRouter(*opts)
	srv := http.Server{
		Addr:    os.Getenv("APP_PORT"),
		Handler: router,
	}

	// Graceful Shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
