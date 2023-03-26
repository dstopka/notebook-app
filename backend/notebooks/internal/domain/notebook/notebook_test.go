package notebook_test

import (
	"strings"
	"testing"

	"github.com/dstopka/notebook-app/backend/notebooks/internal/domain/icon"
	"github.com/dstopka/notebook-app/backend/notebooks/internal/domain/notebook"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTraining(t *testing.T) {
	t.Parallel()
	ntbkUUID := uuid.NewString()
	userUUID := uuid.NewString()
	title := "test notebook title"
	description := "test notebook title"
	var icon *icon.Icon

	ntbk, err := notebook.NewNotebook(ntbkUUID, userUUID, title, description, icon)
	require.NoError(t, err)

	assert.Equal(t, ntbkUUID, ntbk.UUID())
	assert.Equal(t, userUUID, ntbk.UserUUID())
	assert.Equal(t, title, ntbk.Title())
	assert.Equal(t, description, ntbk.Description())
	assert.Equal(t, icon, ntbk.Icon())
	assert.Equal(t, 0, ntbk.NotesNumber())
	assert.Equal(t, ntbk.CreatedTime(), ntbk.LastEditedTime())
}

func TestNewTraining_invalid(t *testing.T) {
	t.Parallel()

	ntbkUUID := uuid.NewString()
	userUUID := uuid.NewString()
	title := "test notebook title"
	description := "test notebook title"
	var icon *icon.Icon

	t.Run("empty notebook uuid", func(t *testing.T) {
		t.Parallel()

		_, err := notebook.NewNotebook("", userUUID, title, description, icon)
		assert.Error(t, err)
	})

	t.Run("empty userUUID", func(t *testing.T) {
		t.Parallel()

		_, err := notebook.NewNotebook(ntbkUUID, "", title, description, icon)
		assert.Error(t, err)
	})

	t.Run("too long description", func(t *testing.T) {
		t.Parallel()

		description := strings.Repeat("a", notebook.MaxDescriptionLen+1)
		_, err := notebook.NewNotebook(ntbkUUID, userUUID, title, description, icon)
		assert.Error(t, err)
	})
}

func newExampleNotebook(t *testing.T) *notebook.Notebook {
	ntbk, err := notebook.NewNotebook(
		uuid.NewString(),
		uuid.NewString(),
		"notebook title",
		"notebook description",
		nil,
	)
	require.NoError(t, err)

	return ntbk
}
