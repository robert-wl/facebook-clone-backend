package model

import "time"

type Group struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	About       string        `json:"about"`
	Privacy     string        `json:"privacy"`
	Background  string        `json:"background"`
	Members     []*Member     `json:"members" gorm:"foreignKey:GroupID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MemberCount int           `json:"memberCount" gorm:"-"`
	Joined      string        `json:"joined" gorm:"-"`
	IsAdmin     bool          `json:"isAdmin" gorm:"-"`
	ChatID      *string       `json:"chatId,omitempty"`
	Chat        *Conversation `json:"chat" gorm:"foreignKey:ChatID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Posts       []*Post       `json:"posts,omitempty" gorm:"foreignKey:GroupID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Files       []*GroupFile  `json:"Files,omitempty" gorm:"foreignKey:GroupID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time     `json:"createdAt"`
}

type Member struct {
	GroupID   string `json:"groupId" gorm:"primaryKey;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    string `json:"userId" gorm:"primaryKey;foreignKey:ConversationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User      *User  `json:"user"`
	Requested bool   `json:"requested"`
	Approved  bool   `json:"approved"`
	Role      string `json:"role"`
}

type GroupFile struct {
	ID         string    `json:"id"`
	GroupID    string    `json:"groupID" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	URL        string    `json:"url"`
	UserID     string    `json:"userId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UploadedBy *User     `json:"uploadedBy" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UploadedAt time.Time `json:"uploadedAt"`
}

type NewGroup struct {
	Name    string `json:"name"`
	About   string `json:"about"`
	Privacy string `json:"privacy"`
}

type NewGroupFile struct {
	Name string `json:"name"`
	Type string `json:"type"`
	URL  string `json:"url"`
}
