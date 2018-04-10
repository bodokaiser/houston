package httpd

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SpecHandlerTestSuite struct {
	suite.Suite

	echo *echo.Echo
}

func (s *SpecHandlerTestSuite) SetupTest() {
	s.echo = echo.New()
}

func (s *SpecHandlerTestSuite) TestListSpecsJSON() {
	req := httptest.NewRequest(echo.GET, "/specs", nil)
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := s.echo.NewContext(req, rec)

	h := WrapContext(ListSpecsHandler)

	if assert.NoError(s.T(), h(ctx)) {
		assert.Equal(s.T(), http.StatusOK, rec.Code)
	}
}

func TestSpecHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(SpecHandlerTestSuite))
}
