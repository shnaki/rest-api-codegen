package jwt

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type contextKey string

const userContextKey contextKey = "userID"

func SuccessHandler(c echo.Context) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId, ok := claims["user_id"]

	if ok {
		// 標準の http.Request の context に userId をセットする。
		ctx := context.WithValue(c.Request().Context(), userContextKey, userId)
		// 更新されたコンテキストを持つリクエストで echo.Context を更新する。
		c.SetRequest(c.Request().WithContext(ctx))
	}
}

func GetUserIdFromContext(ctx context.Context) uint64 {
	userId := ctx.Value(userContextKey)
	return uint64(userId.(float64))
}
