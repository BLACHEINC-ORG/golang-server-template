package middleware

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/opensaucerer/gotemplate/typing"
)

// ParamsF is a middleware that takes a request handler.
// It sets a variable paramas and pushes it into the request's context
// using the ParamsCtxKey.
func ParamsH() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//set variable params
			params := mux.Vars(r)

			//set params to the request context
			ctx := context.WithValue(r.Context(), typing.ParamsCtxKey{}, params)

			//call next handler
			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}
}

// ParamsF is a middleware that takes a controller function.
// It sets a variable paramas and pushes it into the request's context
// using the ParamsCtxKey.
func ParamsF() func(func(http.ResponseWriter, *http.Request)) http.Handler {
	return func(next func(http.ResponseWriter, *http.Request)) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//set variable params
			params := mux.Vars(r)

			//set params to the request context
			ctx := context.WithValue(r.Context(), typing.ParamsCtxKey{}, params)

			//call next handler
			next(w, r.WithContext(ctx))

		})
	}
}
