package presenter

import (
  "net/http"

  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "github.com/a2-ito/go-onion/src/interactor"
)

// ハンドラーを定義
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}

func Init() {
    // Echo instance
    e := echo.New()

    //userController := controllers.NewUserController(NewSqlHandler())
    //userController := controllers.NewUserController()

    i := interactor.NewInteractor()
    h := i.NewUserHandler()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    //e.GET("/", hello)
    e.GET("/", h.Hello)
    //e.GET("/", h.GetUser)
    e.GET("/users", h.GetUsers)
    e.POST("/users", h.CreateUser)

    /*
    e.GET("/", func(c echo.Context) error { return userController.Hello2(c) })
    e.GET("/users", func(c echo.Context) error { return userController.FindAll(c) })
    e.POST("/user", func(c echo.Context) error { return userController.Create(c) })

    e.GET("/users/:id", func(c echo.Context) error { return userController.Show(c) })
    e.PUT("/users/:id", func(c echo.Context) error { return userController.Save(c) })
    e.DELETE("/users/:id", func(c echo.Context) error { return userController.Delete(c) })
    */

    // Start server
    e.Logger.Fatal(e.Start(":1323"))
}
