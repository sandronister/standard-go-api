package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/sandronister/standard-go-api/internal/dto"
	"github.com/sandronister/standard-go-api/internal/entity"
	"github.com/sandronister/standard-go-api/internal/infra/database"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB       database.UserInterface
	Jwt          *jwtauth.JWTAuth
	JWTExpiresIn int
}

func NewUserHanlder(db database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: db,
	}
}

// Create token godoc
// @Summary      Get a user JWT
// @Description  Get a user JWT
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request       body      dto.GetJwtInput  true  "user credentials"
// @Success      200           {object}  dto.GetJWTOutput
// @Failure      404           {object}  Error
// @Failure      500           {object}  Error
// @Router       /users/login  [post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	expiresIn := r.Context().Value("expiresin").(int)
	var login dto.GetJwtInput
	err := json.NewDecoder(r.Body).Decode(&login)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.UserDB.FindByEmail(login.Email)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	if !u.ValidatePassword(login.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(expiresIn)).Unix()})

	accessToken := dto.GetJWTOutput{AccesToken: tokenString}
	w.Header().Set("Content-Type", "json/application")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Create user godoc
// @Summary      Create user
// @Description  Create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request     body      dto.CreateUserDTO  true  "user request"
// @Success      201
// @Failure      500         {object}  Error
// @Router       /users [post]
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
