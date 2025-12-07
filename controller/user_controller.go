package controller

import (
	"context"
	"rest-api-codegen/api"
)

type IUserController interface {
	Csrf(ctx context.Context, req api.CsrfRequestObject) (api.CsrfResponseObject, error)
	Login(ctx context.Context, req api.LoginRequestObject) (api.LoginResponseObject, error)
	Logout(ctx context.Context, req api.LogoutRequestObject) (api.LogoutResponseObject, error)
	SignUp(ctx context.Context, req api.SignUpRequestObject) (api.SignUpResponseObject, error)
}

type userController struct{}

func NewUserController() IUserController {
	return &userController{}
}

func (uc *userController) Csrf(ctx context.Context, req api.CsrfRequestObject) (api.CsrfResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *userController) Login(ctx context.Context, req api.LoginRequestObject) (api.LoginResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *userController) Logout(ctx context.Context, req api.LogoutRequestObject) (api.LogoutResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *userController) SignUp(ctx context.Context, req api.SignUpRequestObject) (api.SignUpResponseObject, error) {
	//TODO implement me
	panic("implement me")
}
