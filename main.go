package main

import (
	"log"
	"net/http"

	"github.com/watworks/watworks-public-api/api"
)

func main() {

	// TODO: load a config file!  This will be fun.  If it doesn't load, must exit
	if false {
		log.Fatal("no config file")
	}

	r := api.NewRouter()
	http.ListenAndServe(":80", r)

	log.Println("Running on :80...")
}
