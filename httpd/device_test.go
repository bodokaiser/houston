package httpd

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DeviceHandlerTestSuite struct {
	suite.Suite

	echo *echo.Echo
}

func (s *DeviceHandlerTestSuite) SetupTest() {
	s.echo = echo.New()
}

func (s *DeviceHandlerTestSuite) TestListDevicesJSON() {
	req := httptest.NewRequest(echo.GET, "/devices", nil)
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := s.echo.NewContext(req, rec)

	h := WrapContext(ListDevicesHandler)

	if assert.NoError(s.T(), h(ctx)) {
		assert.Equal(s.T(), http.StatusOK, rec.Code)
	}
}

func (s *DeviceHandlerTestSuite) TestUpdateDevice() {
	req := httptest.NewRequest(echo.PUT, "/devices/0", nil)
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := s.echo.NewContext(req, rec)

	h := WrapContext(UpdateDeviceHandler)

	if assert.NoError(s.T(), h(ctx)) {
		assert.Equal(s.T(), http.StatusNoContent, rec.Code)
	}
}

func TestDeviceHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(DeviceHandlerTestSuite))
}
