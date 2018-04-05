package httpd

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type HandlerTestSuite struct {
	suite.Suite

	echo *echo.Echo
}

func (s *HandlerTestSuite) SetupTest() {
	s.echo = echo.New()
}

func (s *HandlerTestSuite) TestListDevicesHTML() {
	req := httptest.NewRequest(echo.GET, "/devices", nil)
	req.Header.Set(echo.HeaderAccept, echo.MIMETextHTML)
	rec := httptest.NewRecorder()
	ctx := s.echo.NewContext(req, rec)

	h := WrapContext(ListDevicesHandler)

	assert.Equal(s.T(), h(ctx), echo.ErrUnsupportedMediaType)
}

func (s *HandlerTestSuite) TestListDevicesJSON() {
	req := httptest.NewRequest(echo.GET, "/devices", nil)
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := s.echo.NewContext(req, rec)

	h := WrapContext(ListDevicesHandler)

	if assert.NoError(s.T(), h(ctx)) {
		assert.Equal(s.T(), http.StatusNoContent, rec.Code)
	}
}

func (s *HandlerTestSuite) TestShowDeviceHTML() {
	req := httptest.NewRequest(echo.GET, "/devices/0", nil)
	req.Header.Set(echo.HeaderAccept, echo.MIMETextHTML)
	rec := httptest.NewRecorder()
	ctx := s.echo.NewContext(req, rec)

	h := WrapContext(ShowDeviceHandler)

	assert.Equal(s.T(), h(ctx), echo.ErrUnsupportedMediaType)
}

func (s *HandlerTestSuite) TestShowDeviceJSON() {
	req := httptest.NewRequest(echo.GET, "/devices/0", nil)
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := s.echo.NewContext(req, rec)

	h := WrapContext(ShowDeviceHandler)

	if assert.NoError(s.T(), h(ctx)) {
		assert.Equal(s.T(), http.StatusNoContent, rec.Code)
	}
}

func (s *HandlerTestSuite) TestUpdateDevice() {
	req := httptest.NewRequest(echo.PUT, "/devices/0", nil)
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := s.echo.NewContext(req, rec)

	h := WrapContext(UpdateDeviceHandler)

	if assert.NoError(s.T(), h(ctx)) {
		assert.Equal(s.T(), http.StatusNoContent, rec.Code)
	}
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}
