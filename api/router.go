package api

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// define all the api routes
var routes = []struct {
	method, path string
	handler      http.HandlerFunc
}{
	{"GET", "/", IndexHandler},
	{"GET", "/hello", HelloHandler},
	{"GET", "/goodye", GoodbyeHandler},
	{"GET|POST", "/users", NotImplementedHandler},
	{"GET|PUT|PATCH|DELETE", "/users/{id}", NotImplementedHandler},
}

// NewRouter creates and returns a router with all api routes
// and middlewares registered
func NewRouter() *negroni.Negroni {
	// setup the core router w/ defined routes
	router := mux.NewRouter()
	for _, r := range routes {
		router.HandleFunc(r.path, r.handler).Methods(strings.Split(r.method, "|")...)
	}

	// setup the main handler w/ various middlewares
	// just using the defaults for now... can make this more interesting later

	// TODO: allow cors probably
	n := negroni.Classic()
	n.UseHandler(router)

	return n
}
