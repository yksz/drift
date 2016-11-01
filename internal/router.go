package internal

import (
	"fmt"
	"net/http"

	"goji.io"
	"goji.io/pat"
)

func Router() http.Handler {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/"), index)
	mux.Handle(pat.Get("/static/*"), http.StripPrefix("/static/", http.FileServer(http.Dir("public/static"))))

	mux.Handle(pat.Get("/list"), RedirectHandler("/list/"))
	mux.HandleFunc(pat.Get("/list/*"), serveFileFunc("public/views/list.html"))

	mux.Handle(pat.Get("/api/file"), FileHandler(Conf.RootDir))
	mux.Handle(pat.Get("/api/file/"), RedirectHandler("/api/file"))
	return mux
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "index page")
}

func serveFileFunc(filepath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath)
	}
}
