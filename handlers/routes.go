package handlers

import (
	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo) {
	e.GET("/", Index)

	// authentication
	// e.POST("/api", Register)
}
