package main

import (
	"context"
	"net/http"
	"os"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/shanna/entxid-test/ent"
	"github.com/shanna/entxid-test/graph"
	"go.uber.org/zap"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/shanna/entxid-test/ent/runtime"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	log, _ := zap.NewDevelopment()
	client, err := ent.Open(
		"sqlite3",
		"file:ent?mode=memory&cache=shared&_fk=1",
	)
	client = client.Debug()
	if err != nil {
		log.Fatal("opening ent client", zap.Error(err))
	}
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal("running schema migration", zap.Error(err))
	}

	srv := handler.NewDefaultServer(graph.NewSchema(client))
	srv.Use(entgql.Transactioner{TxOpener: client})
	if true { // TODO: Some debug flag. cli.Debug {
		srv.Use(&debug.Tracer{})
	}

	http.Handle("/",
		playground.Handler("Woot", "/query"),
	)
	http.Handle("/query", srv)

	log.Info("listening on", zap.String("address", ":"+port))
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Error("http server terminated", zap.Error(err))
	}
}
