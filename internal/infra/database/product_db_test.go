package database

import (
	"testing"

	"github.com/sandronister/standart-go-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory.db"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, _ := entity.NewProduct("Prod1", 125.23)
	productDB := NewProductDB(db)
	err = productDB.Create(product)
	assert.Nil(t, err)
	assert.NotNil(t, product)
}

func TestFindByName(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory.db"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	productDB := NewProductDB(db)
	prodFound, err := productDB.FindByName("Prod1")
	assert.Nil(t, err)
	assert.NotNil(t, prodFound)

	prodNotFound, err := productDB.FindByName("Prod2")
	assert.NotNil(t, err)
	assert.Nil(t, prodNotFound)
}

func TestFindByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory.db"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	productDB := NewProductDB(db)
	prodName, _ := productDB.FindByName("Prod1")

	prodFound := productDB.FindById(prodName.ID.String())
	assert.NotNil(t, prodFound)
}
