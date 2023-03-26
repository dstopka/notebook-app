package notebook

import (
	"errors"
	"time"

	"github.com/dstopka/notebook-app/backend/notebooks/internal/domain/icon"
)

// Notebook represents a single notebook object.
type Notebook struct {
	uuid     string
	userUUID string

	title       string
	description string

	icon *icon.Icon

	notesNumber int

	createdTime    time.Time
	lastEditedTime time.Time
}

// ErrDescriptionTooLong is returned when the description is too long.
var ErrDescriptionTooLong = errors.New("description too long")

// MaxDescriptionLen defines max length of a notebook's description.
const MaxDescriptionLen = 1000

// NewNotebook returns a Notebook created using the given values.
func NewNotebook(uuid, userUUID, title, description string, icon *icon.Icon) (*Notebook, error) {
	if uuid == "" {
		return nil, errors.New("empty notebook uuid")
	}
	if userUUID == "" {
		return nil, errors.New("empty userUUID")
	}
	if len(description) > MaxDescriptionLen {
		return nil, ErrDescriptionTooLong
	}

	now := time.Now()

	return &Notebook{
		uuid:           uuid,
		userUUID:       userUUID,
		title:          title,
		description:    description,
		icon:           icon,
		createdTime:    now,
		lastEditedTime: now,
	}, nil
}

// UUID returns uuid of the notebook.
func (n Notebook) UUID() string {
	return n.uuid
}

// UserUUID returns userUUID of the notebook.
func (n Notebook) UserUUID() string {
	return n.userUUID
}

// Title returns title of the notebook.
func (n Notebook) Title() string {
	return n.title
}

// Description returns description of the notebook.
func (n Notebook) Description() string {
	return n.description
}

// Icon returns icon of the notebook.
func (n Notebook) Icon() *icon.Icon {
	return n.icon
}

// NotesNumber returns notesNumber of the notebook.
func (n Notebook) NotesNumber() int {
	return n.notesNumber
}

// CreatedTime returns createdTime of the notebook.
func (n Notebook) CreatedTime() time.Time {
	return n.createdTime
}

// LastEditedTime returns lastEditedTime of the notebook.
func (n Notebook) LastEditedTime() time.Time {
	return n.lastEditedTime
}
