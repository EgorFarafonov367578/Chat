package inmemory

import (
	"chat/internal/domain"
	"chat/internal/usecase"
	"sync"
)

type UserRepository struct {
	list  []*domain.User
	mutex sync.Mutex
}

func NewUserRepository() *UserRepository {
	return &UserRepository{list: make([]*domain.User, 0)}
}

func (ur *UserRepository) SaveUser(user *domain.User) error {
	ur.mutex.Lock()
	defer ur.mutex.Unlock()
	user.Id = int64(len(ur.list))
	ur.list = append(ur.list, user)
	return nil
}

func (ur *UserRepository) GetUserById(id int64) (*domain.User, error) {
	ur.mutex.Lock()
	defer ur.mutex.Unlock()
	if id < int64(len(ur.list)) {
		return ur.list[id], nil
	}
	return nil, usecase.ErrUserNotFound
}
