package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Register(e echo.Context) error {
	// return e.String(http.StatusOK, "Register")
	return e.String(http.StatusOK, "Register")
}
