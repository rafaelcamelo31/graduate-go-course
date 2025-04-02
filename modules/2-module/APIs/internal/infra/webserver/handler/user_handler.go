package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/rafaelcamelo31/graduate-go-course/2-module/APIs/internal/dto"
	"github.com/rafaelcamelo31/graduate-go-course/2-module/APIs/internal/entity"
	"github.com/rafaelcamelo31/graduate-go-course/2-module/APIs/internal/infra/database"
)

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: db,
	}
}

func (uh *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	expiresIn := r.Context().Value("expiresIn").(int)
	var jwtDto dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&jwtDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := uh.UserDB.FindByEmail(jwtDto.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !user.ValidatePassword(jwtDto.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	claims := map[string]any{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(expiresIn)).Unix(),
	}

	_, tokenString, _ := jwt.Encode(claims)
	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)

}

func (uh *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var userDto dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := entity.NewUser(userDto.Name, userDto.Email, userDto.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = uh.UserDB.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
