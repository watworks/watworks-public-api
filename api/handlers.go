package api

import (
	"net/http"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello world!"))
}

func NotImplementedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	w.Write([]byte("Not implemented"))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	NotImplementedHandler(w, r)
}

func GoodbyeHandler(w http.ResponseWriter, r *http.Request) {
	NotImplementedHandler(w, r)
}
