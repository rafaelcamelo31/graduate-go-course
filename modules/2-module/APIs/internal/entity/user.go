package entity

import (
	"github.com/rafaelcamelo31/graduate-go-course/2-module/APIs/pkg/value_object"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       value_object.ID `json:"id"`
	Name     string          `json:"name"`
	Email    string          `json:"email"`
	Password string          `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:       value_object.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

func NewTestUser() (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte("GoPass"), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:       value_object.NewID(),
		Name:     "Gopher",
		Email:    "gopher@gmail.com",
		Password: string(hash),
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
