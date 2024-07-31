package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/graph"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/graph/model"
)

// CreateConversation is the resolver for the createConversation field.
func (r *mutationResolver) CreateConversation(ctx context.Context, username string) (*model.Conversation, error) {
	var user *model.User

	if err := r.DB.First(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}

	userIDCtx := ctx.Value("UserID").(string)
	var conversationUsers []*model.ConversationUsers

	if err := r.DB.Find(&conversationUsers, "user_id in (?)", []string{userIDCtx, user.ID}).Error; err != nil {
		return nil, err
	}

	var conversationID string = ""

	if err := r.DB.Model(&model.ConversationUsers{}).
		Group("conversation_id").
		Where("user_id in (?)", []string{userIDCtx, user.ID}).
		Having("COUNT(user_id) = 2").Pluck("conversation_id", &conversationID).Error; err == nil && conversationID != "" {
		var conversation *model.Conversation

		if err := r.DB.First(&conversation, "id = ?", conversationID).Error; err != nil {
			return nil, err
		}
		return conversation, nil
	}

	conversation := &model.Conversation{
		ID: uuid.NewString(),
	}

	if err := r.DB.Save(&conversation).Error; err != nil {
		return nil, err
	}

	conversationUser1 := &model.ConversationUsers{
		ConversationID: conversation.ID,
		UserID:         userIDCtx,
	}

	if err := r.DB.Save(&conversationUser1).Error; err != nil {
		return nil, err
	}

	conversationUser2 := &model.ConversationUsers{
		ConversationID: conversation.ID,
		UserID:         user.ID,
	}

	if err := r.DB.Save(&conversationUser2).Error; err != nil {
		return nil, err
	}

	conversation.Users = append(conversation.Users, conversationUser1)
	conversation.Users = append(conversation.Users, conversationUser2)

	if err := r.DB.Save(&conversation).Error; err != nil {
		return nil, err
	}

	if err := r.DB.
		Preload("Users").
		Preload("Users.User").
		Preload("Messages").
		First(&conversation, "id = ?", conversation.ID).Error; err != nil {
		return nil, err
	}

	return conversation, nil
}

// SendMessage is the resolver for the sendMessage field.
func (r *mutationResolver) SendMessage(ctx context.Context, conversationID string, message *string, image *string, postID *string) (*model.Message, error) {
	userIDCtx := ctx.Value("UserID").(string)

	messageModel := &model.Message{
		ID:             uuid.NewString(),
		ConversationID: conversationID,
		SenderID:       userIDCtx,
		Message:        message,
		Image:          image,
		PostID:         postID,
		CreatedAt:      time.Now(),
	}

	if err := r.DB.Save(&messageModel).Error; err != nil {
		return nil, err
	}

	go func() {
		for _, convChannel := range r.ConversationChannels {
			if convChannel.ConversationID == conversationID {
				var messages []*model.Message

				if err := r.DB.
					Order("created_at desc").
					Preload("Sender").
					Preload("Post").
					Preload("Post.User").
					Find(&messages, "conversation_id = ?", conversationID).Error; err != nil {
					continue
				}

				convChannel.Channel <- messages
			}
		}
	}()

	r.RedisAdapter.Del([]string{"conversation", conversationID})

	return messageModel, nil
}

// GetConversations is the resolver for the getConversations field.
func (r *queryResolver) GetConversations(ctx context.Context) ([]*model.Conversation, error) {
	var conversations []*model.Conversation
	var conversationUsers []*string
	var groupIDs []*string

	userID := ctx.Value("UserID").(string)

	cacheKey := []string{"conversations", userID}

	err := r.RedisAdapter.GetOrSet(cacheKey, &conversations, func() (interface{}, error) {
		if err := r.DB.Model(&model.Member{}).Where("user_id = ? AND approved = ?", userID, true).Select("group_id").Find(&groupIDs).Error; err != nil {
			return nil, err
		}

		if err := r.DB.Model(&model.ConversationUsers{}).Where("user_id = ?", userID).Select("conversation_id").Find(&conversationUsers).Error; err != nil {
			return nil, err
		}

		if err := r.DB.
			Preload("Users").
			Preload("Users.User").
			Preload("Messages").
			Find(&conversations, "id in (?) OR group_id IN (?)", conversationUsers, groupIDs).Error; err != nil {
			return nil, err
		}

		if len(conversations) == 0 {
			return conversations, nil
		}

		var groupIds []string
		for _, conversation := range conversations {
			if conversation.GroupID == nil {
				continue
			}
			groupIds = append(groupIds, *conversation.GroupID)
		}

		var groups []*model.Group

		if err := r.DB.Find(&groups, "id IN (?)", groupIDs).Error; err != nil {
			return nil, err
		}

		for _, group := range groups {
			for _, conversation := range conversations {
				if conversation.GroupID == &group.ID {
					conversation.Group = group
				}
			}
		}

		return conversations, nil
	}, 10*time.Minute)

	if err != nil {
		return nil, err
	}

	return conversations, nil
}

// ViewConversation is the resolver for the viewConversation field.
func (r *subscriptionResolver) ViewConversation(ctx context.Context, conversationID string) (<-chan []*model.Message, error) {
	channel := make(chan []*model.Message, 1)

	var message []*model.Message

	err := r.RedisAdapter.GetOrSet([]string{"conversation", conversationID}, &message, func() (interface{}, error) {
		if err := r.DB.First(&model.Conversation{}, "id = ?", conversationID).Error; err != nil {
			return nil, err
		}

		if err := r.DB.
			Order("created_at DESC").
			Preload("Sender").
			Preload("Post").
			Preload("Post.User").
			Find(&message, "conversation_id = ?", conversationID).Error; err != nil {
			return nil, err
		}

		return message, nil
	}, 10*time.Minute)

	if err != nil {
		close(channel)
		return nil, err
	}

	go func() {
		<-ctx.Done()
		var convChannel []*model.ConversationChannel

		for _, conv := range r.ConversationChannels {
			if conv.ConversationID != conversationID {
				convChannel = append(convChannel, conv)
			}
		}

		r.ConversationChannels = convChannel

		close(channel)
	}()

	channel <- message

	convChannel := &model.ConversationChannel{
		Channel:        channel,
		ConversationID: conversationID,
	}

	r.ConversationChannels = append(r.ConversationChannels, convChannel)

	return channel, nil
}

// Subscription returns graph.SubscriptionResolver implementation.
func (r *Resolver) Subscription() graph.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
