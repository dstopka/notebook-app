package icon_test

import (
	"testing"

	"github.com/dstopka/notebook-app/backend/notebooks/internal/domain/icon"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewIcon(t *testing.T) {
	t.Parallel()

	emoji := "🙂"
	icn, err := icon.NewIcon(emoji)
	require.NoError(t, err)

	assert.Equal(t, emoji, icn.Emoji())
}

func TestNewIcon_invalid(t *testing.T) {
	t.Parallel()

	_, err := icon.NewIcon("")
	assert.Error(t, err)
}