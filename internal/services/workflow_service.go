package services

import (
	"context"
	"relay-engine/internal/models"
	"relay-engine/internal/repositories"

	"github.com/google/uuid"
)

type WorkflowService struct {
	repository repositories.WorkflowRepository
}

func NewWorkflowService(repository repositories.WorkflowRepository) *WorkflowService {
	return &WorkflowService{repository: repository}
}

func (s *WorkflowService) Create(ctx context.Context, req *models.WorkflowDefinition) (*models.WorkflowDefinition, error) {
	req.ID = uuid.New()

	err := s.repository.Create(ctx, *req)
	if err != nil {
		return nil, err
	}

	return req, nil
}
