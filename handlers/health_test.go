package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestHealthEndpoint(t *testing.T) {
	// Setup echo
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, HealthHandler(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "All good!", rec.Body.String())
	}
}
