package httpx

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
)

// Response represents the result of an HTTP request. It contains the response body,
// status code, headers, and any error that occurred during the request.
type Response struct {
	err              error
	body             []byte
	HttpRequest      *http.Request
	HttpResponse     *http.Response
	Status           string // e.g. "200 OK"
	StatusCode       int    // e.g. 200
	Protocol         string // e.g. "HTTP/1.0"
	ProtocolMajor    int    // e.g. 1
	ProtocolMinor    int    // e.g. 0
	Headers          http.Header
	ContentLength    int64
	TransferEncoding []string
	TLS              *tls.ConnectionState
}

// newResponse creates a new Response object from the given http.Request,
// http.Response, and error. Internal use only.
func newReponse(req *http.Request, resp *http.Response, err error) *Response {
	if err != nil {
		return &Response{err: err, body: []byte{}}
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &Response{err: err, body: []byte{}}
	}

	return &Response{
		HttpRequest:      req,
		HttpResponse:     resp,
		Status:           resp.Status,
		StatusCode:       resp.StatusCode,
		Protocol:         resp.Proto,
		ProtocolMajor:    resp.ProtoMajor,
		ProtocolMinor:    resp.ProtoMinor,
		Headers:          resp.Header,
		ContentLength:    resp.ContentLength,
		TransferEncoding: resp.TransferEncoding,
		TLS:              resp.TLS,
		body:             body,
	}
}

// IsError returns true if there was an error during the HTTP request. You can
// get the error using the Error() method. Notice that a non-2xx status code
// does not count as an error.
func (r *Response) IsError() bool { return r.err != nil }

// Error returns the error that occurred during the HTTP request, or nil if
// there was no error. Notice that a non-2xx status code does not count as an
// error.
func (r *Response) Error() error { return r.err }

// Cookies returns the cookies sent by the server in the response. If there was
// an error during the request, it returns an empty slice.
func (r *Response) Cookies() []*http.Cookie {
	if r.IsError() {
		return []*http.Cookie{}
	}
	return r.HttpResponse.Cookies()
}

// Header returns the value of the specified header key. If the header is not
// present, it returns an empty string. If there was an error during the request,
// it also returns an empty string.
func (r *Response) Header(key string) string {
	return r.Headers.Get(key)
}

// Text returns the response body as a string. If there was an error during the request,
// it returns an empty string.
func (r *Response) Text() string { return string(r.Bytes()) }

// Bytes returns the response body as a byte slice. If there was an error during the request,
// it returns an empty byte slice.
func (r *Response) Bytes() []byte { return r.body }

// Json unmarshals the response body into the provided variable.
func (r *Response) Json(v any) error { return json.Unmarshal(r.body, v) }

// IsSuccess returns true if the HTTP status code is in the 2xx range, indicating a successful response.
func (r *Response) IsSuccess() bool { return r.StatusCode >= 200 && r.StatusCode < 300 }

// Is2xx is an alias for IsSuccess, returning true if the HTTP status code is in the 2xx range.
func (r *Response) Is2xx() bool { return r.IsSuccess() }

// Is3xx returns true if the HTTP status code is in the 3xx range, indicating a redirection response.
func (r *Response) Is3xx() bool { return r.StatusCode >= 300 && r.StatusCode < 400 }

// Is4xx returns true if the HTTP status code is in the 4xx range, indicating a client error response.
func (r *Response) Is4xx() bool { return r.StatusCode >= 400 && r.StatusCode < 500 }

// Is5xx returns true if the HTTP status code is in the 5xx range, indicating a server error response.
func (r *Response) Is5xx() bool { return r.StatusCode >= 500 && r.StatusCode < 600 }
