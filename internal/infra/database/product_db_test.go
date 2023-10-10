package database

import (
	"testing"

	"github.com/google/uuid"
	"github.com/sandronister/standard-go-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func getConnect(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file:memory.db"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	return db
}

func TestCreateProduct(t *testing.T) {
	db := getConnect(t)
	db.AutoMigrate(&entity.Product{})
	product, _ := entity.NewProduct("Prod1", 125.23)
	productDB := NewProductDB(db)
	err := productDB.Create(product)
	assert.Nil(t, err)
	assert.NotNil(t, product)
}

func TestFindByID(t *testing.T) {
	db := getConnect(t)
	entProd, _ := entity.NewProduct("Prod2", 123.23)
	productDB := NewProductDB(db)
	productDB.Create(entProd)

	foundProd, err := productDB.FindById(entProd.ID.String())
	assert.Nil(t, err)
	assert.NotNil(t, foundProd)

	_, err = productDB.FindById(uuid.New().String())
	assert.NotNil(t, err)
}

func TestFindByAll(t *testing.T) {
	db := getConnect(t)
	productDB := NewProductDB(db)
	products, err := productDB.FindAll(0, 0, "asc")
	assert.Nil(t, err)
	assert.NotNil(t, products)

	products, err = productDB.FindAll(1, 2, "asc")
	assert.Nil(t, err)
	assert.NotNil(t, products)

	products, err = productDB.FindAll(1, 2, "")
	assert.Nil(t, err)
	assert.NotNil(t, products)
}

func TestUpdate(t *testing.T) {
	db := getConnect(t)
	productDB := NewProductDB(db)
	entProd, _ := entity.NewProduct("Prod6", 123.23)
	productDB.Create(entProd)
	entProd.Price = 520.30

	err := productDB.Update(entProd)
	assert.Nil(t, err)
	assert.Equal(t, 520.30, entProd.Price)

	entProd.ID = uuid.New()
	err = productDB.Update(entProd)
	assert.NotNil(t, err)
}

func TestDelete(t *testing.T) {
	db := getConnect(t)
	productDB := NewProductDB(db)
	entProd, _ := entity.NewProduct("Prod6", 123.23)
	productDB.Create(entProd)

	err := productDB.Delete(entProd.ID.String())
	assert.Nil(t, err)

	err = productDB.Delete(uuid.New().String())
	assert.NotNil(t, err)
}
