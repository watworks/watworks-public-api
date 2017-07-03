package api

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func NewRouter() *negroni.Negroni {
	// setup the core router w/ api routes
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler).Methods("GET")
	r.HandleFunc("/hello", HelloHandler).Methods("GET")
	r.HandleFunc("/goodbye", GoodbyeHandler).Methods("GET")
	r.HandleFunc("/users", NotImplementedHandler).Methods("GET", "POST")
	r.HandleFunc("/users/{id}", NotImplementedHandler).Methods("GET", "POST", "PUT", "PATCH", "DELETE")

	// setup the main handler w/ various middlewares
	// just using the defaults for now... can make this more interesting later
	n := negroni.Classic()
	n.UseHandler(r)

	return n
}
