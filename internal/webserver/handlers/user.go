package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/sandronister/standart-go-api/internal/dto"
	"github.com/sandronister/standart-go-api/internal/entity"
	"github.com/sandronister/standart-go-api/internal/infra/database"
)

type UserHandler struct {
	UserDB       database.UserInterface
	Jwt          *jwtauth.JWTAuth
	JWTExpiresIn int
}

func NewUserHanlder(db database.UserInterface, jwt *jwtauth.JWTAuth, expiresIn int) *UserHandler {
	return &UserHandler{
		UserDB:       db,
		Jwt:          jwt,
		JWTExpiresIn: expiresIn,
	}
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var login dto.GetJwtInput
	err := json.NewDecoder(r.Body).Decode(&login)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.UserDB.FindByEmail(login.Email)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !u.ValidatePassword(login.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := h.Jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JWTExpiresIn)).Unix()})

	accessToken := struct {
		AccesToken string `json:"acess_token"`
	}{
		AccesToken: tokenString,
	}
	w.Header().Set("Content-Type", "json/application")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserDTO
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
