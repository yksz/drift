package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/yksz/drift/internal"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8888", "port number")
	flag.Parse()
}

func main() {
	log.Printf("Listening and serving HTTP on %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, internal.Router()))
}
