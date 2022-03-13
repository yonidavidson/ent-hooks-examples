package main

import (
	"context"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/yonidavidson/ent-side-effect-hooks-example/cloud"
	"github.com/yonidavidson/ent-side-effect-hooks-example/ent"
)

func main() {
	ctx := context.Background()
	cs := cloud.NewSyncer()
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", ent.CloudSyncer(cs))
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	cs.Start(ctx, client)
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	cs.Close()
}
