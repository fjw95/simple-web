package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	} else if r.URL.Path == "/login" {
		login(w, r)
		return
	} else if r.URL.Path == "/success-login" {
		successLogin(w, r)
		return
	}
	http.NotFound(w, r)
	return

}
func successLogin(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("/home/fiky/work/src/github.com/fjw95/simple-web/success-login.html")
	t.Execute(w, nil)
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //Parse	url	parameters	passed,	then	parse	the	response	packet	for	the	POST	body	(request	body)
	//	attention:	If	you	do	not	call	ParseForm	method,	the	following	data	can	not	be	obtained	form
	fmt.Println(r.Form) //	print	information	on	server	side.
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Welcome") //	write	data	to	response
}
func login(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("method:", r.Method) //get	request	method

	if r.Method == "GET" {
		t, _ := template.ParseFiles("/home/fiky/work/src/github.com/fjw95/simple-web/login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		//var data = r.Form["username"]

		//	logic	part	of	log	in
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}
func main() {
	mux := &MyMux{}
	err := http.ListenAndServe(":8080", mux) //	setting	listening	port
	if err != nil {
		log.Fatal("ListenAndServe:	", err)
	}
}
