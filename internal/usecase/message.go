package usecase

import (
	"chat/internal/domain"
)

type Message struct {
	repository     MessageRepository
	userRepository UserRepository
}

func NewMessage(repository MessageRepository, userRepository UserRepository) *Message {
	return &Message{repository: repository, userRepository: userRepository}
}

func (m *Message) GetMessageById(id int64) (*domain.Message, error) {
	return m.repository.GetMessageById(id)
}

func (m *Message) RegisterMessage(message *domain.Message) error {
	if stringIsMeaningless(message.Text) {
		return ErrEmptyMessage
	}
	if message.Time.IsZero() {
		return ErrInvalidTimestamp
	}
	if _, err := m.userRepository.GetUserById(message.FromUserId); err != nil {
		return ErrUserNotFound
	}
	return m.repository.SaveMessage(message)
}

func (m *Message) GetMessages() ([]*domain.Message, error) {
	return m.repository.GetMessages()
}
