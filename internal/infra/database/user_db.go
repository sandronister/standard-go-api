package database

import (
	"github.com/sandronister/standard-go-api/internal/entity"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUserDB(db *gorm.DB) UserInterface {
	return &User{DB: db}
}

func (u *User) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *User) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := u.DB.Where("email=?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
