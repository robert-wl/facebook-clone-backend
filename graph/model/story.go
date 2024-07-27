package model

import "time"

type Story struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	User      *User     `json:"user"`
	Image     *string   `json:"image,omitempty"`
	Text      *string   `json:"text,omitempty"`
	Font      *string   `json:"font,omitempty"`
	Color     *string   `json:"color,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}

type NewImageStory struct {
	Image string `json:"image"`
}

type NewTextStory struct {
	Text  string `json:"text"`
	Font  string `json:"font"`
	Color string `json:"color"`
}
