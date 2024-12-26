package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServices(t *testing.T) {
	assert.Equal(t, "hello, world", replace("hello, world"), "Prompt without metadata should return the same message")
	assert.Equal(t, "hello, world with users_data", replace("hello, world with METADATA", "METADATA", "users_data"), "Prompt with metadata should return the same message")
	assert.Equal(t, "hello, world with nothing", replace("hello, world with nothing", "METADATA", "users_data"), "Prompt with no metadata should return the same message")
}
