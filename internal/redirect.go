package internal

import "net/http"

func Redirect(w http.ResponseWriter, r *http.Request, path string) {
	if q := r.URL.RawQuery; q != "" {
		path += "?" + q
	}
	w.Header().Set("Location", path)
	w.WriteHeader(http.StatusMovedPermanently)
}

func RedirectHandler(path string) http.Handler {
	return &redirectHandler{path: path}
}

type redirectHandler struct {
	path string
}

func (h *redirectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Redirect(w, r, h.path)
}
