package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bodokaiser/beagle/httpd"
	"github.com/bodokaiser/beagle/model"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SpecTestSuite struct {
	suite.Suite

	e *echo.Echo
	h *Spec
}

func (s *SpecTestSuite) SetupTest() {
	s.e = echo.New()
	s.h = &Spec{
		Specs: model.DefaultDDSSpecs,
	}
}

func (s *SpecTestSuite) TestListJSON() {
	req := httptest.NewRequest(echo.GET, "/specs", nil)
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := s.e.NewContext(req, rec)

	h := httpd.WrapContext(s.h.List)

	if assert.NoError(s.T(), h(ctx)) {
		assert.Equal(s.T(), http.StatusOK, rec.Code)
	}
}

func TestSpecHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(SpecTestSuite))
}
