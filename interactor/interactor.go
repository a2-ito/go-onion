package interactor

import (
	"go-onion/domain/repository"
	"go-onion/infrastructure/database/memory"
	"go-onion/presenter/handler"
	"go-onion/usecase"
)

type Interactor interface {
	NewUserUseCase() usecase.UserUseCase
	NewUserHandler() handler.UserHandler
}

type interactor struct {
}

func NewInteractor() Interactor {
	return &interactor{}
}

type appHandler struct {
	handler.UserHandler
}

func (i *interactor) NewUserRepository() repository.UserRepository {
	return memory.NewUserRepository()
}

func (i *interactor) NewUserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(i.NewUserRepository())
}

func (i *interactor) NewUserHandler() handler.UserHandler {
	appHandler := &appHandler{}
	appHandler.UserHandler = handler.NewUserHandler(i.NewUserUseCase())
	return appHandler
}
