package inmemory

import "chat/internal/domain"

type MessageRepository struct{}

func (mr *MessageRepository) SaveMessage(*domain.Message) error {
	//TODO
	return nil
}

func (mr *MessageRepository) GetMessageById(int64) (*domain.Message, error) {
	//TODO
	return nil, nil
}
