package utils_test

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestParseFileSize(t *testing.T) {
    size, err := ParseFileSize("50MB")
    assert.NoError(t, err)
    assert.Equal(t, int64(50*1024*1024), size)

    size, err = ParseFileSize("10KB")
    assert.NoError(t, err)
    assert.Equal(t, int64(10*1024), size)

    _, err = ParseFileSize("invalid")
    assert.Error(t, err)
}

func TestHashString(t *testing.T) {
    hash := HashString("test string")
    assert.Equal(t, "3474851a3410906697ec77337df7aae4a0e4bb57a4027bfba63667f029db3d07", hash)
}

func TestIsAllowedMimeType(t *testing.T) {
    allowedTypes := []string{"application/pdf", "text/plain"}
    assert.True(t, IsAllowedMimeType("application/pdf", allowedTypes))
    assert.False(t, IsAllowedMimeType("image/png", allowedTypes))
}
