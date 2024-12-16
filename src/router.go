package main

import (
	"back/handlers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func newRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	renderer := Renderer()
	e.Renderer = renderer

	e.GET("/home", index)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	api := e.Group("/api")
	api.POST("/park", handlers.AddParking)
	api.GET("/park", handlers.GetParkings)
	api.GET("/park/:id", handlers.GetParking)
	api.PUT("/park/:id", handlers.UpdateParking)
	api.DELETE("/park/:id", handlers.DeleteParking)

	api.POST("/user", handlers.AddUser)
	api.GET("/user/:name", handlers.GetUser)
	api.GET("/user", handlers.GetUsers)
	api.DELETE("/user/:name", handlers.DeleteUser)

	return e
}

func index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}
