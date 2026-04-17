package httpx

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
)

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

func (r *Response) IsError() bool { return r.err != nil }
func (r *Response) Error() error  { return r.err }
func (r *Response) Cookies() []*http.Cookie {
	if r.IsError() {
		return []*http.Cookie{}
	}
	return r.HttpResponse.Cookies()
}
func (r *Response) Header(key string) string {
	return r.Headers.Get(key)
}
func (r *Response) Text() string     { return string(r.Bytes()) }
func (r *Response) Bytes() []byte    { return r.body }
func (r *Response) Json(v any) error { return json.Unmarshal(r.body, v) }
