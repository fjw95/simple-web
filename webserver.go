package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Page")
}

func static(w http.ResponseWriter, r *http.Request) {
	var t, err = template.ParseFiles("/home/fiky/work/src/simpleweb.com/simple/web/server/page.html")
	if err != nil {

		fmt.Fprintf(w, err.Error())
	}
	t.Execute(w, "")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/static", static)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:	", err)
	}
}
