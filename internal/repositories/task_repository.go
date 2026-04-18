package repositories

import (
	"context"
	"relay-engine/internal/models"
	"time"

	"gorm.io/gorm"
)

type TaskRepositoryInterface interface {
	FetchAndLock(ctx context.Context, workerId string, limit int) ([]models.ActivityTask, error)
	Update(ctx context.Context, req models.ActivityTask) error
}

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(ctx context.Context, req models.ActivityTask) error {
	return r.db.WithContext(ctx).Create(&req).Error
}

func (r *TaskRepository) FetchAndLock(ctx context.Context, workerId string, limit int) ([]models.ActivityTask, error) {
	var tasks []models.ActivityTask

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Raw(`
			SELECT * FROM activity_tasks
			WHERE status = 'pending'
			AND scheduled_at <= NOW()
			ORDER BY scheduled_at DESC
			FOR UPDATE SKIP LOCKED
			LIMIT ?
		`, limit).Scan(&tasks).Error; err != nil {
			return err
		}

		now := time.Now()

		for i, _ := range tasks {
			tasks[i].LockedAt = &now
			tasks[i].Status = models.ActivityTaskStatusRunning
			tasks[i].WorkerID = workerId

			if err := tx.Save(&tasks[i]).Error; err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskRepository) Update(ctx context.Context, req models.ActivityTask) error {
	return r.db.WithContext(ctx).Save(&req).Error
}
