package usecase

import (
	"context"
	"fmt"
	//"github.com/google/uuid"
	// "github.com/labstack/echo/v4"
	"go-onion/domain/repository"
)

/*
type UserRepository interface {
  Hello() error
  //FindById(id string) (domain.User, error)
}


type UserInteractor struct {
  UserRepository repository.UserRepository
}
*/

type UserUseCase interface {
	Hello(ctx context.Context)
	/*
	        GetUser(ctx context.Context, id uuid.UUID) error
		GetUsers(ctx context.Context) error
		CreateUser(ctx context.Context, name string) error
	*/
}

type userUseCase struct {
	repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) UserUseCase {
	return &userUseCase{r}
}

func (usecase *userUseCase) Hello(ctx context.Context) {
	fmt.Println("UseCase Hello")
	usecase.UserRepository.Hello()
}

/*
func (usecase *userUseCase) GetUser(ctx context.Context, id uuid.UUID) (err error) {
	fmt.Println("UseCase FindById ", id)
	usecase.UserRepository.FindById(id)
	return nil
}

func (usecase *userUseCase) GetUsers(ctx context.Context) (err error) {
	fmt.Println("UseCase FindAll")
	usecase.UserRepository.FindAll()
	return nil
}

func (usecase *userUseCase) CreateUser(ctx context.Context, name string) (err error) {
	fmt.Println("UseCase CreateUser ", name)
	usecase.UserRepository.Create(name)
	return nil
}
*/
