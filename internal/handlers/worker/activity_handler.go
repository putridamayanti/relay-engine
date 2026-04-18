package worker

import (
	"context"
	"errors"
	"relay-engine/internal/models"
)

type ActivityHandler func(ctx context.Context, payload []byte) error

type ActivityRegistry struct {
	handlers map[string]ActivityHandler
}

func NewActivityRegistry() *ActivityRegistry {
	return &ActivityRegistry{handlers: make(map[string]ActivityHandler)}
}

func (ar *ActivityRegistry) Register(name string, handler ActivityHandler) {
	ar.handlers[name] = handler
}

func (ar *ActivityRegistry) Execute(ctx context.Context, task models.ActivityTask) error {
	handler, ok := ar.handlers[task.StepName]
	if !ok {
		return errors.New("unknown task step")
	}

	return handler(ctx, task.Payload)
}
