package notebook_test

import (
	"testing"

	"github.com/dstopka/notebook-app/backend/notebooks/internal/domain/notebook"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHasNotes(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		crateNotebook func(t *testing.T) *notebook.Notebook
		expected      bool
	}{
		"with notes": {
			crateNotebook: func(t *testing.T) *notebook.Notebook {
				ntbk := newExampleNotebook(t)
				ntbk.NoteAdded()
				return ntbk
			},
			expected: true,
		},
		"empty notebook": {
			crateNotebook: newExampleNotebook,
			expected:      false,
		},
	}

	for name, tc := range testCases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ntbk := tc.crateNotebook(t)
			hasNotes := ntbk.HasNotes()
			assert.Equal(t, tc.expected, hasNotes)
		})
	}
}

func TestNoteAdded(t *testing.T) {
	t.Parallel()

	ntbk := newExampleNotebook(t)
	preAddNumber := ntbk.NotesNumber()

	ntbk.NoteAdded()
	assert.Equal(t, preAddNumber+1, ntbk.NotesNumber())
}

func TestNoteDeleted(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		ntbk := newExampleNotebook(t)
		ntbk.NoteAdded()

		preDeleteNumber := ntbk.NotesNumber()
		err := ntbk.NoteDeleted()
		require.NoError(t, err)

		assert.Equal(t, preDeleteNumber-1, ntbk.NotesNumber())
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		ntbk := newExampleNotebook(t)

		err := ntbk.NoteDeleted()
		assert.ErrorIs(t, err, notebook.ErrNoNotesInNotebook)
	})
}
