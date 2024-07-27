package model

type Friend struct {
	SenderID   string `json:"senderId" gorm:"primaryKey"`
	Sender     *User  `json:"sender" gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ReceiverID string `json:"receiverId" gorm:"primaryKey"`
	Receiver   *User  `json:"receiver" gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Accepted   bool   `json:"accepted"`
}

type FriendInput struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
}
