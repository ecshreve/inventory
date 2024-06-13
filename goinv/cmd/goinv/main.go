package main

import (
	"context"
	"goinv"
	"log"
	"net/http"

	"goinv/ent"
	"goinv/ent/migrate"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	_ "github.com/mattn/go-sqlite3"
)

var SQLITE_DB = "file:file.db?mode=rwc&cache=shared&_fk=1"

func main() {
	// Create ent.Client and run the schema migration.
	// client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	client, err := ent.Open(dialect.SQLite, SQLITE_DB)
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	// Configure the server and start listening on :8081.
	srv := handler.NewDefaultServer(goinv.NewSchema(client))
	http.Handle("/",
		playground.Handler("inv", "/query"),
	)
	http.Handle("/query", srv)
	log.Println("listening on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("http server terminated", err)
	}
}
