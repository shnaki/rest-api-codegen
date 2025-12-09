package main

import (
	"rest-api-codegen/internal/controller/rest"
)

func main() {
	// REST APIサーバーを開始する。
	r := rest.NewRouter()
	r.Logger.Fatal(r.Start(":1323"))
}
