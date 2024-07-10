package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Index(e echo.Context) error {
	return e.String(http.StatusOK, "Hello, World!")
	// return e.String(http.StatusForbidden, "Hello, World!")
}
