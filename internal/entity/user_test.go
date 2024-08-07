package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Joao Paulo", "jojoo@gmail.com", "HomemAranaha213")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Joao Paulo", user.Name)
	assert.Equal(t, "jojoo@gmail.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Joao Paulo", "jojoo@gmail.com", "HomemAranaha213")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("HomemAranaha213"))
	assert.False(t, user.ValidatePassword("76543"))
	assert.NotEqual(t, "HomemAranaha213", user.Password)

}
