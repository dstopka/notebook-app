package app

import "context"

// Application is controller of the users service.
type Application struct {
	repo UserRepository
}

// New creates a new Application.
func New(repo UserRepository) *Application {
	if repo == nil {
		panic("nil user repository")
	}

	return &Application{
		repo: repo,
	}
}

// HandleGetUser handles getting user by id.
func (a *Application) HandleGetUser(ctx context.Context, id string) (*User, error) {
	return a.repo.GetUser(ctx, id)
}

// HandleSetAvatar handles setting the avatar url of the user.
func (a *Application) HandleSetAvatar(ctx context.Context, userID, avatarURL string) error {
	return a.repo.UpdateUser(ctx, userID, func(u *User) (*User, error) {
		u.AvatarURL = avatarURL
		return u, nil
	})
}

// HandleSetLastIP handles setting last ip of the user.
func (a *Application) HandleSetLastIP(ctx context.Context, userID, ip string) error {
	return a.repo.UpdateUser(ctx, userID, func(u *User) (*User, error) {
		u.LastIP = ip
		return u, nil
	})
}
