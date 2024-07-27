package model

import (
	"time"
)

type NewUser struct {
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Dob       time.Time `json:"dob"`
	Gender    string    `json:"gender"`
}

type User struct {
	ID                string    `json:"id"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	Username          string    `json:"username" gorm:"unique"`
	Email             string    `json:"email" gorm:"unique"`
	Password          string    `json:"password"`
	Dob               time.Time `json:"dob"`
	Gender            string    `json:"gender"`
	Active            bool      `json:"active"`
	MiscId            *string   `json:"miscId,omitempty"`
	Profile           *string   `json:"profile,omitempty"`
	Background        *string   `json:"background,omitempty"`
	CreatedAt         time.Time `json:"createdAt"`
	Posts             []*Post   `json:"posts,omitempty" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	FriendCount       int       `json:"friendCount" gorm:"-"`
	MutualCount       int       `json:"mutualCount" gorm:"-"`
	NotificationCount int       `json:"notificationCount" gorm:"-"`
	Friended          string    `json:"friended" gorm:"-"`
	Theme             string    `json:"theme"`
}

type UpdateUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
	Gender    string `json:"gender"`
}
