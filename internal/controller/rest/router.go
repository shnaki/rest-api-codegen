package rest

import (
	"rest-api-codegen/internal/controller/rest/v1"
	"rest-api-codegen/internal/repository/db"
	"rest-api-codegen/internal/usecase"

	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	client := db.NewClient()
	ur := db.NewUserRepository(client)
	uu := usecase.NewUserUsecase(ur)
	uc := v1.NewUserController(uu)
	tc := v1.NewTaskController()
	handler := v1.NewStrictHandler(v1.NewServer(uc, tc), nil)

	v1.RegisterHandlersWithBaseURL(e, handler, "/api/v1")
	return e
}
