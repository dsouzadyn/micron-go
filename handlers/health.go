package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// HealthHandler handles the health checkups
func HealthHandler(c echo.Context) error {
	return c.String(http.StatusOK, "All good!")
}
