package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Middlewares struct {
	constructors []Middleware
}

func Chain(middlewares ...Middleware) Middlewares {
	return Middlewares{append(([]Middleware)(nil), middlewares...)}
}

func (m Middlewares) To(h http.Handler) http.Handler {
	if h == nil {
		h = http.DefaultServeMux
	}
	for i := range m.constructors {
		h = m.constructors[len(m.constructors)-1-i](h)
	}
	return h
}

func (m Middlewares) ToFunc(fn http.HandlerFunc) http.Handler {
	if fn == nil {
		return m.To(nil)
	}
	return m.To(fn)
}