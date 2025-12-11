package usecase

import (
	"context"
	"errors"
	"rest-api-codegen/internal/entity"
	"rest-api-codegen/internal/repository"
)

//go:generate go tool mockgen -source=$GOFILE -destination=../../mock/usecase/mock/mock_$GOFILE -package=usecasemock

var ErrTaskAlreadyExists = errors.New("task already exists")

type ITaskUsecase interface {
	GetAllTasks(userID uint64) ([]*entity.Task, error)
	GetTaskByID(userID uint64, taskID uint64) (*entity.Task, error)
	CreateTask(task *entity.Task) error
	UpdateTask(task *entity.Task, userID uint64, taskID uint64) error
	DeleteTask(userID uint64, taskID uint64) error
}

type taskUsecase struct {
	tr repository.TaskRepository
}

func NewTaskUsecase(tr repository.TaskRepository) ITaskUsecase {
	return &taskUsecase{tr: tr}
}

func (tu *taskUsecase) GetAllTasks(userID uint64) ([]*entity.Task, error) {
	return tu.tr.GetAllTasks(context.Background(), userID)
}

func (tu *taskUsecase) GetTaskByID(userID uint64, taskID uint64) (*entity.Task, error) {
	return tu.tr.GetTaskByID(context.Background(), userID, taskID)
}

func (tu *taskUsecase) CreateTask(task *entity.Task) error {
	return tu.tr.CreateTask(context.Background(), task)
}

func (tu *taskUsecase) UpdateTask(task *entity.Task, userID uint64, taskID uint64) error {
	return tu.tr.UpdateTask(context.Background(), task, userID, taskID)
}

func (tu *taskUsecase) DeleteTask(userID uint64, taskID uint64) error {
	return tu.tr.DeleteTask(context.Background(), userID, taskID)
}
