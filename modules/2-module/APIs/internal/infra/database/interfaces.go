package database

import "github.com/rafaelcamelo31/graduate-go-course/2-module/APIs/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
