package internal

import (
	"fmt"
	"net/http"

	"github.com/yksz/drift/internal/webapi"

	"goji.io"
	"goji.io/pat"
)

func Router() http.Handler {
	mux := goji.NewMux()

	mux.HandleFunc(pat.Get("/"), index)
	mux.Handle(pat.Get("/static/*"), http.StripPrefix("/static/", http.FileServer(http.Dir("public/static"))))
	mux.Handle(pat.Get("/files/*"), http.StripPrefix("/files/", http.FileServer(http.Dir(Conf.RootDir))))

	// API
	mux.Handle(pat.Get("/api/list"), webapi.ListAPI(Conf.RootDir))
	mux.Handle(pat.Get("/api/list/"), RedirectHandler("/api/list"))

	mux.Handle(pat.Get("/api/open"), webapi.OpenAPI(Conf.RootDir))
	mux.Handle(pat.Get("/api/open/"), RedirectHandler("/api/open"))

	// HTML
	mux.Handle(pat.Get("/list"), RedirectHandler("/list/"))
	mux.HandleFunc(pat.Get("/list/*"), serveFileFunc("public/views/list.html"))

	mux.Handle(pat.Get("/open"), RedirectHandler("/open/"))
	mux.HandleFunc(pat.Get("/open/*"), serveFileFunc("public/views/open.html"))

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
