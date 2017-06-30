package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHelloWorldHandler(t *testing.T) {
	res := callApi(t, "GET", "/", nil)
	require.Equal(t, 200, res.StatusCode)
}

func TestHelloHandler(t *testing.T) {

}

func TestGoodbyeHandler(t *testing.T) {

}

func callApi(t *testing.T, method, path string, body io.Reader) *http.Response {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		t.Fatal(err)
	}
	c := http.Client{}
	s := httptest.NewServer(NewRouter())
	defer s.Close()

	// the url must be absolute, so depends on the test server
	var newPath []string
	newPath = append(newPath, s.URL, req.URL.Path)

	// check for query string and ensure that it's encoded properly
	if "" != req.URL.RawQuery {
		parsedQuery, err := url.ParseQuery(req.URL.RawQuery)
		if err != nil {
			t.Fatal(err)
		}
		newPath = append(newPath, "?", parsedQuery.Encode())
	}
	newUrl, err := url.ParseRequestURI(strings.Join(newPath, ""))
	if err != nil {
		t.Fatal(err)
	}

	req.URL = newUrl

	res, err := c.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	return res
}
