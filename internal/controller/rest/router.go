package rest

import (
	"fmt"
	"os"
	"rest-api-codegen/internal/controller/rest/middleware/jwt"
	"rest-api-codegen/internal/controller/rest/v1"
	"rest-api-codegen/internal/repository/db"
	"rest-api-codegen/internal/usecase"
	"strings"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// 認証を必要としない正確なパスのリストを定義する。
// GET, POST, PUTなどのHTTPメソッドも区別できるようにマップ構造にする。
// パスは/api/v1等のapiBasePrefix以降を記載する。
var publicEndpointsRelative = map[string]bool{
	"POST /signup": true,
	"POST /login":  true,
	"POST /logout": true,
	"GET /csrf":    true,
}

func NewRouter() *echo.Echo {
	e := echo.New()

	client := db.NewClient()
	ur := db.NewUserRepository(client)
	tr := db.NewTaskRepository(client)
	uu := usecase.NewUserUsecase(ur)
	tu := usecase.NewTaskUsecase(tr)
	uc := v1.NewUserController(uu)
	tc := v1.NewTaskController(tu)

	// ミドルウェアを設定する。
	apiBasePrefix := "/api/v1"

	// 認証が必要なパスにはJWTを適用する。
	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:     []byte(os.Getenv("SECRET")),
		TokenLookup:    "cookie:token",
		SuccessHandler: jwt.SuccessHandler,
		Skipper: func(c echo.Context) bool {
			reqMethod := c.Request().Method
			reqPath := c.Request().URL.Path

			// リクエストパスが設定されたAPIBasePrefixで始まっているか確認する。
			if !strings.HasPrefix(reqPath, apiBasePrefix) {
				// APIパス以外（例: /adminコンソールなど）はスキップしない（JWTミドルウェアが全体にかかっている前提なので）。
				// もしAPIパス以外はJWTチェック不要なら return true に変更する。
				return false
			}

			// ベースパス以降の相対パス部分を抽出する。
			relativePath := strings.TrimPrefix(reqPath, apiBasePrefix)

			// "METHOD /relativePath" の形式のキーを作成する。
			key := fmt.Sprintf("%s %s", reqMethod, relativePath)

			// publicEndpointsRelative マップに存在すればスキップ（認証不要）する。
			if _, exists := publicEndpointsRelative[key]; exists {
				return true
			}

			// 存在しなければスキップしない（認証が必要）。
			return false
		},
	}))

	handler := v1.NewStrictHandler(v1.NewServer(uc, tc), nil)

	v1.RegisterHandlersWithBaseURL(e, handler, apiBasePrefix)
	return e
}
