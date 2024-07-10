package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestController(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/register", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// h := handler{}
	// h.helloWorld(c)
	// HelloWorld(c)
	Register(c)

	assert.Equal(t, http.StatusOK, rec.Code)
}
