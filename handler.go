package latte

import (
	"net/http"

	"github.com/go-chi/chi"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func ToHTTPHandlerFunc(h HandlerFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}

func HF(h HandlerFunc) func(w http.ResponseWriter, r *http.Request) {
	return ToHTTPHandlerFunc(h)
}

type Mux struct {
	*chi.Mux
}

func NewRouter() *Mux {
	return &Mux{Mux: chi.NewRouter()}
}

type Router interface {
	chi.Router
	EDelete(pattern string, h Handler)
	EGet(pattern string, h Handler)
	EHead(pattern string, h Handler)
	EOptions(pattern string, h Handler)
	EPatch(pattern string, h Handler)
	EPost(pattern string, h Handler)
	EPut(pattern string, h Handler)
}

func (m *Mux) EDelete(path string, h Handler) {
	m.Method("DELETE", path, h)
}

func (m *Mux) EGet(path string, h Handler) {
	m.Method("GET", path, h)
}

func (m *Mux) EHead(path string, h Handler) {
	m.Method("HEAD", path, h)
}

func (m *Mux) EOptions(path string, h Handler) {
	m.Method("OPTIONS", path, h)
}

func (m *Mux) EPatch(path string, h Handler) {
	m.Method("PATCH", path, h)
}

func (m *Mux) EPost(path string, h Handler) {
	m.Method("POST", path, h)
}

func (m *Mux) EPut(path string, h Handler) {
	m.Method("PUT", path, h)
}

func EDelete(r chi.Router, path string, h Handler) {
	r.Method("DELETE", path, h)
}

func EGet(r chi.Router, path string, h Handler) {
	r.Method("GET", path, h)
}

func EHead(r chi.Router, path string, h Handler) {
	r.Method("HEAD", path, h)
}

func EOptions(r chi.Router, path string, h Handler) {
	r.Method("OPTIONS", path, h)
}

func EPatch(r chi.Router, path string, h Handler) {
	r.Method("PATCH", path, h)
}

func EPost(r chi.Router, path string, h Handler) {
	r.Method("POST", path, h)
}

func EPut(r chi.Router, path string, h Handler) {
	r.Method("PUT", path, h)
}
