package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HelloWorld(e echo.Context) error {
	return e.String(http.StatusOK, "Hello, World!")
}
