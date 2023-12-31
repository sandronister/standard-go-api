package database

import (
	"testing"

	"github.com/sandronister/standard-go-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUserDB(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory.db"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Jonh Doe", "jo@doe.com", "1234")
	userDB := NewUserDB(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id=?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.Name, userFound.Name)
	assert.NotEmpty(t, userFound.Password)
}

func TestFindUserByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory.db"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	userDB := NewUserDB(db)
	userFound, err := userDB.FindByEmail("jo@doe.com")
	assert.Nil(t, err)
	assert.NotNil(t, userFound)

	userFound, err = userDB.FindByEmail("joe@doe.com")
	assert.NotNil(t, err)
	assert.Nil(t, userFound)
}
