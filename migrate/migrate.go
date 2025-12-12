package main

import (
	"context"
	"fmt"
	"log"
	"rest-api-codegen/internal/repository/db"
)

func main() {
	client := db.NewClient()
	defer fmt.Println("Successfully migrated")
	// nolint: errcheck
	defer client.Close()

	// 自動マイグレーションを実行する。
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
