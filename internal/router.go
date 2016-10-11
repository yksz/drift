package internal

import (
	"fmt"
	"net/http"

	"goji.io"
	"goji.io/pat"
	"golang.org/x/net/context"
)

func Router() http.Handler {
	mux := goji.NewMux()
	mux.HandleFuncC(pat.Get("/"), index)
	mux.Handle(pat.Get("/list/*"), http.StripPrefix("/list", NewFileServer(Conf.RootDir)))
	return mux
}

func index(c context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "index page")
}

func RedirectLocal(newPath string, w http.ResponseWriter, r *http.Request) {
	if q := r.URL.RawQuery; q != "" {
		newPath += "?" + q
	}
	w.Header().Set("Location", newPath)
	w.WriteHeader(http.StatusMovedPermanently)
}
