package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
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
		log.Fatal().Err(err).Msgf("app - Run - postgres.New: %v", err)
	}
	defer pg.Close()

	// Up Migrations
	migrateUp(pg)

	// Repo
	projectRepository := repository.NewProject(pg)

	// Use case
	projectUseCases := usecase.NewUsecases(projectRepository)

	// HTTP Server
	router := gin.Default()

	handler.NewRouter(router, *projectUseCases)

	router.Run(fmt.Sprintf(":%d", cfg.Http.Port))
}
