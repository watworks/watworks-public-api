package main

import (
	"log"
	"net/http"

	"github.com/watworks/watworks-public-api/api"
	"github.com/watworks/watworks-public-api/services"
)

func main() {

	// TODO: load a config file!  This will be fun.  If it doesn't load, must exit
	if false {
		log.Fatal("no config file")
	}

	c := services.Config()
	http.ListenAndServe(c.ServerUrl, api.NewRouter())

	log.Println("Running on :80...")
}
