package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"myapp/config"
	"myapp/internal/handler"
	"myapp/internal/usecase"
	"myapp/internal/usecase/repository"
	"myapp/pkg/postgres"
)

func Run(cfg *config.Config) {
	// Repository
	pg, err := postgres.New(cfg.Postgres)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Up Migrations
	migrateUp(pg)

	// Repo
	projectRepository := repository.NewProject(pg)

	// Use case
	projectUseCases := usecase.NewProjectUsecases(projectRepository)

	// HTTP Server
	router := gin.Default()

	handler.NewRouter(router, *projectUseCases)

	router.Run(fmt.Sprintf(":%d", cfg.Http.Port))
}
