package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type stubRepo struct {
	users map[string]User
}

func (r stubRepo) ByEmail(email string) (User, error) {
	u, ok := r.users[email]
	if !ok {
		return User{}, ErrNotFound
	}
	return u, nil
}

func TestService_FindIDByEmail(t *testing.T) {
	t.Run("user found", func(t *testing.T) {
		repo := stubRepo{
			users: map[string]User{
				"user@test.com": {ID: 42, Email: "user@test.com"},
			},
		}
		svc := New(repo)

		id, err := svc.FindIDByEmail("user@test.com")

		require.NoError(t, err)
		assert.Equal(t, int64(42), id)
	})

	t.Run("user not found", func(t *testing.T) {
		repo := stubRepo{
			users: map[string]User{},
		}
		svc := New(repo)

		_, err := svc.FindIDByEmail("missing@test.com")

		assert.ErrorIs(t, err, ErrNotFound)
	})
}
