package database

import "github.com/sandronister/standart-go-api/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type ProducInterface interface {
	Create(product *entity.Product) error
	FindByName(name string) (*entity.Product, error)
	FindById(id string) *entity.Product
}
