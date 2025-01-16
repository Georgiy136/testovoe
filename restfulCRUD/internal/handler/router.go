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
	project := router.Group("/project")
	{
		project.POST("/", projectHandlers.PostProject)
		project.GET("/", projectHandlers.GetAllProjects)
		project.GET("/:id", projectHandlers.GetOneProject)
		project.PUT("/:id", projectHandlers.UpdateProject)
		project.DELETE("/:id", projectHandlers.DeleteProject)

	}
}
