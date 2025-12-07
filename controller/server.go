package controller

import (
	"context"
	"rest-api-codegen/api"
)

type Server struct {
	uc IUserController
	tc ITaskController
}

func NewServer(uc IUserController, tc ITaskController) *Server {
	return &Server{
		uc: uc,
		tc: tc,
	}
}

func (s *Server) Csrf(ctx context.Context, req api.CsrfRequestObject) (api.CsrfResponseObject, error) {
	return s.uc.Csrf(ctx, req)
}

func (s *Server) Login(ctx context.Context, req api.LoginRequestObject) (api.LoginResponseObject, error) {
	return s.uc.Login(ctx, req)
}

func (s *Server) Logout(ctx context.Context, req api.LogoutRequestObject) (api.LogoutResponseObject, error) {
	return s.uc.Logout(ctx, req)
}

func (s *Server) SignUp(ctx context.Context, req api.SignUpRequestObject) (api.SignUpResponseObject, error) {
	return s.uc.SignUp(ctx, req)
}

func (s *Server) GetAllTasks(ctx context.Context, req api.GetAllTasksRequestObject) (api.GetAllTasksResponseObject, error) {
	return s.tc.GetAllTasks(ctx, req)
}

func (s *Server) CreateTask(ctx context.Context, req api.CreateTaskRequestObject) (api.CreateTaskResponseObject, error) {
	return s.tc.CreateTask(ctx, req)
}

func (s *Server) DeleteTask(ctx context.Context, req api.DeleteTaskRequestObject) (api.DeleteTaskResponseObject, error) {
	return s.tc.DeleteTask(ctx, req)
}

func (s *Server) GetTaskById(ctx context.Context, req api.GetTaskByIdRequestObject) (api.GetTaskByIdResponseObject, error) {
	return s.tc.GetTaskById(ctx, req)
}

func (s *Server) UpdateTask(ctx context.Context, req api.UpdateTaskRequestObject) (api.UpdateTaskResponseObject, error) {
	return s.tc.UpdateTask(ctx, req)
}
