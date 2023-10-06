package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Jonh Doe", "joe@ig.com", "1234")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Jonh Doe", user.Name)
	assert.Equal(t, "joe@ig.com", user.Email)
}

func TestValidatePassword(t *testing.T) {
	user, err := NewUser("Jonh Doe", "joe@ig.com", "1234")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("1234"))
	assert.False(t, user.ValidatePassword("2341"))
	assert.NotEqual(t, user.Password, "1234")
}
