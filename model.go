package user

import (
	"time"
)

type Contact struct {
	// ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" `
	ID int `json:"id" bson:"id"`
	PhoneNumber    string    `json:"phoneNumber" bson:"phoneNumber"`
	Email          string    `json:"email" bson:"email"`
	LinkedID       int64     `json:"linkedId" bson:"linkedId"`
	LinkPrecedence string    `json:"linkPrecedence" bson:"linkPrecedence"`
	CreatedAt      time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt" bson:"updatedAt"`
	DeletedAt      time.Time `json:"deletedAt" bson:"deletedAt"`
}

