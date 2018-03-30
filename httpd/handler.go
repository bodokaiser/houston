package httpd

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

// IndexHandler serves index HTML.
func IndexHandler(c echo.Context) error {
	t := c.Request().Header.Get(echo.HeaderContentType)

	if strings.Contains(t, echo.MIMETextHTML) {
		return c.Render(http.StatusOK, "index.html", nil)
	}

	return echo.ErrUnsupportedMediaType
}

// ListSignalGeneratorsHandler serves list of available signal generators
// formated as JSON.
func ListSignalGeneratorsHandler(c echo.Context) error {
	t := c.Request().Header.Get(echo.HeaderContentType)

	if strings.Contains(t, echo.MIMEApplicationJSON) {
		return c.NoContent(http.StatusNoContent)
	}

	return echo.ErrUnsupportedMediaType
}

// ShowSignalGeneratorHandler serves configuration of specified signal
// generator formated as JSON.
func ShowSignalGeneratorHandler(c echo.Context) error {
	t := c.Request().Header.Get(echo.HeaderContentType)

	if strings.Contains(t, echo.MIMEApplicationJSON) {
		return c.NoContent(http.StatusNoContent)
	}

	return echo.ErrUnsupportedMediaType
}

// UpdateSignalGeneratorHandler updates configuration of specified signal
// generator from JSON formated payload.
func UpdateSignalGeneratorHandler(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
