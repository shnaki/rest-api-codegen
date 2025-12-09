package repository

import (
	"context"
	"rest-api-codegen/internal/entity"
)

type ITaskRepository interface {
	GetAllTasks(ctx context.Context, userID uint64) ([]*entity.Task, error)
	GetTaskByID(ctx context.Context, userID uint64, taskID uint64) (*entity.Task, error)
	CreateTask(ctx context.Context, te *entity.Task) error
	UpdateTask(ctx context.Context, te *entity.Task, userID uint64, taskID uint64) error
	DeleteTask(ctx context.Context, userID uint64, taskID uint64) error
}
