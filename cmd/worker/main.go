package main

import (
	"context"
	"os"
	"relay-engine/internal/database"
	"relay-engine/internal/handlers/worker"
	"relay-engine/internal/repositories"
	"relay-engine/internal/services"
)

func main() {
	db := database.ConnectDatabase()

	taskRepo := repositories.NewTaskRepository(db)

	registry := worker.NewActivityRegistry()

	postHandler := worker.NewPostHandler()

	registry.Register("publish_post", postHandler.Publish)

	workerId := os.Getenv("WORKER_ID")

	workerService := services.NewWorkerService(taskRepo, workerId, 5, registry)

	ctx := context.Background()
	workerService.Start(ctx)
}
