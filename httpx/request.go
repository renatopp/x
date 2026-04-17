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

// Fetch performs an HTTP request with the specified method, URL, and options.
// This always returns a Response object, even if there was an error during
// the request. You can check for errors using the IsError() method on the
// Response.
//
// The body of the response is read and stored in the Response struct
// automatically, so you don't need to close it.
//
// By default, Fetch uses a timeout of 60 seconds. You can override this by
// providing a RequestOption with a different Timeout value. Setting Timeout to
// negative means no timeout.
func Fetch(method, url string, opts ...RequestOption) *Response {
	return FetchWithBody(method, url, nil, opts...)
}

// FetchWithBody performs an HTTP request with the specified method, URL, body, and options.
// This always returns a Response object, even if there was an error during
// the request. You can check for errors using the IsError() method on the
// Response.
//
// The body of the response is read and stored in the Response struct
// automatically, so you don't need to close it.
//
// By default, FetchWithBody uses a timeout of 60 seconds. You can override this by
// providing a RequestOption with a different Timeout value. Setting Timeout to
// negative means no timeout.
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

// Get is a shortcut for performing a GET using Fetch.
func Get(url string, opts ...RequestOption) *Response {
	return Fetch(http.MethodGet, url, opts...)
}

// Post is a shortcut for performing a POST using FetchWithBody.
func Post(url string, body []byte, opts ...RequestOption) *Response {
	return FetchWithBody(http.MethodPost, url, body, opts...)
}

// Put is a shortcut for performing a PUT using FetchWithBody.
func Put(url string, body []byte, opts ...RequestOption) *Response {
	return FetchWithBody(http.MethodPut, url, body, opts...)
}

// Delete is a shortcut for performing a DELETE using Fetch.
func Delete(url string, opts ...RequestOption) *Response {
	return Fetch(http.MethodDelete, url, opts...)
}

// Patch is a shortcut for performing a PATCH using FetchWithBody.
func Patch(url string, body []byte, opts ...RequestOption) *Response {
	return FetchWithBody(http.MethodPatch, url, body, opts...)
}

// Head is a shortcut for performing a HEAD using Fetch.
func Head(url string, opts ...RequestOption) *Response {
	return Fetch(http.MethodHead, url, opts...)
}

// Options is a shortcut for performing an OPTIONS using Fetch.
func Options(url string, opts ...RequestOption) *Response {
	return Fetch(http.MethodOptions, url, opts...)
}

// Trace is a shortcut for performing a TRACE using Fetch.
func Trace(url string, opts ...RequestOption) *Response {
	return Fetch(http.MethodTrace, url, opts...)
}

// Connect is a shortcut for performing a CONNECT using Fetch.
func Connect(url string, opts ...RequestOption) *Response {
	return Fetch(http.MethodConnect, url, opts...)
}
