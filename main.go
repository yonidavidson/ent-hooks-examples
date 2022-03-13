package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/yonidavidson/ent-side-effect-hooks-example/cloud"
	"github.com/yonidavidson/ent-side-effect-hooks-example/ent"
	_ "github.com/yonidavidson/ent-side-effect-hooks-example/ent/runtime"
)

func main() {
	ctx := context.Background()
	cs := cloud.NewSyncer()
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", ent.CloudSyncer(cs))
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	cs.Start(ctx, client)
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	cl := client.Cloud.Create().SetWalks(-1).SaveX(ctx)
	d := client.Dog.Create().SetName("Karashindo").SaveX(ctx)
	u := client.User.Create().SetName("Yoni").SetCloud(cl).AddPets(d).SaveX(ctx)
	tx, err := client.Tx(ctx)
	tx.Dog.UpdateOne(d).SetName("Fortuna").ExecX(ctx)
	tx.User.UpdateOne(u).SetName("Gany").ExecX(ctx)
	if err := tx.Commit(); err != nil {
		fmt.Printf("failed to update dog and user")
		return
	}
	cs.Close()
	cl = u.QueryCloud().OnlyX(ctx)
	fmt.Println(cl)
}
