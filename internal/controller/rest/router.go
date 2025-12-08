package rest

import (
	"rest-api-codegen/internal/controller/rest/v1"

	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	handler := v1.NewStrictHandler(v1.NewServer(), nil)
	v1.RegisterHandlersWithBaseURL(e, handler, "")
	return e
}
