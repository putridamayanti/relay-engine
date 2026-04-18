package router

import (
	"relay-engine/internal/database"
	"relay-engine/internal/handlers/http"
	"relay-engine/internal/repositories"
	"relay-engine/internal/services"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.RedirectTrailingSlash = false

	db := database.ConnectDatabase()

	workflowRepo := repositories.NewWorkflowRepository(db)
	workflowService := services.NewWorkflowService(*workflowRepo)
	workflowHandler := http.NewWorkflowHandler(*workflowService)

	executionRepo := repositories.NewExecutionRepository(db)
	executionService := services.NewExecutionService(*executionRepo)
	executionHandler := http.NewExecutionHandler(*executionService)

	api := router.Group("/api/v1")

	workflow := api.Group("/workflows")
	{
		workflow.POST("", workflowHandler.Create)
	}

	execution := api.Group("/executions")
	{
		execution.POST("/start", executionHandler.Start)
	}

	return router
}
