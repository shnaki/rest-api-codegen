package v1

import (
	"context"
	"errors"
	"strings"
	"testing"

	"rest-api-codegen/internal/entity"
	"rest-api-codegen/internal/usecase"
	usecasemock "rest-api-codegen/mock/usecase/mock"

	"go.uber.org/mock/gomock"
)

func TestUserController_Login_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	uu := usecasemock.NewMockIUserUsecase(ctrl)
	uc := NewUserController(uu)

	body := UserRequest{Email: "a@b.c", Password: "pw"}
	req := LoginRequestObject{Body: &body}
	uu.EXPECT().Login(entity.User{Email: body.Email, Password: body.Password}).Return("tok123", nil)

	t.Setenv("API_DOMAIN", "example.com")

	resp, err := uc.Login(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	r, ok := resp.(Login200Response)
	if !ok {
		t.Fatalf("expected Login200Response, got %T", resp)
	}
	if !strings.Contains(r.Headers.SetCookie, "token=tok123") || !strings.Contains(r.Headers.SetCookie, "Domain=example.com") || !strings.Contains(r.Headers.SetCookie, "HttpOnly") || !strings.Contains(r.Headers.SetCookie, "Secure") || !strings.Contains(r.Headers.SetCookie, "SameSite=None") {
		t.Fatalf("unexpected Set-Cookie: %s", r.Headers.SetCookie)
	}
}

func TestUserController_Login_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	uu := usecasemock.NewMockIUserUsecase(ctrl)
	uc := NewUserController(uu)

	body := UserRequest{Email: "a@b.c", Password: "pw"}
	req := LoginRequestObject{Body: &body}
	uu.EXPECT().Login(entity.User{Email: body.Email, Password: body.Password}).Return("", errors.New("boom"))

	resp, err := uc.Login(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if _, ok := resp.(Login500ApplicationProblemPlusJSONResponse); !ok {
		t.Fatalf("expected Login500ApplicationProblemPlusJSONResponse, got %T", resp)
	}
}

func TestUserController_SignUp_Conflict(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	uu := usecasemock.NewMockIUserUsecase(ctrl)
	uc := NewUserController(uu)

	body := UserRequest{Email: "a@b.c", Password: "pw"}
	req := SignUpRequestObject{Body: &body}
	uu.EXPECT().Signup(entity.User{Email: body.Email, Password: body.Password}).Return(entity.User{}, usecase.ErrUserAlreadyExists)

	resp, err := uc.SignUp(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if _, ok := resp.(SignUp409ApplicationProblemPlusJSONResponse); !ok {
		t.Fatalf("expected SignUp409..., got %T", resp)
	}
}

func TestUserController_SignUp_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	uu := usecasemock.NewMockIUserUsecase(ctrl)
	uc := NewUserController(uu)

	body := UserRequest{Email: "a@b.c", Password: "pw"}
	req := SignUpRequestObject{Body: &body}
	uu.EXPECT().Signup(entity.User{Email: body.Email, Password: body.Password}).Return(entity.User{}, errors.New("db"))

	resp, err := uc.SignUp(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if _, ok := resp.(SignUp500ApplicationProblemPlusJSONResponse); !ok {
		t.Fatalf("expected SignUp500..., got %T", resp)
	}
}

func TestUserController_Logout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	uu := usecasemock.NewMockIUserUsecase(ctrl)
	uc := NewUserController(uu)

	t.Setenv("API_DOMAIN", "example.com")
	resp, err := uc.Logout(context.Background(), LogoutRequestObject{})
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	r, ok := resp.(Logout200Response)
	if !ok {
		t.Fatalf("expected Logout200Response, got %T", resp)
	}
	if !strings.Contains(r.Headers.SetCookie, "token=") || !strings.Contains(r.Headers.SetCookie, "Domain=example.com") {
		t.Fatalf("unexpected Set-Cookie: %s", r.Headers.SetCookie)
	}
}
