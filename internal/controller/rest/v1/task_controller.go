package v1

import (
	"context"
)

type ITaskController interface {
	GetAllTasks(ctx context.Context, req GetAllTasksRequestObject) (GetAllTasksResponseObject, error)
	CreateTask(ctx context.Context, req CreateTaskRequestObject) (CreateTaskResponseObject, error)
	DeleteTask(ctx context.Context, req DeleteTaskRequestObject) (DeleteTaskResponseObject, error)
	GetTaskById(ctx context.Context, req GetTaskByIdRequestObject) (GetTaskByIdResponseObject, error)
	UpdateTask(ctx context.Context, req UpdateTaskRequestObject) (UpdateTaskResponseObject, error)
}

type taskController struct{}

func (tc *taskController) GetAllTasks(ctx context.Context, req GetAllTasksRequestObject) (GetAllTasksResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (tc *taskController) CreateTask(ctx context.Context, req CreateTaskRequestObject) (CreateTaskResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (tc *taskController) DeleteTask(ctx context.Context, req DeleteTaskRequestObject) (DeleteTaskResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (tc *taskController) GetTaskById(ctx context.Context, req GetTaskByIdRequestObject) (GetTaskByIdResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (tc *taskController) UpdateTask(ctx context.Context, req UpdateTaskRequestObject) (UpdateTaskResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func NewTaskController() ITaskController {
	return &taskController{}
}
