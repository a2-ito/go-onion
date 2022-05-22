package handler

import (
  "fmt"
  "context"
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/a2-ito/go-onion/src/usecase"
  //"github.com/a2-ito/go-onion/src/domain/model"
)

type UserHandler interface {
  Hello(c echo.Context) error
  GetUsers(c echo.Context) error
  CreateUser(c echo.Context) error
}

type userHandler struct {
  UserUseCase usecase.UserUseCase
}

func NewUserHandler(u usecase.UserUseCase) UserHandler {
  return &userHandler{u}
}

func (h *userHandler) Hello(c echo.Context) error {
  ctx := c.Request().Context()
  if ctx == nil {
    ctx = context.Background()
  }

  h.UserUseCase.Hello(ctx)
  return c.JSON(http.StatusOK, "hello")
}

func (h *userHandler) GetUsers(c echo.Context) error {
  ctx := c.Request().Context()
  if ctx == nil {
    ctx = context.Background()
  }

  h.UserUseCase.GetUsers(ctx)
  return c.JSON(http.StatusOK, "get users")
}

type (
  User struct {
    Id  string `json:"name"`
    Name string `json:"email"`
  }
)

func (h *userHandler) CreateUser(c echo.Context) error {
  //user := &model.User{}
  u := new(User)
  if err := c.Bind(u); err != nil {
    return err
  }
  fmt.Println("handler CreateUser ", u.Id, u.Name)
  /*
  ctx := c.Request().Context()
  if ctx == nil {
    ctx = context.Background()
  }

  h.UserUseCase.CreateUser(ctx, name)
  */
  return c.JSON(http.StatusOK, "create user")
}

