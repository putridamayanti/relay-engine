package services

import (
	"context"
	"log"
	"relay-engine/internal/handlers/worker"
	"relay-engine/internal/models"
	"relay-engine/internal/repositories"
	"sync"
	"time"
)

type WorkerService struct {
	repo        repositories.TaskRepositoryInterface
	workerID    string
	concurrency int
	registry    *worker.ActivityRegistry
}

func NewWorkerService(
	repo repositories.TaskRepositoryInterface,
	workerID string,
	concurrency int,
	registry *worker.ActivityRegistry) *WorkerService {
	return &WorkerService{repo: repo, workerID: workerID, concurrency: concurrency, registry: registry}
}

func (s *WorkerService) Start(ctx context.Context) {
	semaphore := make(chan struct{}, s.concurrency)

	for {
		select {
		case <-ctx.Done():
			log.Println("worker stopped")
			return
		default:
		}

		tasks, err := s.repo.FetchAndLock(ctx, s.workerID, s.concurrency)
		if err != nil {
			log.Println(err)
			time.Sleep(5 * time.Second)
			continue
		}

		if len(tasks) == 0 {
			time.Sleep(5 * time.Second)
			continue
		}

		var wg sync.WaitGroup
		for _, task := range tasks {
			wg.Add(1)
			semaphore <- struct{}{}

			go func() {
				defer wg.Done()
				defer func() { <-semaphore }()

				s.ProcessTask(ctx, &task)
			}()
		}

		wg.Wait()
	}
}

func (s *WorkerService) ProcessTask(ctx context.Context, task *models.ActivityTask) {
	err := s.registry.Execute(ctx, *task)
	if err != nil {
		log.Println("Task Failed: ", task.ID, err)

		task.RetryCount++

		if task.RetryCount > 3 {
			task.Status = models.ActivityTaskStatusFailed
		} else {
			task.Status = models.ActivityTaskStatusPending
			scheduledAt := time.Now().Add(time.Duration(task.RetryCount) * time.Second)
			task.ScheduledAt = &scheduledAt

			errorStr := err.Error()
			task.LastError = &errorStr

			err = s.repo.Update(ctx, *task)
			if err != nil {
				log.Println("Update Task Failed: ", task.ID, err)
			}
			return
		}
	}

	err = s.repo.Update(ctx, *task)
	if err != nil {
		log.Println("Update Task Failed: ", task.ID, err)
		return
	}
}
