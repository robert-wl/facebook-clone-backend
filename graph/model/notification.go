package model

import "time"

type Notification struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	UserID    string    `json:"userId"`
	User      *User     `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	SenderID  string    `json:"senderId"`
	Sender    *User     `json:"sender" gorm:"foreignKey:SenderID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Seen      bool      `json:"seen"`
	PostID    *string   `json:"postId,omitempty"`
	Post      *Post     `json:"post,omitempty" gorm:"foreignKey:PostID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ReelID    *string   `json:"reelId,omitempty"`
	Reel      *Reel     `json:"reel,omitempty" gorm:"foreignKey:ReelID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	StoryID   *string   `json:"storyId,omitempty"`
	Story     *Story    `json:"story,omitempty" gorm:"foreignKey:StoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	GroupID   *string   `json:"groupId,omitempty"`
	Group     *Group    `json:"group,omitempty" gorm:"foreignKey:GroupID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time `json:"createdAt"`
}

type BlockNotification struct {
	SenderID   string `json:"senderId" gorm:"primaryKey"`
	Sender     *User  `json:"sender" gorm:"foreignKey:SenderID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ReceiverID string `json:"receiverId" gorm:"primaryKey"`
	Receiver   *User  `json:"receiver" gorm:"foreignKey:ReceiverID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type NewNotification struct {
	Message string  `json:"message"`
	UserID  string  `json:"userId"`
	PostID  *string `json:"postId,omitempty"`
	ReelID  *string `json:"reelId,omitempty"`
	StoryID *string `json:"storyId,omitempty"`
	GroupID *string `json:"groupId,omitempty"`
}
