package app

import (
	"fmt"

	"log-service/internal/config"
	"log-service/internal/database"
	httpserver "log-service/internal/http"
	"log-service/internal/http/handlers"
	"log-service/internal/logactivity"

	"github.com/gin-gonic/gin"
)

type App struct {
	cfg    config.Config
	router *gin.Engine
}

func New(cfg config.Config) (*App, error) {
	db, err := database.NewPostgres(cfg)
	if err != nil {
		return nil, err
	}

	logActivityRepo := logactivity.NewRepository(db)
	logActivityService := logactivity.NewService(logActivityRepo)
	logActivityHandler := handlers.NewLogActivityHandler(logActivityService)

	router := httpserver.NewRouter(logActivityHandler)

	return &App{
		cfg:    cfg,
		router: router,
	}, nil
}

func (a *App) Run() error {
	address := fmt.Sprintf(":%s", a.cfg.AppPort)

	return a.router.Run(address)
}
