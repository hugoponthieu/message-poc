package service

import (
	"message/repository"
	"message/repository/types"
	"message/types/message"
)

type MessageService struct {
	repository repository.MessageRepository
}

func NewMessageService(repository repository.MessageRepository) MessageService {
	return MessageService{
		repository: repository,
	}
}

func (s MessageService) Get(id string) (*message.Message, error) {
	return s.repository.Get(id)
}

func (s MessageService) Mget(ids []string) (*types.MGetResult, error) {
	return s.repository.MGet(ids)
}

func (s MessageService) Create(message *message.Message) (*message.Message, error) {
	return s.repository.Create(message)
}

func (s MessageService) Update(id string, message *message.UpdateMessage) (*message.Message, error) {
	return s.repository.Update(id, message)
}

func (s MessageService) Delete(id string) error {
	return s.repository.Delete(id)
}

func (s MessageService) GetByChannel(channelID string, page, limit int) (*[]message.Message, error) {
	return s.repository.GetByChannel(channelID, page, limit)
}

func (s MessageService) Search(query string, channelId *string, serverID *string, page, limit int) (*[]message.Message, error) {
	return s.repository.Search(query, channelId, serverID, page, limit)
}
