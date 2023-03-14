package app

import "context"

// UpdateUserFn defines function type used to update the user.
type UpdateUserFn func(*User) (*User, error)

// UserRepository is the contract for user related database operations.
type UserRepository interface {
	// GetUser finds user by ID.
	GetUser(ctx context.Context, id string) (*User, error)
	// UpdateUser performs the update on the user under the given ID.
	UpdateUser(ctx context.Context, id string, updateFn UpdateUserFn) error
}