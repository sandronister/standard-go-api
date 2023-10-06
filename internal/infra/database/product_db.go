package database

import (
	"github.com/sandronister/standart-go-api/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProductDB(db *gorm.DB) ProducInterface {
	return &Product{DB: db}
}

func (p *Product) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) FindByName(name string) (*entity.Product, error) {
	var product entity.Product
	if err := p.DB.Where("name=?", name).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *Product) FindById(id string) *entity.Product {
	var product entity.Product
	p.DB.First(&product, "id=?", id)
	return &product
}
