package main

import (
	"fmt"
	//"log"
	"net/http"
)

type files int

func (f files) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("content-type", "application/json")
	fmt.Fprintf(rw, "<h1>Hello world</h1>")
	fmt.Println(r.Header["Accept-Language"])
	fmt.Println(r.Host)
	fmt.Println(r.URL)
	fmt.Println(r.Response)

}

func main() {

	//l := log.Logger
	//fmt.Print(l)
	var d files
	http.ListenAndServe(":8080", d)

}
