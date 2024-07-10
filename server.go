package main

import (
	"github.com/allanchangcl/jwtauth/database"
	"github.com/allanchangcl/jwtauth/handlers"

	_ "github.com/mattn/go-sqlite3"

	"github.com/labstack/echo/v4"
)

func main() {
	database.Connect()

	e := echo.New()
	handlers.Setup(e)

	e.Logger.Fatal(e.Start(":1323"))
}
