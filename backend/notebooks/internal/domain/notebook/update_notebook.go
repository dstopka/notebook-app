package notebook

import (
	"time"

	"github.com/dstopka/notebook-app/backend/notebooks/internal/domain/icon"
)

// ChangeIcon sets notebook's icon to the given icon.
func (n *Notebook) ChangeIcon(icon *icon.Icon) {
	n.icon = icon
	n.edited()
}

// ChangeTitle sets notebook's title to the given title.
func (n *Notebook) ChangeTitle(title string) {
	n.title = title
	n.edited()
}

// ChangeDescription sets notebook's description to the given description.
func (n *Notebook) ChangeDescription(description string) error {
	if len(description) > MaxDescriptionLen {
		return ErrDescriptionTooLong
	}

	n.description = description
	n.edited()
	return nil
}

func (n *Notebook) edited() {
	n.lastEditedTime = time.Now()
}
