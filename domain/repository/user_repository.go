package repository

import (
	"github.com/google/uuid"
	"go-onion/domain/model"
)

type UserRepository interface {
	Hello()
	FindById(id uuid.UUID) (*model.User, error)
	FindAll() ([]*model.User, error)
	Create(name string) (*model.User, error)
	/*
	  Update(domain.User) (*domain.User, error)
	  Delete(id uuid.UUID) error
	*/
}
