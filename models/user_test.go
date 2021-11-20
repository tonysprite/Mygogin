package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginAction(t *testing.T) {
	result, err := LoginQuery("tony", "ABC")
	assert.Empty(result, err)
}

func TestLogoutAction(t *testing.T) {
	result, err := LogoutAction(1223)
	assert.True(result, assert.Error(err))
}
