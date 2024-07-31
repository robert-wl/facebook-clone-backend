package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/graph/model"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/internal/adapter"
	"time"
)

type MessagesService struct {
	*Service
	MessageAdapter *adapter.MessageAdapter
}

func NewMessagesService(s *Service, ma *adapter.MessageAdapter) *MessagesService {
	return &MessagesService{
		Service:        s,
		MessageAdapter: ma,
	}
}

func (s *MessagesService) CreateConversation(userID string, username string) (*model.Conversation, error) {
	var user *model.User

	if err := s.DB.First(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}
	var conversationUsers []*model.ConversationUsers

	if err := s.DB.Find(&conversationUsers, "user_id in (?)", []string{userID, user.ID}).Error; err != nil {
		return nil, err
	}

	var conversationID = ""

	if err := s.DB.Model(&model.ConversationUsers{}).
		Group("conversation_id").
		Where("user_id in (?)", []string{userID, user.ID}).
		Having("COUNT(user_id) = 2").Pluck("conversation_id", &conversationID).Error; err == nil && conversationID != "" {
		var conversation *model.Conversation

		if err := s.DB.First(&conversation, "id = ?", conversationID).Error; err != nil {
			return nil, err
		}
		return conversation, nil
	}

	conversation := &model.Conversation{
		ID: uuid.NewString(),
	}

	if err := s.DB.Save(&conversation).Error; err != nil {
		return nil, err
	}

	conversationUser1 := &model.ConversationUsers{
		ConversationID: conversation.ID,
		UserID:         userID,
	}

	if err := s.DB.Save(&conversationUser1).Error; err != nil {
		return nil, err
	}

	conversationUser2 := &model.ConversationUsers{
		ConversationID: conversation.ID,
		UserID:         user.ID,
	}

	if err := s.DB.Save(&conversationUser2).Error; err != nil {
		return nil, err
	}

	conversation.Users = append(conversation.Users, conversationUser1)
	conversation.Users = append(conversation.Users, conversationUser2)

	if err := s.DB.Save(&conversation).Error; err != nil {
		return nil, err
	}

	if err := s.DB.
		Preload("Users").
		Preload("Users.User").
		Preload("Messages").
		First(&conversation, "id = ?", conversation.ID).Error; err != nil {
		return nil, err
	}

	return conversation, nil
}

func (s *MessagesService) SendMessage(userID string, conversationID string, message *string, image *string, postID *string) (*model.Message, error) {
	messageModel := &model.Message{
		ID:             uuid.NewString(),
		ConversationID: conversationID,
		SenderID:       userID,
		Message:        message,
		Image:          image,
		PostID:         postID,
		CreatedAt:      time.Now(),
	}

	if err := s.DB.Save(&messageModel).Error; err != nil {
		return nil, err
	}

	go func() {
		for _, convChannel := range s.MessageAdapter.ConversationChannels {
			if convChannel.ConversationID == conversationID {
				var messages []*model.Message

				if err := s.DB.
					Order("created_at desc").
					Preload("Sender").
					Preload("Post").
					Preload("Post.User").
					Find(&messages, "conversation_id = ?", conversationID).Error; err != nil {
					continue
				}

				s.MessageAdapter.BroadcastMessages(messages)
			}
		}
	}()

	if err := s.RedisAdapter.Del([]string{"conversation", conversationID}); err != nil {
		return nil, err
	}

	return messageModel, nil
}

func (s *MessagesService) GetConversations(userID string) ([]*model.Conversation, error) {
	var conversations []*model.Conversation
	var conversationUsers []*string
	var groupIDs []*string

	cacheKey := []string{"conversations", userID}

	err := s.RedisAdapter.GetOrSet(cacheKey, &conversations, func() (interface{}, error) {
		if err := s.DB.Model(&model.Member{}).Where("user_id = ? AND approved = ?", userID, true).Select("group_id").Find(&groupIDs).Error; err != nil {
			return nil, err
		}

		if err := s.DB.Model(&model.ConversationUsers{}).Where("user_id = ?", userID).Select("conversation_id").Find(&conversationUsers).Error; err != nil {
			return nil, err
		}

		if err := s.DB.
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

		if err := s.DB.Find(&groups, "id IN (?)", groupIDs).Error; err != nil {
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

func (s *MessagesService) ViewConversation(ctx context.Context, conversationID string) (<-chan []*model.Message, error) {
	channel := make(chan []*model.Message, 1)

	var message []*model.Message

	err := s.RedisAdapter.GetOrSet([]string{"conversation", conversationID}, &message, func() (interface{}, error) {
		if err := s.DB.First(&model.Conversation{}, "id = ?", conversationID).Error; err != nil {
			return nil, err
		}

		if err := s.DB.
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

	channel <- message

	s.MessageAdapter.AddConversationChannel(channel, conversationID)
	s.MessageAdapter.CloseConversationChannel(ctx, channel, conversationID)
	return channel, nil
}
