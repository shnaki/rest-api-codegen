package router

import (
	"rest-api-codegen/api"
	"rest-api-codegen/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	handler := api.NewStrictHandler(controller.NewServer(), nil)
	api.RegisterHandlersWithBaseURL(e, handler, "")
	return e
}
