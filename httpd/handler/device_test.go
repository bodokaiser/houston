package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/bodokaiser/beagle/httpd"
	"github.com/bodokaiser/beagle/model"
)

type DeviceTestSuite struct {
	suite.Suite

	e *echo.Echo
}

func (s *DeviceTestSuite) SetupTest() {
  h := &Device{
		Devices: model.DefaultDDSDevices,
	}

	s.e = echo.New()
  s.e.Use(httpd.WrapContext)
  s.e.GET("/devices")
	s.e.PUT("/devices/:name")

func (s *DeviceTestSuite) TestListJSON() {
	req := httptest.NewRequest(echo.GET, "/devices", nil)
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

  s.e.ServeHTTP(rec, req)

	assert.Equal(s.T(), http.StatusOK, rec.Code)
}

func (s *DeviceTestSuite) TestUpdateJSON() {
	req := httptest.NewRequest(echo.PUT, "/devices/DDS 0", nil)
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := s.e.NewContext(req, rec)

	h := httpd.WrapContext(s.h.Update)

	if assert.NoError(s.T(), h(ctx)) {
		assert.Equal(s.T(), http.StatusNoContent, rec.Code)
	}
}

func TestDeviceTestSuite(t *testing.T) {
	suite.Run(t, new(DeviceTestSuite))
}
