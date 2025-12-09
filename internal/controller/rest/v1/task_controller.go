package v1

import (
	"context"
	"net/http"
	"rest-api-codegen/internal/controller/rest/middleware/jwt"
	"rest-api-codegen/internal/entity"
	"rest-api-codegen/internal/usecase"
)

type ITaskController interface {
	GetAllTasks(ctx context.Context, req GetAllTasksRequestObject) (GetAllTasksResponseObject, error)
	CreateTask(ctx context.Context, req CreateTaskRequestObject) (CreateTaskResponseObject, error)
	DeleteTask(ctx context.Context, req DeleteTaskRequestObject) (DeleteTaskResponseObject, error)
	GetTaskById(ctx context.Context, req GetTaskByIDRequestObject) (GetTaskByIDResponseObject, error)
	UpdateTask(ctx context.Context, req UpdateTaskRequestObject) (UpdateTaskResponseObject, error)
}

type taskController struct {
	tu usecase.ITaskUsecase
}

func NewTaskController(tu usecase.ITaskUsecase) ITaskController {
	return &taskController{
		tu: tu,
	}
}

func (tc *taskController) GetAllTasks(ctx context.Context, req GetAllTasksRequestObject) (GetAllTasksResponseObject, error) {
	userId := jwt.GetUserIdFromContext(ctx)
	tasks, err := tc.tu.GetAllTasks(userId)
	if err != nil {
		errMessage := err.Error()
		return GetAllTasks500ApplicationProblemPlusJSONResponse(ProblemDetails{
			Type:   "about:blank",
			Title:  "Internal Server Error",
			Status: http.StatusInternalServerError,
			Detail: &errMessage,
		}), nil
	}
	resTasks := make([]TaskResponse, 0, len(tasks))
	for _, t := range tasks {
		var taskResponse TaskResponse
		fromEntityToTaskResponse(t, &taskResponse)
		resTasks = append(resTasks, taskResponse)
	}
	return GetAllTasks200JSONResponse(resTasks), nil
}

func (tc *taskController) CreateTask(ctx context.Context, req CreateTaskRequestObject) (CreateTaskResponseObject, error) {
	userId := jwt.GetUserIdFromContext(ctx)
	t := entity.Task{
		Title:  req.Body.Title,
		UserId: userId,
	}
	if err := tc.tu.CreateTask(&t); err != nil {
		errMessage := err.Error()
		return CreateTask500ApplicationProblemPlusJSONResponse(ProblemDetails{
			Type:   "about:blank",
			Title:  "Internal Server Error",
			Status: http.StatusInternalServerError,
			Detail: &errMessage,
		}), nil
	}
	var taskResponse TaskResponse
	fromEntityToTaskResponse(&t, &taskResponse)
	return CreateTask201JSONResponse(taskResponse), nil
}

func (tc *taskController) GetTaskById(ctx context.Context, req GetTaskByIDRequestObject) (GetTaskByIDResponseObject, error) {
	userId := jwt.GetUserIdFromContext(ctx)
	t, err := tc.tu.GetTaskById(userId, req.TaskID)
	if err != nil {
		errMessage := err.Error()
		return GetTaskByID500ApplicationProblemPlusJSONResponse(ProblemDetails{
			Type:   "about:blank",
			Title:  "Internal Server Error",
			Status: http.StatusInternalServerError,
			Detail: &errMessage,
		}), nil
	}
	var taskResponse TaskResponse
	fromEntityToTaskResponse(t, &taskResponse)
	return GetTaskByID200JSONResponse(taskResponse), nil
}

func (tc *taskController) UpdateTask(ctx context.Context, req UpdateTaskRequestObject) (UpdateTaskResponseObject, error) {
	userId := jwt.GetUserIdFromContext(ctx)
	t := entity.Task{
		Title: req.Body.Title,
	}
	if err := tc.tu.UpdateTask(&t, userId, req.TaskID); err != nil {
		errMessage := err.Error()
		return UpdateTask500ApplicationProblemPlusJSONResponse(ProblemDetails{
			Type:   "about:blank",
			Title:  "Internal Server Error",
			Status: http.StatusInternalServerError,
			Detail: &errMessage,
		}), nil
	}
	var taskResponse TaskResponse
	fromEntityToTaskResponse(&t, &taskResponse)
	return UpdateTask200JSONResponse(taskResponse), nil
}

func (tc *taskController) DeleteTask(ctx context.Context, req DeleteTaskRequestObject) (DeleteTaskResponseObject, error) {
	userId := jwt.GetUserIdFromContext(ctx)
	if err := tc.tu.DeleteTask(userId, req.TaskID); err != nil {
		errMessage := err.Error()
		return DeleteTask500ApplicationProblemPlusJSONResponse(ProblemDetails{
			Type:   "about:blank",
			Title:  "Internal Server Error",
			Status: http.StatusInternalServerError,
			Detail: &errMessage,
		}), nil
	}
	return DeleteTask204Response{}, nil
}
