package worker

import (
	"context"
	"encoding/json"
	"log"
	"relay-engine/internal/models"
)

type PostHandlerInterface interface {
	Publish(ctx context.Context, payload []byte) error
}
type PostHandler struct{}

func NewPostHandler() *PostHandler {
	return &PostHandler{}
}

func (h *PostHandler) Publish(ctx context.Context, payload []byte) error {
	var req models.Post
	if err := json.Unmarshal(payload, &req); err != nil {
		return err
	}

	log.Println("Publish post id: ", req.ID.String())

	return nil
}
