package adapters

import (
	"context"
	"sync"

	"github.com/dstopka/notebook-app/backend/users/internal/app"
)

// MemoryUserRepository is an in-memory implementation of UserRepository.
type MemoryUserRepository struct {
	users map[string]app.User
	lock  *sync.RWMutex
}

// NewMemoryUserRepository creates a new MemoryUserRepository.
func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		users: make(map[string]app.User),
		lock:  &sync.RWMutex{},
	}
}

func (r *MemoryUserRepository) GetUser(_ context.Context, id string) (*app.User, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	return r.getOrCreateUser(id)
}

func (r *MemoryUserRepository) getOrCreateUser(id string) (*app.User, error) {
	user, ok := r.users[id]
	if !ok {
		return &app.User{}, nil
	}

	return &user, nil
}

func (r *MemoryUserRepository) UpdateUser(_ context.Context, id string, updateFn app.UpdateUserFn) error {
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
