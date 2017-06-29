package services

import (
	"net/http"
)

var counterService *CounterService

type CounterService struct {
	url    string
	client *http.Client
}

func (s *CounterService) Get() float64 {
	return 0
}
func (s *CounterService) Incr() bool {
	return true
}
func (s *CounterService) Decr() bool {
	return true
}

func Counter() *CounterService {
	if counterService == nil {
		initCounterService()
	}

	return counterService
}

func initCounterService() {
	counterService = &CounterService{
		url:    "http://localhost:8004",
		client: &http.Client{},
	}
}
