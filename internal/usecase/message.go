package usecase

import (
	"chat/internal/domain"
	"unicode"
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

func stringIsMeaningless(s string) bool {
	for _, c := range s {
		if !unicode.IsSpace(c) {
			return false
		}
	}
	return true
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
