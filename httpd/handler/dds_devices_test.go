package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
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
	h *DDSDevices
}

func (s *DDSTestSuite) SetupTest() {
	s.h = &DDSDevices{
		Devices: []model.DDSDevice{
			model.DDSDevice{
				Name:      "DDS0",
				ID:        3,
				Amplitude: 1.0,
				Frequency: 250e6,
			},
		},
		Driver: &driver.MockedDDSArray{},
	}

	s.e = echo.New()
	s.e.Use(httpd.WrapContext)
	s.e.GET("/devices/dds", s.h.List)
	s.e.PUT("/devices/dds/:name", s.h.Update)
}

func (s *DDSTestSuite) TestList() {
	req := httptest.NewRequest(echo.GET, "/devices/dds", nil)
	req.Header.Set(echo.HeaderAccept, echo.MIMETextHTML)
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)

	assert.Equal(s.T(), http.StatusUnsupportedMediaType, rec.Code)
}

func (s *DDSTestSuite) TestListJSON() {
	req := httptest.NewRequest(echo.GET, "/devices/dds", nil)
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)

	assert.Equal(s.T(), http.StatusOK, rec.Code)
	assert.Equal(s.T(),
		`[{"name":"DDS0","amplitude":1,"frequency":250000000,"phase":0}]`,
		rec.Body.String())
}

func (s *DDSTestSuite) TestListXML() {
	req := httptest.NewRequest(echo.GET, "/devices/dds", nil)
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationXML)
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)

	assert.Equal(s.T(), http.StatusOK, rec.Code)
	assert.True(s.T(), strings.HasPrefix(rec.Body.String(), `<?xml version="1.0"`))
}

func (s *DDSTestSuite) TestUpdate() {
	req := httptest.NewRequest(echo.PUT, "/devices/dds/DDS5", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)

	assert.Equal(s.T(), http.StatusNotFound, rec.Code)
}

func (s *DDSTestSuite) TestUpdateJSON() {
	req := httptest.NewRequest(echo.PUT, "/devices/dds/DDS0",
		bytes.NewBuffer([]byte(`{"name":"DDS0","amplitude":0.5,"frequency":100000000}`)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)

	assert.Equal(s.T(), http.StatusNoContent, rec.Code)
	assert.EqualValues(s.T(), 0.5, s.h.Devices[0].Amplitude)
	assert.EqualValues(s.T(), 100e6, s.h.Devices[0].Frequency)
	assert.EqualValues(s.T(), 3, s.h.Driver.(*driver.MockedDDSArray).Address)
	assert.EqualValues(s.T(), 100e6, s.h.Driver.(*driver.MockedDDSArray).Frequency)
}

func TestDeviceTestSuite(t *testing.T) {
	suite.Run(t, new(DDSTestSuite))
}
