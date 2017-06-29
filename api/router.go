package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func NewRouter() *negroni.Negroni {
	// setup the core router w/ api routes
	r := mux.NewRouter()
	r.HandleFunc("/", HelloWorldHandler).Methods("GET")
	r.HandleFunc("/hello", HelloWorldHandler).Methods("PUT")
	r.HandleFunc("/goodbye", HelloWorldHandler).Methods("PUT")
	r.HandleFunc("/users", NotImplementedHandler).Methods("GET", "POST")
	r.HandleFunc("/users/{id}", NotImplementedHandler).Methods("GET", "POST", "PUT", "PATCH", "DELETE")

	// setup the main handler w/ various middlewares
	// just using the defaults for now... can make this more interesting later
	n := negroni.Classic()
	n.UseHandler(r)

	return n
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello world!"))
}

func NotImplementedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	w.Write([]byte("Not implemented"))
}
