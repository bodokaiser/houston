package httpd

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bodokaiser/beagle/model"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DeviceHandlerTestSuite struct {
	suite.Suite

	echo    *echo.Echo
	handler *DeviceHandler
}

func (s *DeviceHandlerTestSuite) SetupTest() {
	s.echo = echo.New()
	s.handler = &DeviceHandler{
		Devices: []model.Device{},
	}
}

func (s *DeviceHandlerTestSuite) TestListJSON() {
	req := httptest.NewRequest(echo.GET, "/devices", nil)
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := s.echo.NewContext(req, rec)

	h := WrapContext(s.handler.List)

	if assert.NoError(s.T(), h(ctx)) {
		assert.Equal(s.T(), http.StatusOK, rec.Code)
	}
}

func (s *DeviceHandlerTestSuite) TestUpdate() {
	req := httptest.NewRequest(echo.PUT, "/devices/0", nil)
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := s.echo.NewContext(req, rec)

	h := WrapContext(s.handler.Update)

	if assert.NoError(s.T(), h(ctx)) {
		assert.Equal(s.T(), http.StatusNoContent, rec.Code)
	}
}

func TestDeviceHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(DeviceHandlerTestSuite))
}
