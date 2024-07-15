package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/fdfdfdfdfregister", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	auth := Auth{}
	// auth.Register(c)

	// Assertions
	if assert.NoError(t, auth.Register(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}
