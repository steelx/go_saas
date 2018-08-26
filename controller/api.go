package controllers

import (
	"context"
	"fmt"
	"go_saas/engine"
	"net/http"
)

// API is the starting point of our API.
// Responsible for routing the request to the correct handler
type API struct {
	User *engine.Route
}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, engine.ContextOriginalPath, r.URL.Path)
	var next *engine.Route
	var head string
	head, r.URL.Path = engine.ShiftPath(r.URL.Path)
	if head == "user" {
		next = newUser()
	} else {
		next = newError(fmt.Errorf("path not found"), http.StatusNotFound)
	}
	if next.Logger {
		next.Handler = engine.Logger(next.Handler)
	}
	next.Handler.ServeHTTP(w, r.WithContext(ctx))
}
func newError(err error, statusCode int) *engine.Route {
	return &engine.Route{
		Logger: true,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			engine.Respond(w, r, statusCode, err)
		}),
	}
}
