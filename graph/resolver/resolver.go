package resolver

import (
	"github.com/yahkerobertkertasnya/facebook-clone-backend/internal/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	UserService         *services.UserService
	StoryService        *services.StoryService
	ReelsService        *services.ReelsService
	PostService         *services.PostService
	NotificationService *services.NotificationService
	MessagesService     *services.MessagesService
	GroupService        *services.GroupService
	FriendsService      *services.FriendsService
}
