package models

import (
	"time"

	"github.com/google/uuid"
)

type PostStatus string

const (
	PostStatusDraft     PostStatus = "draft"
	PostStatusPublished PostStatus = "published"
)

type Post struct {
	ID        uuid.UUID  `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Title     string     `json:"title"`
	Content   string     `json:"content" gorm:"type:text"`
	Status    PostStatus `json:"status"`
	CreatedAt time.Time  `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"default:current_timestamp"`
}
