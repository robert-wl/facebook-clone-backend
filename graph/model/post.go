package model

import "time"

type Post struct {
	ID           string            `json:"id"`
	UserID       string            `json:"userId"`
	User         *User             `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Content      string            `json:"content"`
	Privacy      string            `json:"privacy"`
	Visibility   []*PostVisibility `json:"visibility,omitempty" gorm:"foreignKey:PostID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PostTags     []*PostTag        `json:"postTags,omitempty" gorm:"foreignKey:PostID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	LikeCount    int               `json:"likeCount" gorm:"-"`
	CommentCount int               `json:"commentCount" gorm:"-"`
	ShareCount   int               `json:"shareCount"`
	GroupID      *string           `json:"groupId,omitempty"`
	Group        *Group            `json:"group,omitempty" gorm:"foreignKey:GroupID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Files        []*string         `json:"files,omitempty" gorm:"json"`
	Likes        []*PostLike       `json:"likes,omitempty" gorm:"foreignKey:PostID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Comments     []*Comment        `json:"comments,omitempty" gorm:"foreignKey:ParentPostID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Liked        *bool             `json:"liked,omitempty" gorm:"-"`
	CreatedAt    time.Time         `json:"createdAt"`
}

type Comment struct {
	ID              string         `json:"id"`
	UserID          string         `json:"userId"`
	User            *User          `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Content         string         `json:"content"`
	LikeCount       int            `json:"likeCount" gorm:"-"`
	ReplyCount      int            `json:"replyCount" gorm:"-"`
	ParentPostID    *string        `json:"parentPostId,omitempty"`
	ParentPost      *Post          `json:"parentPost,omitempty" gorm:"foreignKey:ParentPostID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ParentCommentID *string        `json:"parentCommentId,omitempty"`
	ParentComment   *Comment       `json:"parentComment,omitempty" gorm:"foreignKey:ParentCommentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Likes           []*CommentLike `json:"likes,omitempty" gorm:"foreignKey:CommentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Comments        []*Comment     `json:"comments,omitempty" gorm:"foreignKey:ParentCommentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Liked           *bool          `json:"liked,omitempty" gorm:"-"`
	CreatedAt       time.Time      `json:"createdAt"`
}

type PostTag struct {
	PostID string `json:"postId" gorm:"primaryKey;foreignKey:PostID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserID string `json:"userId" gorm:"primaryKey"`
	User   *User  `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type PostVisibility struct {
	PostID string `json:"postId" gorm:"primaryKey;foreignKey:PostID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserID string `json:"userId" gorm:"primaryKey"`
	User   *User  `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type CommentLike struct {
	CommentID string `json:"commentId" gorm:"primaryKey;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserID    string `json:"userId" gorm:"primaryKey;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User      *User  `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type NewPost struct {
	Content    string    `json:"content"`
	Privacy    string    `json:"privacy"`
	Files      []*string `json:"files,omitempty"`
	GroupID    *string   `json:"groupId,omitempty"`
	Visibility []*string `json:"visibility,omitempty"`
	Tags       []*string `json:"tags,omitempty"`
}

type NewComment struct {
	Content       string  `json:"content"`
	ParentPost    *string `json:"parentPost,omitempty"`
	ParentComment *string `json:"parentComment,omitempty"`
}

type PostLike struct {
	PostID string `json:"postId" gorm:"primaryKey;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserID string `json:"userId" gorm:"primaryKey;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User   *User  `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
