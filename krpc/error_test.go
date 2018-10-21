package krpc

import (
	"testing"

	"github.com/amivin123/tortortor/bencode"
	"github.com/stretchr/testify/require"
)

// https://github.com/amivin123/tortortor/issues/166
func TestUnmarshalBadError(t *testing.T) {
	var e Error
	err := bencode.Unmarshal([]byte(`l5:helloe`), &e)
	require.Error(t, err)
}
