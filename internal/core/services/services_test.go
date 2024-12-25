package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServices(t *testing.T) {
	assert.Equal(t, "hello, world", NewPrompt("hello, world"), "Prompt without metadata should return the same message")
	assert.Equal(t, "hello, world with users_data", NewPrompt("hello, world with METADATA", "METADATA", "users_data"), "Prompt with metadata should return the same message")
}
