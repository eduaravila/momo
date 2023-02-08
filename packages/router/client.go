package router

import (
	"io"
	"net/http"
)

type RequestParams struct {
	Url     string
	Body    io.Reader
	Headers [][]string
}

func MakeRequest(method string, params RequestParams) (*http.Response, error) {
	request, err := http.NewRequest(method, params.Url, params.Body)
	if err != nil {
		return nil, err
	}

	for _, header := range params.Headers {
		request.Header.Set(header[0], header[1])
	}
	request.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	return client.Do(request)
}

// Post make a post request to a generic url with a body
func Post(params RequestParams) (*http.Response, error) {
	return MakeRequest(http.MethodPost, params)
}

// Get make a post request to a generic url with a body
func Get(params RequestParams) (*http.Response, error) {
	return MakeRequest(http.MethodGet, params)
}
