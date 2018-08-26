package engine

import "net/http"

type Route struct {
	Logger  bool
	Tester  bool
	Handler http.Handler
}
