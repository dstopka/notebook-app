package icon

import "errors"

// Icon represents a single icon object.
type Icon struct {
	emoji string
}

// NewIcon returns an icon created using the given emoji.
func NewIcon(emoji string) (*Icon, error) {
	if emoji == "" {
		return nil, errors.New("empty emoji")
	}

	return &Icon{emoji: emoji}, nil
}

// Emoji returns emoji of the icon.
func (i Icon) Emoji() string {
	return i.emoji
}