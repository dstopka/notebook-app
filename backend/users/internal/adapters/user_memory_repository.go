package adapters

import (
	"context"
	"sync"
)

type UpdateUserFn func(*User) (*User, error)

type MemoryUserRepository struct {
	users map[string]User
	lock  *sync.RWMutex
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		users: make(map[string]User),
		lock:  &sync.RWMutex{},
	}
}

func (r *MemoryUserRepository) GetUser(_ context.Context, id string) (*User, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	return r.getOrCreateUser(id)
}

func (r *MemoryUserRepository) getOrCreateUser(id string) (*User, error) {
	user, ok := r.users[id]
	if !ok {
		return &User{}, nil
	}

	return &user, nil
}

func (r *MemoryUserRepository) UpdateUser(_ context.Context, id string, updateFn UpdateUserFn) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	user, err := r.getOrCreateUser(id)
	if err != nil {
		return err
	}

	updatedUser, err := updateFn(user)
	if err != nil {
		return err
	}

	r.users[id] = *updatedUser
	return nil
}
