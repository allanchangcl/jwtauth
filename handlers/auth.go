package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Auth struct{}

func (a *Auth) Register(e echo.Context) error {
	// return e.String(http.StatusOK, "Register")
	return e.String(http.StatusOK, "Register")
}
