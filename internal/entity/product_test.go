package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Prod 1", 120.0)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, product.Name, "Prod 1")
	assert.Equal(t, product.Price, 120.0)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", 120)
	assert.Equal(t, err, ErrNameIsRequired)
	assert.Nil(t, product)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	product, err := NewProduct("Prod", 0)
	assert.Equal(t, err, ErrPriceIsRequired)
	assert.Nil(t, product)
}

func TestProductWhenInvalidPrice(t *testing.T) {
	product, err := NewProduct("Prod", -12.0)
	assert.Equal(t, err, ErrInvalidPrice)
	assert.Nil(t, product)
}

func TestProductValidate(t *testing.T) {
	product, err := NewProduct("Prod", 12.0)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Nil(t, product.Validate())
}
