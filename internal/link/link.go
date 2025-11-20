package link

import (
	"errors"
	"time"
)

type Link struct {
	URL       string    `bson:"url" json:"url"`
	ShortCode string    `bson:"shortCode" json:"short_code"`
	Id        string    `bson:"_id,omitempty" json:"id"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updated_at"`
	CreatedAt time.Time `bson:"createdAt" json:"created_at"`
}

type ShortenLinkRequest struct {
	URL string `json:"url" binding:"required,url"`
}

var ErrNotFound = errors.New("record not found")
