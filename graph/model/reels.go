package model

import "time"

type Reel struct {
	ID           string         `json:"id"`
	UserID       string         `json:"userId"`
	User         *User          `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Content      string         `json:"content"`
	Video        string         `json:"video"`
	LikeCount    int            `json:"likeCount" gorm:"-"`
	CommentCount int            `json:"commentCount" gorm:"-"`
	ShareCount   int            `json:"shareCount"`
	Likes        []*ReelLike    `json:"likes,omitempty" gorm:"foreignKey:ReelID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Comments     []*ReelComment `json:"comments,omitempty" gorm:"foreignKey:ParentReelID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Liked        *bool          `json:"liked,omitempty" gorm:"-"`
	CreatedAt    time.Time      `json:"createdAt"`
}

type ReelComment struct {
	ID              string             `json:"id"`
	UserID          string             `json:"userId"`
	User            *User              `json:"user"`
	Content         string             `json:"content"`
	LikeCount       int                `json:"likeCount" gorm:"-"`
	ReplyCount      int                `json:"replyCount" gorm:"-"`
	ParentReelID    *string            `json:"parentReelId,omitempty"`
	ParentReel      *Reel              `json:"parentReel,omitempty" gorm:"foreignKey:ParentReelID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ParentCommentID *string            `json:"parentCommentId,omitempty"`
	ParentComment   *ReelComment       `json:"parentComment,omitempty" gorm:"foreignKey:ParentCommentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Likes           []*ReelCommentLike `json:"likes,omitempty" gorm:"foreignKey:ReelCommentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Comments        []*ReelComment     `json:"comments,omitempty" gorm:"foreignKey:ParentCommentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Liked           *bool              `json:"liked,omitempty" gorm:"-"`
	CreatedAt       time.Time          `json:"createdAt"`
}

type ReelCommentLike struct {
	ReelCommentID string `json:"reelCommentId" gorm:"primaryKey"`
	UserID        string `json:"userId" gorm:"primaryKey"`
	User          *User  `json:"user"`
}

type ReelLike struct {
	ReelID string `json:"reelId" gorm:"primaryKey"`
	UserID string `json:"userId" gorm:"primaryKey"`
	User   *User  `json:"user"`
}

type NewReel struct {
	Content string `json:"content"`
	Video   string `json:"video"`
}

type NewReelComment struct {
	Content       string  `json:"content"`
	ParentReel    *string `json:"parentReel,omitempty"`
	ParentComment *string `json:"parentComment,omitempty"`
}
