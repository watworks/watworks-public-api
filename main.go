package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/watworks/watworks-public-api/api"
	"github.com/watworks/watworks-public-api/services"
)

func main() {
	// get config and create the server
	c := services.Config()
	s := &http.Server{Addr: ":80", Handler: api.NewRouter()}
	done := make(chan os.Signal, 2)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	// listen for interrupt to shutdown the server
	go func() {
		<-done
		fmt.Print("Shutting down...")
		if err := s.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// run server... will be closed from goroutine on interrupt
	fmt.Printf("Starting with config: \n%+v\n", c)
	s.ListenAndServe()
	fmt.Print(" done!\n")
	os.Exit(0)
}
