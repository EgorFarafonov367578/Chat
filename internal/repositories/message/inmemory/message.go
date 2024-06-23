package inmemory

import (
	"chat/internal/domain"
	"chat/internal/usecase"
	"sync"
)

type MessageRepository struct {
	list  []*domain.Message
	mutex sync.Mutex
}

func NewMessageRepository() *MessageRepository {
	return &MessageRepository{list: make([]*domain.Message, 0)}
}

func (mr *MessageRepository) SaveMessage(message *domain.Message) error {
	mr.mutex.Lock()
	defer mr.mutex.Unlock()
	message.Id = int64(len(mr.list))
	mr.list = append(mr.list, message)
	return nil
}

func (mr *MessageRepository) GetMessageById(id int64) (*domain.Message, error) {
	mr.mutex.Lock()
	defer mr.mutex.Unlock()
	if id < int64(len(mr.list)) {
		return mr.list[id], nil
	}
	return nil, usecase.ErrMessageNotFound
}

func (mr *MessageRepository) GetMessages() ([]*domain.Message, error) {
	mr.mutex.Lock()
	defer mr.mutex.Unlock()
	ans := make([]*domain.Message, len(mr.list))
	copy(ans, mr.list)
	return ans, nil
}
