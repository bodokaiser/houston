// Package handler provides HTTP handlers to mediate between the model
// and the driver package.
package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

// HTTPError implements custom HTTP error handling for echo.
func HTTPError(err error, ctx echo.Context) {
	if valErr, ok := err.(validator.ValidationErrors); ok {
		err = echo.NewHTTPError(http.StatusBadRequest,
			fmt.Sprintf("%s is invalid", valErr[0].Field()))
	}

	ctx.Echo().DefaultHTTPErrorHandler(err, ctx)
}
