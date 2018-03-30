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
	s.echo.Renderer = NewTemplate("../views/*.html")
}

func (s *HandlerTestSuite) TestIndexHTML() {
	req := httptest.NewRequest(echo.GET, "/", nil)
	req.Header.Set(echo.HeaderAccept, echo.MIMETextHTML)
	rec := httptest.NewRecorder()
	ctx := s.echo.NewContext(req, rec)

	h := ExtendContext(IndexHandler)

	if assert.NoError(s.T(), h(ctx)) {
		assert.Equal(s.T(), http.StatusOK, rec.Code)
	}
}

func (s *HandlerTestSuite) TestIndexJSON() {
	req := httptest.NewRequest(echo.GET, "/", nil)
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := s.echo.NewContext(req, rec)

	h := ExtendContext(IndexHandler)

	assert.Equal(s.T(), h(ctx), echo.ErrUnsupportedMediaType)
}

func (s *HandlerTestSuite) TestListSignalGeneratorsHTML() {
	req := httptest.NewRequest(echo.GET, "/signal-generators", nil)
	req.Header.Set(echo.HeaderAccept, echo.MIMETextHTML)
	rec := httptest.NewRecorder()
	ctx := s.echo.NewContext(req, rec)

	h := ExtendContext(ListSignalGeneratorsHandler)

	assert.Equal(s.T(), h(ctx), echo.ErrUnsupportedMediaType)
}

func (s *HandlerTestSuite) TestListSignalGeneratorsJSON() {
	req := httptest.NewRequest(echo.GET, "/signal-generators", nil)
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := s.echo.NewContext(req, rec)

	h := ExtendContext(ListSignalGeneratorsHandler)

	if assert.NoError(s.T(), h(ctx)) {
		assert.Equal(s.T(), http.StatusNoContent, rec.Code)
	}
}

func (s *HandlerTestSuite) TestShowSignalGeneratorHTML() {
	req := httptest.NewRequest(echo.GET, "/signal-generators/0", nil)
	req.Header.Set(echo.HeaderAccept, echo.MIMETextHTML)
	rec := httptest.NewRecorder()
	ctx := s.echo.NewContext(req, rec)

	h := ExtendContext(ShowSignalGeneratorHandler)

	assert.Equal(s.T(), h(ctx), echo.ErrUnsupportedMediaType)
}

func (s *HandlerTestSuite) TestShowSignalGeneratorJSON() {
	req := httptest.NewRequest(echo.GET, "/signal-generators/0", nil)
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := s.echo.NewContext(req, rec)

	h := ExtendContext(ShowSignalGeneratorHandler)

	if assert.NoError(s.T(), h(ctx)) {
		assert.Equal(s.T(), http.StatusNoContent, rec.Code)
	}
}

func (s *HandlerTestSuite) TestUpdateSignalGenerator() {
	req := httptest.NewRequest(echo.PUT, "/signal-generators/0", nil)
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := s.echo.NewContext(req, rec)

	h := ExtendContext(UpdateSignalGeneratorHandler)

	if assert.NoError(s.T(), h(ctx)) {
		assert.Equal(s.T(), http.StatusNoContent, rec.Code)
	}
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}
