package bear

import (
	"context"
	"net/http"
)

type Context struct {
	context.Context
	Authorization string
	RequestId     string
	Origin        string
}

func ReadContext(r *http.Request) Context {
	return Context{
		Context:       r.Context(),
		Authorization: r.Header.Get(HeaderAuthorization),
		RequestId:     r.Header.Get(HeaderRequestId),
		Origin:        r.Header.Get(HeaderOrigin),
	}
}

func (c Context) Write(r *http.Request) {
	r.Header.Set(HeaderAuthorization, c.Authorization)
	r.Header.Set(HeaderRequestId, c.RequestId)
	r.Header.Set(HeaderOrigin, c.Origin)
}
