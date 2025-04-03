package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/rafaelcamelo31/graduate-go-course/2-module/APIs/internal/dto"
	"github.com/rafaelcamelo31/graduate-go-course/2-module/APIs/internal/entity"
	"github.com/rafaelcamelo31/graduate-go-course/2-module/APIs/internal/infra/database"
	"github.com/rafaelcamelo31/graduate-go-course/2-module/APIs/pkg/value_object"
)

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: db,
	}
}

// GetJWT godoc
// @Summary      Get a user JWT
// @Description  Get a user JWT
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request   body     dto.GetJWTInput  true  "user credentials"
// @Success      200  {object}  dto.GetJWTOutput
// @Failure      404  {object}  value_object.Error
// @Failure      500  {object}  value_object.Error
// @Router       /users/get_token [post]
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
	accessToken := dto.GetJWTOutput{AccessToken: tokenString}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)

}

// Create user godoc
// @Summary      Create user
// @Description  Create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request     body      dto.CreateUserInput  true  "user request"
// @Success      201
// @Failure      500         {object}  value_object.Error
// @Router       /users [post]
func (uh *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var userDto dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := value_object.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	user, err := entity.NewUser(userDto.Name, userDto.Email, userDto.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := value_object.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	err = uh.UserDB.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := value_object.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusOK)
}
