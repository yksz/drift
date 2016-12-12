package webapi

import (
	"fmt"
	"net/http"
	"path"
	"path/filepath"
)

func OpenAPI(root string) http.Handler {
	return &openAPI{root: filepath.Clean(root)}
}

type openAPI struct {
	root string
}

func (a *openAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filename := path.Clean(query.Get("path"))
	if err := a.serve(filename, w, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (a *openAPI) serve(filename string, w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "%s", filename)
	return nil
}
