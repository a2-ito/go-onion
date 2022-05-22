package usecase

import (
  "fmt"
  "context"
  "github.com/labstack/echo/v4"
  "github.com/a2-ito/go-onion/src/domain/repository"
  "github.com/google/uuid"
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
  Hello(ctx context.Context) error
  GetUser(ctx context.Context, id uuid.UUID) error
  GetUsers(ctx context.Context) error
  CreateUser(ctx context.Context, name string) error
}

type userUseCase struct {
  repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) UserUseCase {
  return &userUseCase{r}
}

func HelloTest(c echo.Context) (err error) {
  fmt.Println("interactor HelloTest")
  //interactor.UserRepository.Hello()
  //interactor.userRepository.FindById()
  return c.JSON(200, "Hello, World!")
}

func (usecase *userUseCase) Hello(ctx context.Context) (err error) {
  fmt.Println("UseCase Hello")
  usecase.UserRepository.Hello()
  return nil
}

func (usecase *userUseCase) GetUser(ctx context.Context, id uuid.UUID) (err error) {
  fmt.Println("UseCase FindById")
  usecase.UserRepository.FindById(id)
  return nil
}

func (usecase *userUseCase) GetUsers(ctx context.Context) (err error) {
  fmt.Println("UseCase FindById")
  usecase.UserRepository.FindAll()
  return nil
}

func (usecase *userUseCase) CreateUser(ctx context.Context, name string) (err error) {
  fmt.Println("UseCase CreateUser")
  usecase.UserRepository.Create(name)
  return nil
}
