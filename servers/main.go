package main

import (
	"fmt"
	"net/http"
)

type Server struct{}

func (s Server) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello Server with Go!")
}

func main() {
	var s Server
	http.ListenAndServe("localhost:4000", s)

	exercise()
}
