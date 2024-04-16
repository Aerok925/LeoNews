package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Aerok925/LeoNews/bootstrap"
	"github.com/Aerok925/LeoNews/configs"
	"github.com/Aerok925/LeoNews/internal/graph"
	"github.com/Aerok925/LeoNews/internal/repository/postgres"
	"github.com/Aerok925/LeoNews/internal/service"
	"log"
	"net/http"
)

func main() {
	cfg, err := configs.Init()
	if err != nil {
		log.Println(err)
		return
	}
	db, err := bootstrap.NewDB(cfg.Database)
	if err != nil {
		panic(err.Error())
		log.Fatal(err.Error())
	}
	repos := postgres.Init(db)
	service.Init(repos)
	defer db.Close()
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", 8080)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
