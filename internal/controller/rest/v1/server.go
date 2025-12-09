package v1

import (
	"context"
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

func (s *Server) Csrf(ctx context.Context, req CsrfRequestObject) (CsrfResponseObject, error) {
	return s.uc.Csrf(ctx, req)
}

func (s *Server) Login(ctx context.Context, req LoginRequestObject) (LoginResponseObject, error) {
	return s.uc.Login(ctx, req)
}

func (s *Server) Logout(ctx context.Context, req LogoutRequestObject) (LogoutResponseObject, error) {
	return s.uc.Logout(ctx, req)
}

func (s *Server) SignUp(ctx context.Context, req SignUpRequestObject) (SignUpResponseObject, error) {
	return s.uc.SignUp(ctx, req)
}

func (s *Server) GetAllTasks(ctx context.Context, req GetAllTasksRequestObject) (GetAllTasksResponseObject, error) {
	return s.tc.GetAllTasks(ctx, req)
}

func (s *Server) CreateTask(ctx context.Context, req CreateTaskRequestObject) (CreateTaskResponseObject, error) {
	return s.tc.CreateTask(ctx, req)
}

func (s *Server) DeleteTask(ctx context.Context, req DeleteTaskRequestObject) (DeleteTaskResponseObject, error) {
	return s.tc.DeleteTask(ctx, req)
}

func (s *Server) GetTaskByID(ctx context.Context, req GetTaskByIDRequestObject) (GetTaskByIDResponseObject, error) {
	return s.tc.GetTaskById(ctx, req)
}

func (s *Server) UpdateTask(ctx context.Context, req UpdateTaskRequestObject) (UpdateTaskResponseObject, error) {
	return s.tc.UpdateTask(ctx, req)
}
