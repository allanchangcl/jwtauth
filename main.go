package main

import (
	"database/sql"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/labstack/echo/v4"
)

type handler struct{}

// Handler
func (h *handler) helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
	// return c.String(http.StatusBadRequest, "Hello, World!")
}

func main() {
	db, err := sql.Open("sqlite3", "./jwtauth.db") // not clear on what is needed for the user and password
	if err != nil {
		panic("Could not connect to the database")
		// fmt.Println(err)
		// fmt.Println("error one tripped")
		// return
	}
	defer db.Close()

	e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, world!")
	// })

	h := &handler{}
	e.GET("/", h.helloWorld)
	e.Logger.Fatal(e.Start(":1323"))
}
