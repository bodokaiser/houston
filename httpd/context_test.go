package httpd

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ContextTestSuite struct {
	suite.Suite

	echo *echo.Echo
}

func (s *ContextTestSuite) SetupTest() {
	s.echo = echo.New()
	s.echo.Use(WrapContext)
}

func (s *HandlerTestSuite) TestContentTypeHTML() {
	req := httptest.NewRequest(echo.GET, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMETextHTML)
	rec := httptest.NewRecorder()
	ctx := Context{s.echo.NewContext(req, rec)}

	assert.Equal(s.T(), echo.MIMETextHTML, ctx.ContentType())
}

func (s *HandlerTestSuite) TestContentTypeJSON() {
	req := httptest.NewRequest(echo.GET, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := Context{s.echo.NewContext(req, rec)}

	assert.Equal(s.T(), echo.MIMEApplicationJSON, ctx.ContentType())
}

func (s *HandlerTestSuite) TestContentTypeInvalid() {
	req := httptest.NewRequest(echo.GET, "/", nil)
	req.Header.Set(echo.HeaderContentType, "?")
	rec := httptest.NewRecorder()
	ctx := Context{s.echo.NewContext(req, rec)}

	assert.Equal(s.T(), "", ctx.ContentType())
}

func (s *HandlerTestSuite) TestAcceptChrome() {
	req := httptest.NewRequest(echo.GET, "/", nil)
	req.Header.Set(echo.HeaderAccept, "text/html,application/xhtml+xml")
	rec := httptest.NewRecorder()
	ctx := Context{s.echo.NewContext(req, rec)}

	assert.Equal(s.T(), []string{"text/html", "application/xhtml+xml"},
		ctx.Accept())
}

func (s *HandlerTestSuite) TestAcceptsHTML() {
	req := httptest.NewRequest(echo.GET, "/", nil)
	req.Header.Set(echo.HeaderAccept, "text/html,application/xhtml+xml")
	rec := httptest.NewRecorder()
	ctx := Context{s.echo.NewContext(req, rec)}

	assert.True(s.T(), ctx.Accepts(echo.MIMETextHTML))
	assert.False(s.T(), ctx.Accepts(echo.MIMEApplicationJavaScript))
}

func (s *HandlerTestSuite) TestWrapContext() {
	req := httptest.NewRequest(echo.GET, "/", nil)
	req.Header.Set(echo.HeaderAccept, "text/html,application/xhtml+xml")
	rec := httptest.NewRecorder()
	ctx := s.echo.NewContext(req, rec)

	h := WrapContext(func(c echo.Context) error {
		return c.(*Context).NoContent(http.StatusOK)
	})

	if assert.NoError(s.T(), h(ctx)) {
		assert.Equal(s.T(), http.StatusOK, rec.Code)
	}
}

func TestContextSuite(t *testing.T) {
	suite.Run(t, new(ContextTestSuite))
}
