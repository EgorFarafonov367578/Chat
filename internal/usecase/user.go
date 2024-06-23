package usecase

import (
	"chat/internal/domain"
)

type User struct {
	repository UserRepository
}

func NewUser(repository UserRepository) *User {
	return &User{repository: repository}
}

func (u *User) GetUserById(id int64) (*domain.User, error) {
	return u.repository.GetUserById(id)
}

func (u *User) RegisterUser(user *domain.User) error {
	if stringIsMeaningless(user.Name) {
		return ErrInvalidUserName
	}
	return u.repository.SaveUser(user)
}
