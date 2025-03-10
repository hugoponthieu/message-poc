package repository

import (
	"message/repository/types"
	"message/types/message"
)

type MessageRepository interface {
	Get(id string) (*message.Message, error)
	MGet(ids []string) (*types.MGetResult, error)
	Create(msg *message.Message) (*message.Message, error)
	Update(id string, msg *message.UpdateMessage) (*message.Message, error)
	Delete(id string) error
	GetByChannel(channelId string, page int, limit int) (*[]message.Message, error)
	Search(query string, channelId *string, serverId *string, page int, limit int) (*[]message.Message, error)
}

