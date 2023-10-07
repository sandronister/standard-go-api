package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sandronister/standart-go-api/internal/dto"
	"github.com/sandronister/standart-go-api/internal/entity"
	"github.com/sandronister/standart-go-api/internal/infra/database"
)

type productHandler struct {
	ProductDB database.ProducInterface
}

func NewProductHandler(db database.ProducInterface) *productHandler {
	return &productHandler{
		ProductDB: db,
	}
}

func (h *productHandler) Create(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductDTO
	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := entity.NewProduct(product.Name, product.Price)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
