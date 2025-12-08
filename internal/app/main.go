package main

import (
	"rest-api-codegen/internal/controller/rest"
)

func main() {
	r := rest.NewRouter()
	r.Logger.Fatal(r.Start(":1323"))
}
