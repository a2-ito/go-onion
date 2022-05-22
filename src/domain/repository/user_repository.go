package repository

import (
  "github.com/google/uuid"
  "github.com/a2-ito/go-onion/src/domain/model"
)

type UserRepository interface {
  Hello() error
  FindById(id uuid.UUID) (*model.User, error)
  FindAll() ([]*model.User, error)
  Create(name string) (*model.User, error)
  /*
  Update(domain.User) (*domain.User, error)
  Delete(id uuid.UUID) error
  */
}
