package main

import (
	"context"
	"log"
	"net/http"

	"github.com/yahkerobertkertasnya/facebook-clone-backend/graph/directives"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/internal/adapter"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/internal/auth"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/internal/database"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/internal/services"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/internal/utils"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	cors2 "github.com/rs/cors"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/graph"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/graph/resolver"
)

const defaultPort = "8080"

func main() {

	port := utils.GetDotENVVariable("PORT", defaultPort)

	router := chi.NewRouter()

	cors := cors2.New(cors2.Options{
		AllowedOrigins:   []string{"http://localhost", "http://localhost:5173", "http://localhost:8080", "chrome-extension://flnheeellpciglgpaodhkhmapeljopja"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		Debug:            true,
	})

	router.Use(cors.Handler)
	router.Use(auth.AuthMiddleware)

	service := services.NewService(database.GetDBInstance(), adapter.NewRedisCacheAdapter())
	messageAdapter := adapter.NewMessageAdapter()

	userService := services.NewUserService(service)
	reelsService := services.NewReelsService(service)
	notificationService := services.NewNotificationService(service)
	messagesService := services.NewMessagesService(service, messageAdapter)
	postService := services.NewPostService(service, notificationService, messagesService)
	storyService := services.NewStoryService(service, notificationService)
	groupService := services.NewGroupService(service, notificationService)
	friendsService := services.NewFriendsService(service, notificationService)

	c := graph.Config{Resolvers: &resolver.Resolver{
		UserService:         userService,
		StoryService:        storyService,
		ReelsService:        reelsService,
		PostService:         postService,
		NotificationService: notificationService,
		MessagesService:     messagesService,
		GroupService:        groupService,
		FriendsService:      friendsService,
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

	//database.DropDatabase()
	//database.MigrateDatabase()
	//database.FakeData()

	adapter.NewRedisCacheAdapter().DelAll()

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	srv.SetQueryCache(lru.New(1000))
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
