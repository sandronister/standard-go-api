package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/sandronister/standard-go-api/internal/dto"
	"github.com/sandronister/standard-go-api/internal/entity"
	"github.com/sandronister/standard-go-api/internal/infra/database"
	entityPKG "github.com/sandronister/standard-go-api/pkg/entity"
)

type productHandler struct {
	ProductDB database.ProducInterface
}

func NewProductHandler(db database.ProducInterface) *productHandler {
	return &productHandler{
		ProductDB: db,
	}
}

// Createproduct godoc
// @Summary      Create Product
// @Description  Create Product
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        request     body      dto.CreateProductDTO  true  "product request"
// @Success      201
// @Failure      500         {object}  Error
// @Router       /products [post]
// @Security     ApiKeyAuth
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

// GetProduct godoc
// @Summary      Get Product
// @Description  Get Product
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id          path        string    true    "Product ID" Format(uuid)
// @Success      200         {object}    entity.Product
// @Failure      400         {object}    Error
// @Failure      500         {object}    Error
// @Router       /products/{id} [get]
// @Security     ApiKeyAuth
func (h *productHandler) FindOne(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		err := Error{Message: "Invalid ID"}
		json.NewEncoder(w).Encode(err)
		return
	}

	product, err := h.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&product)
}

// UpdateProduct godoc
// @Summary      Update Product
// @Description  Update Product
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id          path      string    true    "Product ID" Format(uuid)
// @Param        request     body      dto.CreateProductDTO  true  "product request"
// @Success      200
// @Failure      400         {object}   Error
// @Failure      500         {object}   Error
// @Router       /products/{id} [put]
// @Security     ApiKeyAuth
func (h *productHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product := entity.Product{}
	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product.ID, err = entityPKG.ParseId(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&product)
}

// DeleteProduct godoc
// @Summary      Update Product
// @Description  Update Product
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id          path      string    true    "Product ID" Format(uuid)
// @Success      200
// @Failure      400         {object}   Error
// @Failure      500         {object}   Error
// @Router       /products/{id} [delete]
// @Security     ApiKeyAuth
func (h *productHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := h.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Listproduct godoc
// @Summary      List Products
// @Description  Get All Products
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        page        query      string    false    "page number"
// @Param        limit       query      string    false    "limit"
// @Success      200         {array}    entity.Product
// @Failure      500         {object}   Error
// @Router       /products [get]
// @Security     ApiKeyAuth
func (h *productHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}

	sort := r.URL.Query().Get("sort")

	products, err := h.ProductDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
