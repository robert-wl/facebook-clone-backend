package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	cors2 "github.com/rs/cors"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/database/postgresql"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/database/redis"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/graph"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/graph/resolver"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/helper"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/helper/directives"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/middleware"
	"log"
	"net/http"
)

const defaultPort = "8080"

func main() {

	port := helper.GetDotENVVariable("PORT", defaultPort)

	router := chi.NewRouter()

	cors := cors2.New(cors2.Options{
		AllowedOrigins:   []string{"http://localhost", "http://localhost:5173", "http://localhost:8080", "chrome-extension://flnheeellpciglgpaodhkhmapeljopja"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		Debug:            true,
	})

	router.Use(cors.Handler)
	router.Use(middleware.AuthMiddleware)

	c := graph.Config{Resolvers: &resolver.Resolver{
		DB:    postgresql.GetInstance(),
		Redis: redis.GetInstance(),
	}}

	c.Directives.Auth = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		return directives.AuthDirectives(ctx, next)
	}

	srv := handler.New(graph.NewExecutableSchema(c))

	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	srv.SetQueryCache(lru.New(1000))
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})
	//srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	postgresql.MigrateDatabase()

	//helper.SendVerification("robert.wiliam12345@gmail.com", "12345")

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
