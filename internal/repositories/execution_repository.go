package repositories

import (
	"context"
	"relay-engine/internal/models"

	"gorm.io/gorm"
)

type ExecutionRepository struct {
	db *gorm.DB
}

func NewExecutionRepository(db *gorm.DB) *ExecutionRepository {
	return &ExecutionRepository{
		db: db,
	}
}

func (r *ExecutionRepository) Create(ctx context.Context, req models.WorkflowExecution) error {
	return r.db.WithContext(ctx).Create(&req).Error
}

func (r *ExecutionRepository) FindByID(ctx context.Context, id string) (*models.WorkflowExecution, error) {
	var item models.WorkflowExecution
	err := r.db.WithContext(ctx).First(&item, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *ExecutionRepository) Update(ctx context.Context, req *models.WorkflowExecution) error {
	return r.db.WithContext(ctx).Save(&req).Error
}
