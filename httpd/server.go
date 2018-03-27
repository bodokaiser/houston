package httpd

import (
	"net/http"

	"github.com/labstack/echo"
)

func Handler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello")
}
