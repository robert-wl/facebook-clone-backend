package model

import "time"

type Conversation struct {
	ID       string               `json:"id"`
	Users    []*ConversationUsers `json:"users" gorm:"foreignKey:ConversationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Messages []*Message           `json:"messages,omitempty" gorm:"foreignKey:ConversationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Message struct {
	ID             string    `json:"id"`
	ConversationID string    `json:"conversationId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SenderID       string    `json:"senderId"`
	Sender         *User     `json:"sender" gorm:"foreignKey:SenderID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Message        *string   `json:"message,omitempty"`
	Image          *string   `json:"image,omitempty"`
	PostID         *string   `json:"postId,omitempty"`
	Post           *Post     `json:"post,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt      time.Time `json:"createdAt"`
}

type ConversationUsers struct {
	ConversationID string `json:"conversationId" gorm:"primaryKey;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID         string `json:"userId" gorm:"primaryKey"`
	User           *User  `json:"user"`
}

type ConversationChannel struct {
	Channel        chan []*Message
	ConversationID string
}
