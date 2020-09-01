package cursor_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"soccer-api/pkg/cursor"
)

func TestEncode(t *testing.T) {
	newTime := time.Date(2020, time.August, 10, 9, 20, 0, 0, time.UTC)

	want := "MjAyMC0wOC0xMFQwOToyMDowMFo="
	got := cursor.Encode(newTime)
	require.Equal(t, want, got)
}

func TestDecode(t *testing.T) {
	csr := "MjAyMC0wOC0xMFQwOToyMDowMFo="
	want := time.Date(2020, time.August, 10, 9, 20, 0, 0, time.UTC)

	got, err := cursor.Decode(csr)
	require.Equal(t, want, got)
	require.NoError(t, err)
}
