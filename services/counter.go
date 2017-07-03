package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var counterService *CounterService

type CounterService struct {
	url    string
	client *http.Client
	err    error
}

type apiReturn struct {
	Value     float64 `json:"value,omitempty"`
	PrevValue float64 `json:"prevValue,omitempty"`
}

func (s *CounterService) call(method, path string) (*http.Response, error) {
	uri, err := url.ParseRequestURI(strings.Join([]string{s.url, path}, ""))
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, uri.String(), nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(req)
}

func (s *CounterService) Get() float64 {
	res, err := s.call("GET", "/counters/main")
	if err != nil {
		s.err = err
		return 0
	}
	defer res.Body.Close()
	in, err := ioutil.ReadAll(res.Body)
	if err != nil {
		s.err = err
		return 0
	}
	var v apiReturn
	s.err = json.Unmarshal(in, &v)
	return v.Value
}

func (s *CounterService) Incr() bool {
	if _, s.err = s.call("PUT", "/counters/main/increment/1"); s.err != nil {
		return false
	}
	return true
}

func (s *CounterService) Decr() bool {
	if _, s.err = s.call("PUT", "/counters/main/decrement/1"); s.err != nil {
		return false
	}
	return true
}

func (s *CounterService) Init() bool {
	if _, s.err = s.call("PUT", "/counters/main"); s.err != nil {
		return false
	}
	return true
}

func (s *CounterService) Reset() bool {
	if _, s.err = s.call("PUT", "/counters/main/0"); s.err != nil {
		return false
	}
	return true
}

func (s *CounterService) Err() error {
	return s.err
}

func Counter() *CounterService {
	if counterService != nil {
		return counterService
	}

	counterService = &CounterService{
		url:    Config().CounterServiceUrl,
		client: &http.Client{},
	}

	return counterService
}
