package main

import "rest-api-codegen/router"

func main() {
	r := router.NewRouter()
	r.Logger.Fatal(r.Start(":1323"))
}
