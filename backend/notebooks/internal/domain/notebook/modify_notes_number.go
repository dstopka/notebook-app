package notebook

import "errors"

// ErrNoNotesInNotebook is returned when NoteDeleted is called
// and there are no notes in the notebook.
var ErrNoNotesInNotebook = errors.New("notebook does not have any notes")

// HasNotes returns true if there are any notes in the notebook.
func (n Notebook) HasNotes() bool {
	return n.notesNumber > 0
}

// NoteAdded signals that note was added to the notebook
// and increases notesNumber by one.
func (n *Notebook) NoteAdded() {
	n.notesNumber++
}

// NoteDeleted signals that note was deleted from the notebook
// and reduces notesNumber by one.
func (n *Notebook) NoteDeleted() error {
	if !n.HasNotes() {
		return ErrNoNotesInNotebook
	}

	n.notesNumber--
	return nil
}
