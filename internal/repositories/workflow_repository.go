package repositories

import (
	"context"
	"relay-engine/internal/models"

	"gorm.io/gorm"
)

type WorkflowRepository struct {
	db *gorm.DB
}

func NewWorkflowRepository(db *gorm.DB) *WorkflowRepository {
	return &WorkflowRepository{db: db}
}

func (r *WorkflowRepository) Create(ctx context.Context, req models.WorkflowDefinition) error {
	return r.db.WithContext(ctx).Create(req).Error
}

func (r *WorkflowRepository) FindByID(ctx context.Context, id string) (*models.WorkflowDefinition, error) {
	var item models.WorkflowDefinition
	err := r.db.WithContext(ctx).First(&item, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &item, nil
}
