package usecase

import (
	"chat/internal/domain"
	"errors"
)

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrInvalidUserName  = errors.New("invalid user name")
	ErrEmptyMessage     = errors.New("empty message text")
	ErrInvalidTimestamp = errors.New("invalid timestamp")
)

type UserRepository interface {
	SaveUser(*domain.User) error
	GetUserById(int64) (*domain.User, error)
}

type MessageRepository interface {
	SaveMessage(*domain.Message) error
	GetMessageById(int64) (*domain.Message, error)
}
