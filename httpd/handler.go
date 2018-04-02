package httpd

import (
	"net/http"

	"github.com/labstack/echo"
)

// ListSignalGeneratorsHandler serves list of available signal generators
// formated as JSON.
func ListSignalGeneratorsHandler(ctx echo.Context) error {
	c := ctx.(*Context)

	if c.Accepts(echo.MIMEApplicationJSON) {
		return c.NoContent(http.StatusNoContent)
	}
	if c.Accepts(echo.MIMETextHTML) {
		return c.File("public/index.html")
	}

	return echo.ErrUnsupportedMediaType
}

// ShowSignalGeneratorHandler serves configuration of specified signal
// generator formated as JSON.
func ShowSignalGeneratorHandler(ctx echo.Context) error {
	c := ctx.(*Context)

	if c.Accepts(echo.MIMEApplicationJSON) {
		return c.NoContent(http.StatusNoContent)
	}

	return echo.ErrUnsupportedMediaType
}

// UpdateSignalGeneratorHandler updates configuration of specified signal
// generator from JSON formated payload.
func UpdateSignalGeneratorHandler(ctx echo.Context) error {
	c := ctx.(*Context)

	return c.NoContent(http.StatusNoContent)
}
