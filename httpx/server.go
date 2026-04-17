package httpx

import (
	"encoding/json"
	"net/http"
)

type responseWriter struct {
	http.ResponseWriter
}

func (w *responseWriter) WithStatus(status int) ResponseWriter {
	w.WriteHeader(status)
	return w
}
func (w *responseWriter) WithHeader(key, value string) ResponseWriter {
	w.Header().Set(key, value)
	return w
}
func (w *responseWriter) WithCookie(cookie *http.Cookie) ResponseWriter {
	http.SetCookie(w.ResponseWriter, cookie)
	return w
}
func (w *responseWriter) WriteBytes(b []byte) (int, error) {
	return w.ResponseWriter.Write(b)
}
func (w *responseWriter) WriteString(s string) (int, error) {
	return w.ResponseWriter.Write([]byte(s))
}
func (w *responseWriter) WriteJson(v any) (int, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return 0, err
	}

	return w.ResponseWriter.Write(bytes)
}
func (w *responseWriter) WriteJsonIndent(v any, prefix, indent string) (int, error) {
	bytes, err := json.MarshalIndent(v, prefix, indent)
	if err != nil {
		return 0, err
	}
	return w.ResponseWriter.Write(bytes)
}
func (w *responseWriter) Redirect(url string, code int) {
	http.Redirect(w.ResponseWriter, nil, url, code)
}

type ResponseWriter interface {
	http.ResponseWriter
	WithHeader(key, value string) ResponseWriter
	WithCookie(cookie *http.Cookie) ResponseWriter
	WithStatus(int) ResponseWriter
	WriteBytes(b []byte) (int, error)
	WriteString(s string) (int, error)
	WriteJson(v any) (int, error)
	WriteJsonIndent(v any, prefix, indent string) (int, error)
	Redirect(url string, code int)
}

func HandleFunc(pattern string, handler func(w ResponseWriter, r *http.Request)) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		handler(&responseWriter{w}, r)
	})
}

func ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, nil)
}
