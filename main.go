package main

import (
	"flag"
	"log"
	"net/http"

	"./internal/router"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8888", "port number")
	flag.Parse()
}

func main() {
	r := router.NewRouter()
	log.Printf("Listening and serving HTTP on %s\n", port)
	http.ListenAndServe(":"+port, r)
}
