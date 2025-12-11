package repository

import (
	"context"
	"rest-api-codegen/internal/entity"
)

//go:generate go tool mockgen -source=$GOFILE -destination=../../mock/repository/mock/mock_task_repository.go -package=repositorymock

type TaskRepository interface {
	GetAllTasks(ctx context.Context, userID uint64) ([]*entity.Task, error)
	GetTaskByID(ctx context.Context, userID uint64, taskID uint64) (*entity.Task, error)
	CreateTask(ctx context.Context, te *entity.Task) error
	UpdateTask(ctx context.Context, te *entity.Task, userID uint64, taskID uint64) error
	DeleteTask(ctx context.Context, userID uint64, taskID uint64) error
}
