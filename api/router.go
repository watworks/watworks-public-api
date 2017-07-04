package api

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type route struct {
	method, path string
	handler      http.HandlerFunc
}

var routes = []route{
	route{"GET", "/", IndexHandler},
	route{"GET", "/hello", HelloHandler},
	route{"GET", "/goodye", GoodbyeHandler},
	route{"GET|POST", "/users", NotImplementedHandler},
	route{"GET|PUT|PATCH|DELETE", "/users/{id}", NotImplementedHandler},
}

func NewRouter() *negroni.Negroni {
	// setup the core router w/ api routes
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
