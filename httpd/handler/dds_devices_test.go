package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/bodokaiser/houston/driver"
	"github.com/bodokaiser/houston/httpd"
	"github.com/bodokaiser/houston/model"
)

type DDSTestSuite struct {
	suite.Suite

	e *echo.Echo
}

func (s *DDSTestSuite) SetupTest() {
	h := &DDSDevices{
		Devices: model.DefaultDDSDevices,
		Driver:  &driver.MockedDDSArray{},
	}

	s.e = echo.New()
	s.e.Use(httpd.WrapContext)
	s.e.GET("/devices/dds", h.List)
	s.e.PUT("/devices/dds/:name", h.Update)
}

func (s *DDSTestSuite) TestListJSON() {
	req := httptest.NewRequest(echo.GET, "/devices/dds", nil)
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)

	assert.Equal(s.T(), http.StatusOK, rec.Code)
}

func (s *DDSTestSuite) TestUpdateJSON() {
	req := httptest.NewRequest(echo.PUT, "/devices/dds/DDS0", nil)
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)
}

func TestDeviceTestSuite(t *testing.T) {
	suite.Run(t, new(DDSTestSuite))
}
