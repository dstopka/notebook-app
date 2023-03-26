package notebook_test

import (
	"strings"
	"testing"

	"github.com/dstopka/notebook-app/backend/notebooks/internal/domain/icon"
	"github.com/dstopka/notebook-app/backend/notebooks/internal/domain/notebook"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChangeIcon(t *testing.T) {
	t.Parallel()
	notebook := newExampleNotebook(t)

	icon, err := icon.NewIcon("😀")
	require.NoError(t, err)

	preUpdateTime := notebook.LastEditedTime()
	notebook.ChangeIcon(icon)

	assert.Equal(t, icon, notebook.Icon())
	assert.True(t, preUpdateTime.Before(notebook.LastEditedTime()))
}

func TestChangeTitle(t *testing.T) {
	t.Parallel()
	notebook := newExampleNotebook(t)

	newTitle := "updated notebook title"
	preUpdateTime := notebook.LastEditedTime()
	notebook.ChangeTitle(newTitle)

	assert.Equal(t, newTitle, notebook.Title())
	assert.True(t, preUpdateTime.Before(notebook.LastEditedTime()))
}

func TestChangeDescription(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		ntbk := newExampleNotebook(t)

		newDesc := "updated notebook description"
		preUpdateTime := ntbk.LastEditedTime()
		err := ntbk.ChangeDescription(newDesc)
		require.NoError(t, err)
	
		assert.Equal(t, newDesc, ntbk.Description())
		assert.True(t, preUpdateTime.Before(ntbk.LastEditedTime()))
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()
		ntbk := newExampleNotebook(t)

		newDesc := strings.Repeat("a", notebook.MaxDescriptionLen+1)
		err := ntbk.ChangeDescription(newDesc)
		assert.Error(t, err)
	})
}
