package handler

import (
	"context"
	"fmt"
	//"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	//"github.com/a2-ito/go-onion/src/domain/model"
	"go-onion/usecase"
)

type UserHandler interface {
	Hello(c echo.Context) error
        /*
	GetUser(c echo.Context) error
	GetUsers(c echo.Context) error
	CreateUser(c echo.Context) error
        */
}

type userHandler struct {
	UserUseCase usecase.UserUseCase
}

func NewUserHandler(u usecase.UserUseCase) UserHandler {
	return &userHandler{u}
}

func (h *userHandler) Hello(c echo.Context) error {
	fmt.Println("handler Hello")
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	h.UserUseCase.Hello(ctx)
	return c.JSON(http.StatusOK, "hello")
}

/*
func (h *userHandler) GetUser(c echo.Context) error {
	id := uuid.MustParse(c.Param("id"))
	fmt.Println("handler GetUser ", id)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	h.UserUseCase.GetUser(ctx, id)
	return c.JSON(http.StatusOK, "get users")
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
		Id   string `json:"name"`
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

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	h.UserUseCase.CreateUser(ctx, u.Name)
	return c.JSON(http.StatusOK, "create user")
}
*/
