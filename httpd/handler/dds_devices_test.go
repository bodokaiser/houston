package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/bodokaiser/houston/driver/dds"
	"github.com/bodokaiser/houston/driver/mux"
	"github.com/bodokaiser/houston/httpd"
	"github.com/bodokaiser/houston/model"
)

type DDSTestSuite struct {
	suite.Suite

	e *echo.Echo
	h *DDSDevices
	d model.DDSDevice
}

func (s *DDSTestSuite) SetupTest() {
	s.d = model.DDSDevice{
		ID:   3,
		Name: "Champ",
		Amplitude: model.DDSParam{
			Const: model.DDSConst{Value: 1.0},
		},
		Frequency: model.DDSParam{
			Const: model.DDSConst{Value: 250e6},
		},
	}
	s.h = &DDSDevices{
		Devices: model.DDSDevices{s.d},
		Mux:     &mux.Mockup{},
		DDS:     &dds.Mockup{},
	}

	s.e = echo.New()
	s.e.Use(httpd.WrapContext)
	s.e.GET("/devices/dds", s.h.List)
	s.e.PUT("/devices/dds/:id", s.h.Update)
	s.e.DELETE("/devices/dds/:id", s.h.Delete)
	s.e.POST("/devices/dds/:id/trigger", s.h.Trigger)
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

	d := model.DDSDevices{}
	assert.Equal(s.T(), http.StatusOK, rec.Code)
	assert.NoError(s.T(), json.Unmarshal([]byte(rec.Body.String()), &d))
	assert.Equal(s.T(), s.h.Devices, d)
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
	req := httptest.NewRequest(echo.PUT, "/devices/dds/2", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)

	assert.Equal(s.T(), http.StatusNotFound, rec.Code)
}

func (s *DDSTestSuite) TestUpdateSweepJSON() {
	d := model.DDSDevice{
		ID:   3,
		Name: "Abi Haft",
		Amplitude: model.DDSParam{
			Mode: model.ModeSweep,
			Sweep: model.DDSSweep{
				Limits:   [2]float64{0, 1.0},
				Duration: 10,
			},
		},
		Frequency: model.DDSParam{
			Const: model.DDSConst{Value: 250e6},
		},
	}
	json, err := json.Marshal(d)
	assert.NoError(s.T(), err)

	req := httptest.NewRequest(echo.PUT, "/devices/dds/3", bytes.NewBuffer(json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)

	assert.Equal(s.T(), http.StatusNoContent, rec.Code)
	assert.Equal(s.T(), d, s.h.Devices[0])
	assert.Equal(s.T(), dds.SweepConfig{
		Limits:   [2]float64{0, 1.0},
		Duration: 10 * time.Second,
		Param:    dds.ParamAmplitude,
	}, s.h.DDS.Sweep())
	assert.Equal(s.T(), 250e6, s.h.DDS.Frequency())
	assert.Equal(s.T(), float64(0), s.h.DDS.PhaseOffset())
	assert.True(s.T(), s.h.DDS.(*dds.Mockup).HadExec)
}

func (s *DDSTestSuite) TestUpdateConstJSON() {
	s.h.Devices[0].Amplitude = model.DDSParam{
		Mode: model.ModeSweep,
		Sweep: model.DDSSweep{
			Limits:   [2]float64{0, 1.0},
			Duration: 10,
		},
	}

	d := model.DDSDevice{
		ID:   3,
		Name: "Abi Haft",
		Amplitude: model.DDSParam{
			Const: model.DDSConst{Value: 0.5},
		},
		Frequency: model.DDSParam{
			Const: model.DDSConst{Value: 200e6},
		},
	}
	json, err := json.Marshal(d)
	assert.NoError(s.T(), err)

	req := httptest.NewRequest(echo.PUT, "/devices/dds/3", bytes.NewBuffer(json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)

	assert.Equal(s.T(), http.StatusNoContent, rec.Code)
	assert.Equal(s.T(), d, s.h.Devices[0])
	assert.Equal(s.T(), 0.5, s.h.DDS.Amplitude())
	assert.Equal(s.T(), 200e6, s.h.DDS.Frequency())
	assert.Equal(s.T(), float64(0), s.h.DDS.PhaseOffset())
	assert.True(s.T(), s.h.DDS.(*dds.Mockup).HadExec)
}

func (s *DDSTestSuite) TestDelete() {
	req := httptest.NewRequest(echo.DELETE, "/devices/dds/3", nil)
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)

	assert.Equal(s.T(), http.StatusNoContent, rec.Code)
	assert.EqualValues(s.T(), 3, s.h.Mux.(*mux.Mockup).Selected)
	assert.True(s.T(), s.h.DDS.(*dds.Mockup).HadReset)
}

func (s *DDSTestSuite) TestTrigger() {
	req := httptest.NewRequest(echo.POST, "/devices/dds/3/trigger", nil)
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)

	assert.Equal(s.T(), http.StatusNoContent, rec.Code)
	assert.True(s.T(), s.h.DDS.(*dds.Mockup).HadUpdate)
}

func TestDeviceTestSuite(t *testing.T) {
	suite.Run(t, new(DDSTestSuite))
}
