package v1

import (
	"context"
)

type IUserController interface {
	Csrf(ctx context.Context, req CsrfRequestObject) (CsrfResponseObject, error)
	Login(ctx context.Context, req LoginRequestObject) (LoginResponseObject, error)
	Logout(ctx context.Context, req LogoutRequestObject) (LogoutResponseObject, error)
	SignUp(ctx context.Context, req SignUpRequestObject) (SignUpResponseObject, error)
}

type userController struct{}

func NewUserController() IUserController {
	return &userController{}
}

func (uc *userController) Csrf(ctx context.Context, req CsrfRequestObject) (CsrfResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *userController) Login(ctx context.Context, req LoginRequestObject) (LoginResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *userController) Logout(ctx context.Context, req LogoutRequestObject) (LogoutResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *userController) SignUp(ctx context.Context, req SignUpRequestObject) (SignUpResponseObject, error) {
	//TODO implement me
	panic("implement me")
}
