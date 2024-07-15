package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/allanchangcl/jwtauth/handlers"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	// req, err := http.NewRequest("GET", "/", nil)
	// if err != nil {
	//   t.Fatal(err)
	// }

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// h := handler{}
	// h.helloWorld(c)
	handlers.Index(c)

	assert.Equal(t, http.StatusOK, rec.Code)
}
