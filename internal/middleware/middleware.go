package middleware

import (
	"context"
)

type HandlerFunc func(ctx context.Context, req interface{}) (interface{}, error)

type Middleware func(HandlerFunc) HandlerFunc

func Chain(middlewares ...Middleware) Middleware {
	return func(final HandlerFunc) HandlerFunc {
		for i := len(middlewares) - 1; i >= 0; i-- {
			final = middlewares[i](final)
		}
		return final
	}
}
