package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	r1, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatalln("error: ", err)
	}
	defer r1.Body.Close()

	r2, err := http.Head("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatalln("error: ", err)
	}
	defer r2.Body.Close()

	form := url.Values{}
	form.Add("foo", "bar")
	// r3, _ := http.Post("https://www.google.com/robots.txt", "application/x-www-form-urlencoded",
	// 	strings.NewReader(form.Encode()))
	r3, err := http.PostForm("http://www.google.com/robots.txt", form)
	if err != nil {
		log.Fatalln("error: ", err)
	}

	defer r3.Body.Close()
	fmt.Printf("%s", r3.Body)
}
