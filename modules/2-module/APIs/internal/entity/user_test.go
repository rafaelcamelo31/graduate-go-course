package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Gopher", "gopher@google.com", "GoPass")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Gopher", user.Name)
	assert.Equal(t, "gopher@google.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Gopher", "gopher@google.com", "GoPass")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("GoPass"))
	assert.False(t, user.ValidatePassword("PoGass"))
	assert.NotEqual(t, "GoPass", user.Password)
}
