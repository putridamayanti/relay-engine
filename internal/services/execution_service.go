package services

import (
	"context"
	"relay-engine/internal/models"
	"relay-engine/internal/repositories"
	"time"

	"github.com/google/uuid"
)

type ExecutionService struct {
	repository repositories.ExecutionRepository
}

func NewExecutionService(repository repositories.ExecutionRepository) *ExecutionService {
	return &ExecutionService{repository: repository}
}

func (s *ExecutionService) Start(ctx context.Context, req *models.WorkflowExecution) (*models.WorkflowExecution, error) {
	req.ID = uuid.New()
	req.Status = models.WorkflowExecutionStatusRunning

	now := time.Now()
	req.StartedAt = &now

	err := s.repository.Create(ctx, *req)
	if err != nil {
		return nil, err
	}

	return req, nil
}
