package adapters

import "context"

// UserRepository is the contract for user related database operations.
type UserRepository interface {
	GetUser(ctx context.Context, id string) (*User, error)
	UpdateUser(ctx context.Context, id string, updateFn UpdateUserFn) error
}

// User is the representation of an app user.
type User struct {
	UUID string

	Name      string
	Role      string
	AvatarURL string
	LastIP    string
}