package httpd

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetHandler responds with rendered html.
func GetHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}
