package httpd

import (
	"mime"
	"strings"

	"github.com/labstack/echo"
)

// Context extends echo.Context.
type Context struct {
	echo.Context
}

// ContentType returns the mime type of the request content.
func (c *Context) ContentType() string {
	ctype := c.Request().Header.Get(echo.HeaderContentType)
	mtype, _, _ := mime.ParseMediaType(ctype)

	return mtype
}

// Accept returns the mime type the client accepts as response.
func (c *Context) Accept() []string {
	atypes := strings.Split(c.Request().Header.Get(echo.HeaderAccept), ",")

	for i, atype := range atypes {
		atypes[i], _, _ = mime.ParseMediaType(atype)
	}

	return atypes
}

// Accepts returns true if the client accepts given type.
func (c *Context) Accepts(mtype string) bool {
	for _, atype := range c.Accept() {
		if strings.Contains(atype, mtype) {
			return true
		}
	}

	return false
}

// ExtendContext wraps echo.Context as Context.
func ExtendContext(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := &Context{c}

		return h(ctx)
	}
}
