package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type trivial struct{}

func (t *trivial) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("Executing negroni mdw")
	next(w, r)
}

func main() {
	r := mux.NewRouter()
	n := negroni.Classic()
	n.Use(&trivial{})
	n.UseHandler(r)

	http.ListenAndServe(":33000", n)
}
