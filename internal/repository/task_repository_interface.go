package repository

import (
	"context"
	"rest-api-codegen/internal/entity"
)

type ITaskRepository interface {
	GetAllTasks(ctx context.Context, userId uint64) ([]*entity.Task, error)
	GetTaskById(ctx context.Context, userId uint64, taskId uint64) (*entity.Task, error)
	CreateTask(ctx context.Context, te *entity.Task) error
	UpdateTask(ctx context.Context, te *entity.Task, userId uint64, taskId uint64) error
	DeleteTask(ctx context.Context, userId uint64, taskId uint64) error
}
