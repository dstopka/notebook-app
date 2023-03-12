package adapters_test

import (
	"context"
	"testing"

	"github.com/dstopka/notebook-app/backend/users/internal/adapters"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserRepositoryIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	t.Parallel()

	repositories := []Repository{
		{
			Name: "memory",
			Impl: adapters.NewMemoryUserRepository(),
		},
	}

	for i := range repositories {
		repo := repositories[i]

		t.Run(repo.Name, func(t *testing.T) {
			t.Parallel()

			t.Run("testGetUser", func(t *testing.T) {
				t.Parallel()
				testGetUser(t, repo.Impl)
			})

			t.Run("testUpdateUser", func(t *testing.T) {
				t.Parallel()
				testUpdateUser(t, repo.Impl)
			})

			t.Run("testUpdateUser_existing", func(t *testing.T) {
				t.Parallel()
				testUpdateUser_existing(t, repo.Impl)
			})
		})
	}
}

type Repository struct {
	Name string
	Impl adapters.UserRepository
}

func testGetUser(t *testing.T, repo adapters.UserRepository) {
	t.Helper()
	ctx := context.Background()

	testCases := map[string]struct {
		expectedUser *adapters.User
		insertUser   bool
	}{
		"existing_user": {
			expectedUser: newSimpleUser(),
			insertUser: true,
		},
		"not_existing_user": {
			expectedUser: &adapters.User{},
		},
	}

	for name, tc := range testCases {
		name, tc := name, tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			userID := tc.expectedUser.UUID
			if userID == "" {
				userID = uuid.NewString()
			}

			if tc.insertUser {
				err := repo.UpdateUser(ctx, userID, func(_ *adapters.User) (*adapters.User, error) {
					return tc.expectedUser, nil
				})
				require.NoError(t, err)
			}

			user, err := repo.GetUser(ctx, userID)
			require.NoError(t, err)

			assert.Equal(t, tc.expectedUser, user)
		})
	}
}

func testUpdateUser(t *testing.T, repo adapters.UserRepository) {
	t.Helper()
	ctx := context.Background()

	testCases := map[string]struct {
		getUser func() *adapters.User
	}{
		"with_name_and_role": {
			getUser: func() *adapters.User {
				return newSimpleUser()
			},
		},
		"with_avatarURL": {
			getUser: func() *adapters.User {
				user := newSimpleUser()
				user.AvatarURL = "images.com/test-image"

				return user
			},
		},
		"with_lastIP": {
			getUser: func() *adapters.User {
				user := newSimpleUser()
				user.LastIP = "127.0.0.1"
				
				return user
			},
		},
	}

	for name, tc := range testCases {
		name, tc := name, tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			user := tc.getUser()

			err := repo.UpdateUser(ctx, user.UUID, func(_ *adapters.User) (*adapters.User, error) {
				return user, nil
			})
			require.NoError(t, err)

			assertUserInRepository(ctx, t, repo, user)
		})
	}
}

func testUpdateUser_existing(t *testing.T, repo adapters.UserRepository) {
	t.Helper()
	ctx := context.Background()

	user := newSimpleUser()

	err := repo.UpdateUser(ctx, user.UUID, func(_ *adapters.User) (*adapters.User, error) {
		return user, nil
	})
	require.NoError(t, err)
	assertUserInRepository(ctx, t, repo, user)

	var expectedUser *adapters.User
	err = repo.UpdateUser(ctx, user.UUID, func(u *adapters.User) (*adapters.User, error) {
		u.AvatarURL = "images.com/test-image"
		u.LastIP = "127.0.0.1"

		expectedUser = u
		return u, nil
	})
	require.NoError(t, err)

	assertUserInRepository(ctx, t, repo, expectedUser)
}

func newSimpleUser() *adapters.User {
	return &adapters.User{
		UUID:   uuid.NewString(),
		Name:   "John Doe",
		Role:   "user",
	}
}

func assertUserInRepository(ctx context.Context, t *testing.T, repo adapters.UserRepository, user *adapters.User) {
	require.NotNil(t, user)

	userFromRepo, err := repo.GetUser(ctx, user.UUID)
	require.NoError(t, err)

	assert.Equal(t, user, userFromRepo)
}
