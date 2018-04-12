package httpd

import (
	"mime"
	"strings"

	"github.com/labstack/echo"
)

// Context extends echo.Context.
//
// A context can be seen as the bigger scope in which a HTTP response and
// request life. We extend the default context provided by the echo framework
// to give us an entry point for utilities.
type Context struct {
	echo.Context
}

// ContentType returns the normalized mime type of the request content.
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

// Accepts returns true if the client accepts the given mime type.
//
// By using this method we can distinguish between the data format used by
// different clients. For example a web application might prefer JSON formated
// data while a java application might prefer XML.
//
// In case of the client accepting all mime types "*/*" this function will
// just return true.
func (c *Context) Accepts(mtype string) bool {
	for _, atype := range c.Accept() {
		if strings.Contains(atype, mtype) {
			return true
		}
		if atype == "*/*" {
			return true
		}
	}

	return false
}

// WrapContext wraps echo.Context as Context.
//
// You can register this function as middleware in echo such that our
// extended Context type can be interferred from our handlers.
func WrapContext(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := &Context{c}

		return h(ctx)
	}
}
