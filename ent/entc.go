//go:build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
)

func main() {
	opts := []entc.Option{
		entc.Dependency(
			entc.DependencyName("CacheSyncer"),
			entc.DependencyTypeInfo(&field.TypeInfo{
				Ident:   "hook.Syncer",
				PkgPath: "github.com/yonidavidson/ent-hooks-examples/hook",
			}),
		),
	}
	if err := entc.Generate("./schema", &gen.Config{}, opts...); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
