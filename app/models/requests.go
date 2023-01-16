package models

import (
	"time"
)

// Database structures

type Request struct {
	ID         string    `gorm:"type:uuid;primary_key" json:"id"`
	SenderID   string    `gorm:"uniqueIndex:request_index" json:"senderID"`
	Sender     User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"sender" validate:"omitempty"`
	ReceiverID string    `gorm:"uniqueIndex:request_index" json:"receiverID" validate:"required,min=2,max=64"`
	Receiver   User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"receiver" validate:"omitempty"`
	CreatedAt  time.Time `gorm:"autoCreateTime:false" json:"createdAt"`
}

// Output structures

type PrivateRequestWithSender struct {
	ID       string                    `json:"id"`
	SenderID string                    `json:"senderID"`
	Sender   *PublicUserWithoutFriends `json:"sender"`
}

type PrivateRequestWithReceiver struct {
	ID         string                    `json:"id"`
	ReceiverID string                    `json:"receiverID"`
	Receiver   *PublicUserWithoutFriends `json:"receiver"`
}
