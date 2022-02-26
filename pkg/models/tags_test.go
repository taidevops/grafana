package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseTagPairs(t *testing.T) {
	t.Run("Can parse one empty tag", func(t *testing.T) {
		tags := ParseTagPairs([]string{""})
		require.Empty(t, tags)
	})
}
