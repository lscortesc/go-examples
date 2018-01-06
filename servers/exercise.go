package main

import (
	"fmt"
	"net/http"
)

type String string

func (s String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, s)
}

type Greetings struct {
	Greeting string
	Punct    string
	Who      string
}

func (g Greetings) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, g.Greeting, g.Punct, g.Who)
}

func exercise() {
	http.Handle("/string", String("I'm a string"))
	http.Handle("/greetings", Greetings{"Hello", ":", "Gophers!"})

	http.ListenAndServe("localhost:4001", nil)
}
