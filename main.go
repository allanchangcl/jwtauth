package main

import (
	"net/http"

	"github.com/allanchangcl/jwtauth/database"
	_ "github.com/mattn/go-sqlite3"

	"github.com/labstack/echo/v4"
)

// type handler struct{}

// Handler
// func (h *handler) helloWorld(c echo.Context) error {
func helloWorld(c echo.Context) error {
	// return c.String(http.StatusBadRequest, "Hello, World!")
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	database.Connection()

	e := echo.New()
	// h := &handler{}
	e.GET("/", helloWorld)
	e.Logger.Fatal(e.Start(":1323"))
}
