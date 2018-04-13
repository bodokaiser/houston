package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	validator "gopkg.in/go-playground/validator.v9"
)

type HandlerTestSuite struct {
	suite.Suite

	e *echo.Echo
}

func (s *HandlerTestSuite) SetupTest() {
	s.e = echo.New()
	s.e.HTTPErrorHandler = HTTPError
	s.e.GET("/echo-err", func(ctx echo.Context) error {
		return echo.ErrUnauthorized
	})
	s.e.GET("/validation-err", func(ctx echo.Context) error {
		return validator.New().Var("foo", "email")
	})
}

func (s *HandlerTestSuite) TestUpdate() {
	req := httptest.NewRequest(echo.PUT, "/devices/dds/DDS5", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)

	assert.Equal(s.T(), http.StatusNotFound, rec.Code)
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}
