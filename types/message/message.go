package message

import "time"

type Message struct {
	ID          *string   `json:"id" bson:"_id"`
	OwnerID     *string   `json:"owner_id" bson:"owner_id"`
	Content     string    `json:"content" bson:"content"`
	Attachments []string  `json:"attachments" bson:"attachments"`
	ServerID    *string   `json:"server_id" bson:"server_id"`
	ChannelID   string    `json:"channel_id" bson:"channel_id"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}
