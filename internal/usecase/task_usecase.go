package usecase

import (
	"context"
	"errors"
	"rest-api-codegen/internal/entity"
	"rest-api-codegen/internal/repository"
)

var ErrTaskAlreadyExists = errors.New("task already exists")

type ITaskUsecase interface {
	GetAllTasks(userId uint64) ([]*entity.Task, error)
	GetTaskById(userId uint64, taskId uint64) (*entity.Task, error)
	CreateTask(task *entity.Task) error
	UpdateTask(task *entity.Task, userId uint64, taskId uint64) error
	DeleteTask(userId uint64, taskId uint64) error
}

type taskUsecase struct {
	tr repository.ITaskRepository
}

func NewTaskUsecase(tr repository.ITaskRepository) ITaskUsecase {
	return &taskUsecase{tr: tr}
}

func (tu *taskUsecase) GetAllTasks(userId uint64) ([]*entity.Task, error) {
	return tu.tr.GetAllTasks(context.Background(), userId)
}

func (tu *taskUsecase) GetTaskById(userId uint64, taskId uint64) (*entity.Task, error) {
	return tu.tr.GetTaskById(context.Background(), userId, taskId)
}

func (tu *taskUsecase) CreateTask(task *entity.Task) error {
	return tu.tr.CreateTask(context.Background(), task)
}

func (tu *taskUsecase) UpdateTask(task *entity.Task, userId uint64, taskId uint64) error {
	return tu.tr.UpdateTask(context.Background(), task, userId, taskId)
}

func (tu *taskUsecase) DeleteTask(userId uint64, taskId uint64) error {
	return tu.tr.DeleteTask(context.Background(), userId, taskId)
}
