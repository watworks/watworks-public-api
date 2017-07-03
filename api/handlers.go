package api

import (
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: maybe return redirect to UI service index, or call it internally?
	// ... or is it up to the UI service to render the app?
	// It should be up to the UI service to render the app.
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
