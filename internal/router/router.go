package router

import (
	"fmt"
	"net/http"

	"goji.io"
	"goji.io/pat"
	"golang.org/x/net/context"
)

func NewRouter() http.Handler {
	mux := goji.NewMux()
	mux.HandleFuncC(pat.Get("/"), index)
	return mux
}

func index(c context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "index page")
}
