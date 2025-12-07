package controller

import (
	"context"
	"rest-api-codegen/api"
)

type ITaskController interface {
	GetAllTasks(ctx context.Context, req api.GetAllTasksRequestObject) (api.GetAllTasksResponseObject, error)
	CreateTask(ctx context.Context, req api.CreateTaskRequestObject) (api.CreateTaskResponseObject, error)
	DeleteTask(ctx context.Context, req api.DeleteTaskRequestObject) (api.DeleteTaskResponseObject, error)
	GetTaskById(ctx context.Context, req api.GetTaskByIdRequestObject) (api.GetTaskByIdResponseObject, error)
	UpdateTask(ctx context.Context, req api.UpdateTaskRequestObject) (api.UpdateTaskResponseObject, error)
}

type taskController struct{}

func (tc *taskController) GetAllTasks(ctx context.Context, req api.GetAllTasksRequestObject) (api.GetAllTasksResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (tc *taskController) CreateTask(ctx context.Context, req api.CreateTaskRequestObject) (api.CreateTaskResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (tc *taskController) DeleteTask(ctx context.Context, req api.DeleteTaskRequestObject) (api.DeleteTaskResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (tc *taskController) GetTaskById(ctx context.Context, req api.GetTaskByIdRequestObject) (api.GetTaskByIdResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (tc *taskController) UpdateTask(ctx context.Context, req api.UpdateTaskRequestObject) (api.UpdateTaskResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func NewTaskController() ITaskController {
	return &taskController{}
}
