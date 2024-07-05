package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct{}

// Handler
func (h *handler) helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, world!")
	// })

	h := &handler{}
	e.GET("/", h.helloWorld)
	e.Logger.Fatal(e.Start(":1323"))
}
