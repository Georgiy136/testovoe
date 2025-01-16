package handler

import (
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "myapp/docs"
	"myapp/internal/usecase"
)

func NewRouter(router *gin.Engine, ps usecase.ProjectUseCases) {
	projectHandlers := &ProjectHandler{
		us: ps,
	}

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Routers
	project := router.Group("/currency")
	{
		project.POST("/add", projectHandlers.PostProject)
		project.GET("/price", projectHandlers.GetOneProject)
		project.DELETE("/remove", projectHandlers.DeleteProject)

	}
}
