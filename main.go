package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"time"

	"github.com/watworks/watworks-public-api/api"
	"github.com/watworks/watworks-public-api/services"
)

func main() {
	// get config and create the server
	c := services.Config()
	s := &http.Server{Addr: ":80", Handler: api.NewRouter()}
	done := make(chan bool)
	interrupt := make(chan os.Signal, 2)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// listen for interrupt to shutdown the server
	go func() {
		defer close(done)
		<-interrupt
		fmt.Print("Shutting down...")

		// try to gracefully shutdown within 10 seconds
		d := time.Now().Add(time.Second * 10)
		ctx, cancel := context.WithDeadline(context.Background(), d)
		defer cancel()
		if err := s.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
		fmt.Print(" done!\n")
	}()

	// run server... will be closed from goroutine on interrupt
	fmt.Printf("Starting with config: \n%+v\n", c)
	s.ListenAndServe()
	<-done
	fmt.Println("Goodybe.")
	os.Exit(0)
}
