package main

import (
	"github.com/rs/zerolog/log"
	"myapp/config"
	"myapp/internal/app"
)

//	@title			Swagger API
//	@version		1.0
//	@description	Swagger API for Golang Project

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

//	@host		localhost:8080
//	@BasePath	/

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal().Err(err).Msgf("Config error: %s:", err)
	}
	// Run
	app.Run(cfg)
}
