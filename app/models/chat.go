package models

import (
	"time"
)

// Database structures

type Message struct {
	ID         string    `gorm:"type:uuid;primary_key" json:"id"`
	SenderID   string    `gorm:"uniqueIndex:request_index" json:"senderID"`
	Sender     User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"sender" validate:"omitempty"`
	ReceiverID string    `gorm:"uniqueIndex:request_index" json:"receiverID" validate:"required,min=2,max=64"`
	Receiver   User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"receiver" validate:"omitempty"`
	Text       string    `gorm:"not null;size:512;" json:"text" validate:"required,min=1,max=512"`
	CreatedAt  time.Time `gorm:"autoCreateTime:false" json:"createdAt"`
}

// Output structures

type MessageWithSender struct {
	ID        string                    `json:"id"`
	SenderID  string                    `json:"senderID"`
	Sender    *PublicUserWithoutFriends `json:"sender"`
	Text      string                    `json:"text"`
	CreatedAt time.Time                 `json:"createdAt"`
}

type MessageWithoutSender struct {
	ID        string    `json:"id"`
	SenderID  string    `json:"senderID"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}

type Chat struct {
	SenderID   string                    `json:"senderID"`
	ReceiverID string                    `json:"receiverID"`
	Text       string                    `json:"text"`
	Sender     *PublicUserWithoutFriends `json:"sender"`
	Receiver   *PublicUserWithoutFriends `json:"receiver"`
	CreatedAt  time.Time                 `json:"createdAt"`
}
