package main

import (
	// "net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()
	e.Static("/", "lib/www")
	e.Static("/src", "lib/src")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!\n")
	// })

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}