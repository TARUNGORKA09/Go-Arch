package main

import (
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"net/http"
)


type Person struct {
	First string
}

func main() {

	p1 := Person{First: "Tarun"}

	sm := http.DefaultServeMux

	sm.HandleFunc("/",p1.GetPerson)

}

func (p *Person) GetPerson(rw http.ResponseWriter, r *http.Request){

	password := r.Header["username"]

	fmt.Fprintf(rw,"Password",password)

}


//func HashPassword(password string) ([]byte.,error) {
	bcrypt.
//}