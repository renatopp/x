package httpx

import (
	"bytes"
	"maps"
	"net/http"
	"time"
)

type RequestOption struct {
	Headers map[string]string
	Timeout time.Duration
}

func Fetch(method, url string, opts ...RequestOption) *Response {
	return FetchWithBody(method, url, nil, opts...)
}

func FetchWithBody(method, url string, body []byte, opts ...RequestOption) *Response {
	options := RequestOption{
		Headers: make(map[string]string),
		Timeout: 60 * time.Second, // Default timeout
	}

	for _, opt := range opts {
		if opt.Headers != nil {
			maps.Copy(options.Headers, opt.Headers)
		}
		if opt.Timeout != 0 {
			options.Timeout = opt.Timeout
		}
		if opt.Timeout < 0 {
			options.Timeout = 0 // No timeout
		}
	}

	client := &http.Client{Timeout: options.Timeout}
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return newReponse(req, nil, err)
	}

	for k, v := range options.Headers {
		req.Header.Set(k, v)
	}

	response, err := client.Do(req)
	return newReponse(req, response, err)
}

func Get(url string, opts ...RequestOption) *Response {
	return Fetch(http.MethodGet, url, opts...)
}

func Post(url string, body []byte, opts ...RequestOption) *Response {
	return FetchWithBody(http.MethodPost, url, body, opts...)
}

func Put(url string, body []byte, opts ...RequestOption) *Response {
	return FetchWithBody(http.MethodPut, url, body, opts...)
}

func Delete(url string, opts ...RequestOption) *Response {
	return Fetch(http.MethodDelete, url, opts...)
}

func Patch(url string, body []byte, opts ...RequestOption) *Response {
	return FetchWithBody(http.MethodPatch, url, body, opts...)
}

func Head(url string, opts ...RequestOption) *Response {
	return Fetch(http.MethodHead, url, opts...)
}

func Options(url string, opts ...RequestOption) *Response {
	return Fetch(http.MethodOptions, url, opts...)
}

func Trace(url string, opts ...RequestOption) *Response {
	return Fetch(http.MethodTrace, url, opts...)
}

func Connect(url string, opts ...RequestOption) *Response {
	return Fetch(http.MethodConnect, url, opts...)
}
