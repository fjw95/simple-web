package main

import (
	"fmt"
	"net/http"
)

type MyMuxxxx struct {
}

func (p *MyMuxxxx) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	} else if r.URL.Path == "/nakam" {
		nakam(w, r)
		return
	}
	http.NotFound(w, r)
	return
}
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello	myroute!")
}

func nakam(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "NAKAM")
}

func main() {
	mux := &MyMuxxxx{}
	http.ListenAndServe(":8080", mux)
}
