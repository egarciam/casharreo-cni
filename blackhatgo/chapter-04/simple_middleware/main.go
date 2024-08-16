package main

import (
	"fmt"
	"log"
	"net/http"
)

type logger struct {
	Inner http.Handler
}

func (l *logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Print("start: ", r)
	l.Inner.ServeHTTP(w, r)
	log.Print("finish")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world\n")
}

func main() {
	f := http.HandlerFunc(hello)
	l := logger{Inner: f}
	http.ListenAndServe(":6555", &l)
}
