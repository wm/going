package main

import (
    "fmt"
    "net/http"
)

type Hello struct{}

type String string

type Struct struct {
  Greeting string
  Punct    string
  Who      string
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello!")
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, s)
}

func (s Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, s.Greeting + s.Punct + s.Who)
}

// Package http serves HTTP requests using any value that implements http.Handler
func main() {
    h := Hello{}
    http.Handle("/", h)
    http.Handle("/string", String("I'm a frayed knot."))
    http.Handle("/struct", Struct{"Hello", ":", "Gophers!"})
    http.ListenAndServe("localhost:4000", nil)
}

