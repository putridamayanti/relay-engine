package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type WorkflowExecutionStatus string
type ActivityTaskStatus string

const (
	WorkflowExecutionStatusPending   WorkflowExecutionStatus = "pending"
	WorkflowExecutionStatusRunning   WorkflowExecutionStatus = "running"
	WorkflowExecutionStatusCompleted WorkflowExecutionStatus = "completed"
	WorkflowExecutionStatusFailed    WorkflowExecutionStatus = "failed"

	ActivityTaskStatusPending ActivityTaskStatus = "pending"
	ActivityTaskStatusRunning ActivityTaskStatus = "running"
	ActivityTaskStatusSuccess ActivityTaskStatus = "success"
	ActivityTaskStatusFailed  ActivityTaskStatus = "failed"
)

type WorkflowDefinition struct {
	ID         uuid.UUID       `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name       string          `json:"name" gorm:"unique;not null"`
	Version    int             `json:"version"`
	Definition json.RawMessage `json:"definition" gorm:"type:jsonb;not null"`
	CreatedAt  time.Time       `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt  time.Time       `json:"updated_at" gorm:"default:current_timestamp"`
}

type WorkflowExecution struct {
	ID          uuid.UUID               `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	WorkflowID  string                  `json:"workflow_id" gorm:"not null"`
	Status      WorkflowExecutionStatus `json:"status" gorm:"not null;default:pending"`
	Input       json.RawMessage         `json:"input" gorm:"type:jsonb"`
	Output      json.RawMessage         `json:"output" gorm:"type:jsonb"`
	CurrentStep string                  `json:"current_step"`
	StartedAt   *time.Time              `json:"started_at"`
	CompletedAt *time.Time              `json:"completed_at"`
}

type ActivityTask struct {
	ID          uuid.UUID          `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	ExecutionID string             `json:"execution_id" gorm:"not null"`
	StepName    string             `json:"step_name"`
	Payload     json.RawMessage    `json:"payload" gorm:"type:jsonb"`
	RetryCount  int                `json:"retry_count" gorm:"default:0"`
	ScheduledAt *time.Time         `json:"scheduled_at"`
	LockedAt    *time.Time         `json:"locked_at"`
	WorkerID    string             `json:"worker_id" gorm:"not null;default:worker-1"`
	LastError   *string            `json:"last_error" gorm:"type:text"`
	Status      ActivityTaskStatus `json:"status" gorm:"not null;default:pending"`
	CreatedAt   *time.Time         `json:"created_at" gorm:"default:current_timestamp"`
}

type Event struct {
	ID          uuid.UUID       `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	ExecutionID string          `json:"execution_id" gorm:"not null"`
	Type        string          `json:"type" gorm:"not null"`
	Payload     json.RawMessage `json:"payload" gorm:"type:jsonb"`
}
