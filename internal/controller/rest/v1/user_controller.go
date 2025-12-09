package v1

import (
	"context"
	"errors"
	"net/http"
	"os"
	"rest-api-codegen/internal/entity"
	"rest-api-codegen/internal/usecase"
	"time"
)

const CookieNameToken = "token"

type IUserController interface {
	Csrf(ctx context.Context, req CsrfRequestObject) (CsrfResponseObject, error)
	Login(ctx context.Context, req LoginRequestObject) (LoginResponseObject, error)
	Logout(ctx context.Context, req LogoutRequestObject) (LogoutResponseObject, error)
	SignUp(ctx context.Context, req SignUpRequestObject) (SignUpResponseObject, error)
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{
		uu: uu,
	}
}

func (uc *userController) SignUp(ctx context.Context, req SignUpRequestObject) (SignUpResponseObject, error) {
	user := entity.User{}
	user.Email = req.Body.Email
	user.Password = req.Body.Password
	newUser, err := uc.uu.Signup(user)
	errMessage := err.Error()
	if errors.Is(err, usecase.ErrUserAlreadyExists) {
		return SignUp409ApplicationProblemPlusJSONResponse(ProblemDetails{
			Type:   "about:blank",
			Title:  "Resource Conflict",
			Status: http.StatusConflict,
			Detail: &errMessage,
		}), nil
	} else if err != nil {
		return SignUp500ApplicationProblemPlusJSONResponse(ProblemDetails{
			Type:   "about:blank",
			Title:  "Internal Server Error",
			Status: http.StatusInternalServerError,
			Detail: &errMessage,
		}), nil
	}
	userRes := UserResponse{
		Email: newUser.Email,
		Id:    newUser.ID,
	}
	return SignUp201JSONResponse(userRes), nil
}

func (uc *userController) Login(ctx context.Context, req LoginRequestObject) (LoginResponseObject, error) {
	user := entity.User{}
	user.Email = req.Body.Email
	user.Password = req.Body.Password
	tokenString, err := uc.uu.Login(user)
	if err != nil {
		errMessage := err.Error()
		return Login500ApplicationProblemPlusJSONResponse(ProblemDetails{
			Type:   "about:blank",
			Title:  "Internal Server Error",
			Status: http.StatusInternalServerError,
			Detail: &errMessage,
		}), nil
	}

	// トークンをクッキーに設定しておく。
	cookie := new(http.Cookie)
	cookie.Name = CookieNameToken
	cookie.Value = tokenString
	// クッキーの有効期限は24時間としておく。
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	cookieHeader := cookie.String()
	response := Login200Response{
		Headers: Login200ResponseHeaders{
			SetCookie: cookieHeader,
		},
	}
	return response, nil
}

func (uc *userController) Logout(ctx context.Context, req LogoutRequestObject) (LogoutResponseObject, error) {
	cookie := new(http.Cookie)
	cookie.Name = CookieNameToken
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	cookieHeader := cookie.String()
	response := Logout200Response{
		Headers: Logout200ResponseHeaders{
			SetCookie: cookieHeader,
		},
	}
	return response, nil
}

func (uc *userController) Csrf(ctx context.Context, req CsrfRequestObject) (CsrfResponseObject, error) {
	//TODO implement me
	panic("implement me")
}
