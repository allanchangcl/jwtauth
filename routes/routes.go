package routes

import (
	"github.com/allanchangcl/jwtauth/controllers"
	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo) {
	e.GET("/", controllers.HelloWorld)
}
