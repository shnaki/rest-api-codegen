//go:build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
)

func main() {
	if err := entc.Generate("./schema", &gen.Config{
		IDType: &field.TypeInfo{
			Type: field.TypeUint64, // Goの uint64 に対応するentの内部型定数を使用
		},
	}); err != nil {
		log.Fatalf("running entc command: %v", err)
	}
}
